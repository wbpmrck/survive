package battle
import (
	"survive/server/logic/consts/nature"
	"survive/server/logic/dataStructure"
	"survive/server/logic/dataStructure/attribute"
	"survive/server/logic/character"
	"math"
)


//Character可以embed这个类型，来实现战斗
type Warrior struct {
	*character.Character //战斗者 首先是一个角色

	Size *attribute.Attribute //占据长度
	NormalAttackSection *dataStructure.Section //普通攻击范围
	NormalAttackNature nature.Nature //普通攻击属性
	ActSeq *attribute.ComputedAttribute //行动顺序

	OP *attribute.Attribute //行动点数
	MaxOp,OPRecover *attribute.ComputedAttribute //最大行动点数,行动点数恢复速度（每一个时间tick恢复量）

	EachActionCostOP *attribute.ComputedAttribute //普通攻击\技能等动作，需要消耗的Op数量
	EachOpMoveDistance *attribute.Attribute //每一个Op可以移动的距离长度(越大代表移动速度越快)

	AttackPhysical,AttackMagical *attribute.ComputedAttribute //物理、魔法攻击力
	CriticalRatePhysical,CriticalRateMagical *attribute.ComputedAttribute //物理、魔法暴击率
	DefencePhysical,DefenceMagical *attribute.ComputedAttribute //物理、魔法防御力
	FleeRatePhysical,FleeRateMagical *attribute.ComputedAttribute //物理、魔法闪避率
	HitRatePhysical,HitRateMagical *attribute.ComputedAttribute //物理、魔法命中率

	HP *attribute.Attribute //生命值
	MaxHp *attribute.ComputedAttribute //最大生命值
	HpRecover *attribute.ComputedAttribute //生命值回复速度

	AP *attribute.Attribute //怒气值
	MaxAP *attribute.Attribute //最大怒气值
	ApRecover *attribute.ComputedAttribute //怒气值回复速度

	//下面是在战斗中才有意义的一些属性
	BattleIn *Battle //所处的战场
	Position dataStructure.BattlePos //战斗中，当前所处的位置(战场为一条线，左边是0，右边为增大方向)
}


//创建一个战斗角色
func NewWarrior(character *character.Character,normalAttackNature nature.Nature,size,attackFrom,attackTo,op,hp,ap,eachOpMoveDistance,maxAp float64) *Warrior{
	w:=&Warrior{
		Character:character,
		Size:attribute.NewAttribute("size","大小",size),
		NormalAttackSection:dataStructure.NewSection(attackFrom,attackTo),
		NormalAttackNature:normalAttackNature,
		EachOpMoveDistance:attribute.NewAttribute("EachOpMoveDistance","每一个Op可以移动的距离长度",eachOpMoveDistance),
		OP:attribute.NewAttribute("OP","当前行动点数",op),
		HP:attribute.NewAttribute("HP","当前Hp",hp),
		AP:attribute.NewAttribute("AP","当前怒气",ap),
		MaxAP:attribute.NewAttribute("MaxAP","最大怒气",maxAp),
	}
	allAttr:= w.Character.GetAllAttr()
	AGI := allAttr[attribute.AGI]
	AWARE := allAttr[attribute.AWARE]
	STR := allAttr[attribute.STR]
	VIT := allAttr[attribute.VIT]
	INT := allAttr[attribute.INT]
	LUCK := allAttr[attribute.LUCK]
	UNDERSTAND := allAttr[attribute.UNDERSTAND]
	//计算属性：行动顺序
	w.ActSeq = attribute.NewComputedAttribute("ActSeq","行动顺序",0,
		func(dependencies... attribute.AttributeLike) float64{
			base := 50.0
			agi :=dependencies[0].GetValue().Get()
			aware :=dependencies[1].GetValue().Get()
		return base +( 3*agi + 5.3*math.Floor(agi/7) ) - ( 0.2*aware + 0.5*math.Floor(aware/9) )
	},AGI,AWARE)

	//计算属性：最大行动点数
	w.MaxOp = attribute.NewComputedAttribute("MaxOp","最大行动点数",0,
		func(dependencies... attribute.AttributeLike) float64{
			base := 128.0
			agi :=dependencies[0].GetValue().Get()
			vit :=dependencies[1].GetValue().Get()
			return base +math.Floor( agi/3 ) + math.Floor( vit/9 )
		},AGI,VIT)

	//计算属性：行动点数恢复速度(每秒)
	w.OPRecover= attribute.NewComputedAttribute("OPRecover","行动点数恢复速度",0,
		func(dependencies... attribute.AttributeLike) float64{
			base := 28.0
			agi :=dependencies[0].GetValue().Get()
			vit :=dependencies[1].GetValue().Get()
			return base +( 0.6*math.Floor( agi/3 )) + ( 0.2*math.Floor( vit/5 ))
		},AGI,VIT)
	//计算属性：普通攻击\技能等动作，需要消耗的Op数量
	w.EachActionCostOP =attribute.NewComputedAttribute("EachActionCostOP","动作消耗的Op数量",0,
		func(dependencies... attribute.AttributeLike) float64{
			base := 128.0
			agi :=dependencies[0].GetValue().Get()
			return base -math.Floor( agi/50 )
		},AGI)
	//计算属性：物理攻击力
	w.AttackPhysical =attribute.NewComputedAttribute("AttackPhysical","物理攻击力",0,
		func(dependencies... attribute.AttributeLike) float64{
			str :=dependencies[0].GetValue().Get()
			vit :=dependencies[1].GetValue().Get()
			return (str + 11*math.Floor( str/9 )) + (0.3*vit + 2*math.Floor( vit/8 ))
		},STR,VIT)
	//计算属性：魔法攻击力
	w.AttackMagical =attribute.NewComputedAttribute("AttackMagical","魔法攻击力",0,
		func(dependencies... attribute.AttributeLike) float64{
			intel :=dependencies[0].GetValue().Get()
			vit :=dependencies[1].GetValue().Get()
			return (intel + 9*math.Floor( intel/7 )) + (0.3*vit + 2*math.Floor( vit/8 ))
		},INT,VIT)
	//计算属性：物理暴击率
	w.CriticalRatePhysical =attribute.NewComputedAttribute("CriticalRatePhysical","物理暴击率",0,
		func(dependencies... attribute.AttributeLike) float64{
			base:=0.15
			str :=dependencies[attribute.STR].GetValue().Get()
			agi :=dependencies[attribute.AGI].GetValue().Get()
			vit :=dependencies[attribute.VIT].GetValue().Get()
			luk :=dependencies[attribute.LUCK].GetValue().Get()
			aware :=dependencies[attribute.AWARE].GetValue().Get()

			return base+(0.02*math.Floor( luk/3 ))+(0.01*math.Floor( aware/5 )) + (0.01*math.Floor( (str-agi+vit)/30 ))
		})

	//计算属性：魔法暴击率
	w.CriticalRateMagical =attribute.NewComputedAttribute("CriticalRateMagical","魔法暴击率",0,
		func(dependencies... attribute.AttributeLike) float64{
			base:=0.15
			intel :=dependencies[attribute.INT].GetValue().Get()
			luk :=dependencies[attribute.LUCK].GetValue().Get()
			aware :=dependencies[attribute.AWARE].GetValue().Get()

			return base+(0.02*math.Floor( luk/3 ))+(0.01*math.Floor( aware/5 )) +(0.01*math.Floor( intel/7 ))
		})

	//计算属性：物理防御力
	w.DefencePhysical =attribute.NewComputedAttribute("DefencePhysical","物理防御力",0,
		func(dependencies... attribute.AttributeLike) float64{
			str :=dependencies[attribute.STR].GetValue().Get()
			vit :=dependencies[attribute.VIT].GetValue().Get()
			aware :=dependencies[attribute.AWARE].GetValue().Get()

			return (math.Floor( (str+aware)/4 )) + (vit + 1.3*math.Floor( vit/3 ))
		})
	//计算属性：魔法防御力
	w.DefenceMagical =attribute.NewComputedAttribute("DefenceMagical","魔法防御力",0,
		func(dependencies... attribute.AttributeLike) float64{
			intel :=dependencies[attribute.INT].GetValue().Get()
			vit :=dependencies[attribute.VIT].GetValue().Get()
			aware :=dependencies[attribute.AWARE].GetValue().Get()

			return  0.3*vit+2.3* math.Floor( (intel+vit+aware)/4 )
		})
	//计算属性：物理闪避率
	w.FleeRatePhysical =attribute.NewComputedAttribute("FleeRatePhysical","物理闪避率",0,
		func(dependencies... attribute.AttributeLike) float64{
			base:=0.05
			str :=dependencies[attribute.STR].GetValue().Get()
			vit :=dependencies[attribute.VIT].GetValue().Get()
			agi :=dependencies[attribute.AGI].GetValue().Get()
			luk :=dependencies[attribute.LUCK].GetValue().Get()
			aware :=dependencies[attribute.AWARE].GetValue().Get()

			return base+(0.01*math.Floor( (agi+luk)/4 )) + (0.01*math.Floor( aware/5 )) - ( 0.01 * math.Floor( (str+vit)/15 ))
		})

	//计算属性：魔法闪避率
	w.FleeRateMagical =attribute.NewComputedAttribute("FleeRateMagical","魔法闪避率",0,
		func(dependencies... attribute.AttributeLike) float64{
			base:=0.05
			agi :=dependencies[attribute.AGI].GetValue().Get()
			intel :=dependencies[attribute.INT].GetValue().Get()
			luk :=dependencies[attribute.LUCK].GetValue().Get()
			aware :=dependencies[attribute.AWARE].GetValue().Get()

			return base+(0.01*math.Floor( (luk/4 )) + ( 0.01 * math.Floor( (agi+intel+aware)/15 )))
		})
	//计算属性：物理命中率
	w.HitRatePhysical =attribute.NewComputedAttribute("HitRatePhysical","物理命中率",0,
		func(dependencies... attribute.AttributeLike) float64{
			base:=1.0
			luk :=dependencies[attribute.LUCK].GetValue().Get()
			aware :=dependencies[attribute.AWARE].GetValue().Get()

			return base+(0.01*math.Floor( (aware+luk)/4 ))
		})

	//计算属性：魔法命中率
	w.HitRateMagical =attribute.NewComputedAttribute("HitRateMagical","魔法命中率",0,
		func(dependencies... attribute.AttributeLike) float64{
			base:=1.0
			luk :=dependencies[attribute.LUCK].GetValue().Get()
			intel :=dependencies[attribute.INT].GetValue().Get()

			return base+(0.01*math.Floor( (intel+luk)/4 ))
		})

	//计算属性：最大生命值
	w.MaxHp =attribute.NewComputedAttribute("MaxHp","最大生命值",0,
		func(dependencies... attribute.AttributeLike) float64{
			base:=120.0
			str :=dependencies[attribute.STR].GetValue().Get()
			vit :=dependencies[attribute.VIT].GetValue().Get()
			intel :=dependencies[attribute.INT].GetValue().Get()

			return base+str+vit+(23*math.Floor(vit/3)) - 3*math.Floor(intel/3)
		})

	//计算属性：生命值回复(每秒)
	w.HpRecover =attribute.NewComputedAttribute("HpRecover","生命值回复",0,
		func(dependencies... attribute.AttributeLike) float64{
			base:=1.0
			vit :=dependencies[attribute.VIT].GetValue().Get()

			return base+0.1*vit+(1*math.Floor(vit/5))
		})
	//计算属性：怒气值回复(每秒)
	w.ApRecover =attribute.NewComputedAttribute("ApRecover","怒气值回复",0,
		func(dependencies... attribute.AttributeLike) float64{
			base:=0.1
			aware :=dependencies[attribute.AWARE].GetValue().Get()

			return base+0.05*aware+(0.25*math.Floor(aware/7))
		})
	return w
}
