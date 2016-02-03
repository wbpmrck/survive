package battle
import (
	"survive/server/logic/consts/nature"
	"survive/server/logic/dataStructure"
	"survive/server/logic/dataStructure/attribute"
	"survive/server/logic/character"
	"math"
	"fmt"
	"time"
	"survive/server/logic/rule/event"
	"survive/server/logic/skill"
	"survive/server/logic/skill/targetChoose"
)

//事件
const (
	EVENT_WARRIOR_DEAD ="dead" //死亡

//	EVENT_WARRIOR_NEW_ACTION_RECORD ="NEW_ACTION_RECORD" //产生了新的动作日志

	EVENT_WARRIOR_SKILL_CHOOSE ="SKILL_CHOOSE" //技能动作选择阶段(该阶段决定了用户是否能主动发出某个技能)

	EVENT_WARRIOR_NORMAL_ATTACK_PRE ="NORMAL_ATTACK_PRE" //普通攻击动作阶段（前）
	EVENT_WARRIOR_NORMAL_ATTACK_AFTER ="NORMAL_ATTACK_AFTER" //普通攻击动作阶段（后）

)

//状态
const (
	WARRIOR_INACTIVE	int =0 //未激活的(从未获得过时间片)
	WARRIOR_ACTIVE 		int =1 //激活的
	WARRIOR_DEAD		int =2 //死亡的
)

//规则
//const (

//)

//Character可以embed这个类型，来实现战斗
type Warrior struct {
	*character.Character //战斗者 首先是一个角色

	Status int //状态
	Size *attribute.Attribute //占据长度
	NormalAttackSection *dataStructure.Section //普通攻击范围
	NormalAttackNature nature.Nature //普通攻击属性
	NormalAttackChooser targetChoose.TargetChooser //普通攻击的时候，chooser决定攻击对象
	ActSeq *attribute.ComputedAttribute //行动顺序

	OP *attribute.RegeneratedAttribute //行动点数
//	OP *attribute.Attribute //行动点数
	MaxOp,OPRecover *attribute.ComputedAttribute //最大行动点数,行动点数恢复速度（每一个时间tick恢复量）

	EachActionCostOP *attribute.ComputedAttribute //普通攻击\技能等动作，需要消耗的Op数量
	EachOpMoveDistance *attribute.Attribute //每一个Op可以移动的距离长度(越大代表移动速度越快)

	AttackPhysical,AttackMagical *attribute.ComputedAttribute //物理、魔法攻击力
	CriticalRatePhysical,CriticalRateMagical *attribute.ComputedAttribute //物理、魔法暴击率
	DefencePhysical,DefenceMagical *attribute.ComputedAttribute //物理、魔法防御力
	FleeRatePhysical,FleeRateMagical *attribute.ComputedAttribute //物理、魔法闪避率
	HitRatePhysical,HitRateMagical *attribute.ComputedAttribute //物理、魔法命中率

//	HP *attribute.Attribute //生命值
	HP *attribute.RegeneratedAttribute //生命值
	MaxHp *attribute.ComputedAttribute //最大生命值
	HpRecover *attribute.ComputedAttribute //生命值回复速度

//	AP *attribute.Attribute //怒气值
	AP *attribute.RegeneratedAttribute //怒气值
	MaxAP *attribute.Attribute //最大怒气值
	ApRecover *attribute.ComputedAttribute //怒气值回复速度

	//下面是在战斗中才有意义的一些属性
	BattleIn *Battle //所处的战场
	Position dataStructure.BattlePos //战斗中，当前所处的位置(战场为一条线，左边是0，右边为增大方向)

	//其他


}

////向外输出日志事件
//func(self *Warrior) FireNewRecord(r ActionRecord){
//	self.Emit(EVENT_WARRIOR_NEW_ACTION_RECORD,r)
//}

//激活
func (self *Warrior) Active(){
	fmt.Printf("开始激活[%s] \n",self)
	self.Status = WARRIOR_ACTIVE
	//初始化其他事项:
	//1.技能
	for _,v := range self.GetAllSkills(){
		fmt.Printf("-->开始安装技能[%s] \n",v.GetInfo())
		v.Install(self)
	}
}
//增减 hp
func (self *Warrior) AddHP(v float64) {

	//已经dead 的无法修改hp
	if self.Status == WARRIOR_ACTIVE {

		cur := self.HP.GetValue().Get()
		max := self.MaxHp.GetValue().Get()
		//判断上限
		if cur + v > max {
			v = max - cur
		}
		self.HP.GetValue().AddRaw(v)
		//判断死亡
		if self.HP.GetValue().Get() < 0 {
			self.Status = WARRIOR_DEAD
			//对外抛送事件
			self.Emit(EVENT_WARRIOR_DEAD)
		}
	}
}

//进行移动。
//目前移动逻辑很简单，移动到能够让对方角色进入攻击范围的位置
func (self *Warrior)MoveAction(){

}

//尝试进行普通攻击，返回是否成功进行攻击动作的标记
//PS:不管是否命中，只要发动了攻击动作，就视为 true
func (self *Warrior)NormalAttackAction() bool{
	attacked := false
	//先获取普通攻击的对象，如果有，进行普通攻击
	targets := self.getNormalAttackTargets()
	if len(targets>0){
		attacked = true
		//对每一个对象，进行攻击
		for _,t := range targets{
			self.attack(t)
		}
		//如果成功，则扣除动作点数
		self.OP.GetValue().AddRaw(self.EachActionCostOP.GetValue().Get())
	}
	return attacked
}
//对某个对象进行普通攻击操作
func (self *Warrior) attack(target *Warrior){

}

//获取当前普通攻击的对象
func (self *Warrior)getNormalAttackTargets()[]*Warrior{
	warriorTargets := make([]*Warrior,0)
	if self.NormalAttackChooser != nil{
		targets,err := self.NormalAttackChooser.Choose(self,EVENT_WARRIOR_NORMAL_ATTACK_PRE)

		//如果获取对象成功
		if !err && targets!= nil && len(targets)>0{
			for _,target := range targets{

				//尝试把对象转化为warrior
				w,errTrans := target.(*Warrior)
				if errTrans{
					warriorTargets = append(warriorTargets,w)
				}
			}
		}
	}
	return warriorTargets
}

//执行技能动作阶段，并返回成功释放的技能项列表(没有的话为0)
//PS:不管是否命中，只要发动了技能，就视为 有技能释放结果
func (self *Warrior)SkillAction() []*skill.SkillItem{
	skillItems := make([]*skill.SkillItem,0)
	//遍历所有技能动作钩子，执行并拿到执行结果，返回
	results := self.Emit(EVENT_WARRIOR_SKILL_CHOOSE)
	if len(results)>0{
		for _,r := range results{
			if r.HandleResult!=nil{
				skillItems = append(skillItems,r.HandleResult.(*skill.SkillItem))
			}
		}
		//如果释放成功，则扣除动作点数
		self.OP.GetValue().AddRaw(self.EachActionCostOP.GetValue().Get())
	}

	return skillItems
}
/*
	订阅阶段事件
	阶段：选择动作
	HandleFunc返回值含义
		isCancel:
			true: 无意义
			false:无意义
		handleResult interface{}:
			类型：*SkillItem
			含义：代表被释放成功的技能项
 */
func (self *Warrior) OnSkillChoose(handler *event.EventHandler) event.HandlerId{
	return self.On(EVENT_WARRIOR_SKILL_CHOOSE,handler)
}






/*
	角色进行动作
	1.通过OP的概念，把角色和时间进行了隔离，角色只需要行动到OP点数不够为止，而不再需要关注时间的概念
 */
func(self *Warrior) Act(){
	/*
		判断行动点数是否足够进行一次动作
	 */
	canAction := self.OP.GetValue().Get()>self.EachActionCostOP.GetValue().Get()

	//如果是，则进行动作选择阶段
	if canAction{

		/*
			执行阶段：【技能选择】
			1.在技能选择阶段,每个技能都可以通过在这个阶段注册事件，并以一定的概率释放自己
			2.一旦技能选择阶段有技能被释放，那么本次普通攻击阶段就会跳过
		 */
		releasedSkillItemRecords := self.SkillAction()
		//没有技能被释放
		if len(releasedSkillItemRecords)<1{
			//进行普通攻击
			attackRecords := self.NormalAttackAction()

			//如果无法普通攻击
			if !attackRecords {
				//尝试移动
				self.MoveAction()
				//移动过后，重新执行Action
				self.Act()
			}
		}
	}else{
		//如果无法完成一次动作(攻击、技能等)，则尝试移动
		self.MoveAction()
	}
	return
}

//开始行动
//接收一个时间片，开始决定自己的行动
//每次接收时间片，要返回一个动作列表，表示这次行动得到的反馈结果
func(self *Warrior) Receive(ts dataStructure.TimeSpan){
	fmt.Printf("Warrior:[%v] ActSeq:[%v] receive time:%v \n",self.GetShowName(),self.ActSeq.GetValue(),ts)

	//如果是非激活状态，则进行激活
	if self.Status == WARRIOR_INACTIVE{
		self.Active()
	}

	/*
		属性回复
	 */
	self.HP.TimeAcquire(ts.GameSpan)
	self.AP.TimeAcquire(ts.GameSpan)
	self.OP.TimeAcquire(ts.GameSpan)

	/*
		开始角色的行动：
	 */
	return self.Act()
}






//创建一个战斗角色
func NewWarrior(character *character.Character,normalAttackNature nature.Nature,normalAttackChooser targetChoose.TargetChooser,size,attackFrom,attackTo,op,hp,ap,eachOpMoveDistance,maxAp float64) *Warrior{
	w:=&Warrior{
		Character:character,
		Status:WARRIOR_INACTIVE,
		Size:attribute.NewAttribute("size","大小",size),
		NormalAttackSection:dataStructure.NewSection(attackFrom,attackTo),
		NormalAttackNature:normalAttackNature,
		NormalAttackChooser:normalAttackChooser,
		EachOpMoveDistance:attribute.NewAttribute("EachOpMoveDistance","每一个Op可以移动的距离长度",eachOpMoveDistance),
//		OP:attribute.NewAttribute("OP","当前行动点数",op),
//		HP:attribute.NewAttribute("HP","当前Hp",hp),
//		AP:attribute.NewAttribute("AP","当前怒气",ap),
		MaxAP:attribute.NewAttribute("MaxAP","最大怒气",maxAp),
	}
	allAttr:= w.Character.GetAllAttr()
	AGI := allAttr[attribute.AGI]
	AWARE := allAttr[attribute.AWARE]
	STR := allAttr[attribute.STR]
	VIT := allAttr[attribute.VIT]
	INT := allAttr[attribute.INT]
	LUCK := allAttr[attribute.LUCK]
//	UNDERSTAND := allAttr[attribute.UNDERSTAND]
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

	w.OP = attribute.NewRegeneratedAttribute("OP","当前行动点数",op,w.MaxOp.GetValue(),w.OPRecover.GetValue(),1*time.Second)

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
			str :=dependencies[0].GetValue().Get()
			agi :=dependencies[1].GetValue().Get()
			vit :=dependencies[2].GetValue().Get()
			luk :=dependencies[3].GetValue().Get()
			aware :=dependencies[4].GetValue().Get()

			return base+(0.02*math.Floor( luk/3 ))+(0.01*math.Floor( aware/5 )) + (0.01*math.Floor( (str-agi+vit)/30 ))
		},STR,AGI,VIT,LUCK,AWARE)

	//计算属性：魔法暴击率
	w.CriticalRateMagical =attribute.NewComputedAttribute("CriticalRateMagical","魔法暴击率",0,
		func(dependencies... attribute.AttributeLike) float64{
			base:=0.15
			intel :=dependencies[0].GetValue().Get()
			luk :=dependencies[1].GetValue().Get()
			aware :=dependencies[2].GetValue().Get()

			return base+(0.02*math.Floor( luk/3 ))+(0.01*math.Floor( aware/5 )) +(0.01*math.Floor( intel/7 ))
		},INT,LUCK,AWARE)

	//计算属性：物理防御力
	w.DefencePhysical =attribute.NewComputedAttribute("DefencePhysical","物理防御力",0,
		func(dependencies... attribute.AttributeLike) float64{
			str :=dependencies[0].GetValue().Get()
			vit :=dependencies[1].GetValue().Get()
			aware :=dependencies[2].GetValue().Get()

			return (math.Floor( (str+aware)/4 )) + (vit + 1.3*math.Floor( vit/3 ))
		},STR,VIT,AWARE)
	//计算属性：魔法防御力
	w.DefenceMagical =attribute.NewComputedAttribute("DefenceMagical","魔法防御力",0,
		func(dependencies... attribute.AttributeLike) float64{
			intel :=dependencies[0].GetValue().Get()
			vit :=dependencies[1].GetValue().Get()
			aware :=dependencies[2].GetValue().Get()

			return  0.3*vit+2.3* math.Floor( (intel+vit+aware)/4 )
		},INT,VIT,AWARE)
	//计算属性：物理闪避率
	w.FleeRatePhysical =attribute.NewComputedAttribute("FleeRatePhysical","物理闪避率",0,
		func(dependencies... attribute.AttributeLike) float64{
			base:=0.05
			str :=dependencies[0].GetValue().Get()
			vit :=dependencies[1].GetValue().Get()
			agi :=dependencies[2].GetValue().Get()
			luk :=dependencies[3].GetValue().Get()
			aware :=dependencies[4].GetValue().Get()

			return base+(0.01*math.Floor( (agi+luk)/4 )) + (0.01*math.Floor( aware/5 )) - ( 0.01 * math.Floor( (str+vit)/15 ))
		},STR,VIT,AGI,LUCK,AWARE)

	//计算属性：魔法闪避率
	w.FleeRateMagical =attribute.NewComputedAttribute("FleeRateMagical","魔法闪避率",0,
		func(dependencies... attribute.AttributeLike) float64{
			base:=0.05
			agi :=dependencies[0].GetValue().Get()
			intel :=dependencies[1].GetValue().Get()
			luk :=dependencies[2].GetValue().Get()
			aware :=dependencies[3].GetValue().Get()

			return base+(0.01*math.Floor( (luk/4 )) + ( 0.01 * math.Floor( (agi+intel+aware)/15 )))
		},AGI,INT,LUCK,AWARE)
	//计算属性：物理命中率
	w.HitRatePhysical =attribute.NewComputedAttribute("HitRatePhysical","物理命中率",0,
		func(dependencies... attribute.AttributeLike) float64{
			base:=1.0
			luk :=dependencies[0].GetValue().Get()
			aware :=dependencies[1].GetValue().Get()

			return base+(0.01*math.Floor( (aware+luk)/4 ))
		},LUCK,AWARE)

	//计算属性：魔法命中率
	w.HitRateMagical =attribute.NewComputedAttribute("HitRateMagical","魔法命中率",0,
		func(dependencies... attribute.AttributeLike) float64{
			base:=1.0
			luk :=dependencies[0].GetValue().Get()
			intel :=dependencies[1].GetValue().Get()

			return base+(0.01*math.Floor( (intel+luk)/4 ))
		},LUCK,INT)

	//计算属性：最大生命值
	w.MaxHp =attribute.NewComputedAttribute("MaxHp","最大生命值",0,
		func(dependencies... attribute.AttributeLike) float64{
			base:=120.0
			str :=dependencies[0].GetValue().Get()
			vit :=dependencies[1].GetValue().Get()
			intel :=dependencies[2].GetValue().Get()

			return base+str+vit+(23*math.Floor(vit/3)) - 3*math.Floor(intel/3)
		},STR,VIT,INT)

	//计算属性：生命值回复(每秒)
	w.HpRecover =attribute.NewComputedAttribute("HpRecover","生命值回复",0,
		func(dependencies... attribute.AttributeLike) float64{
			base:=1.0
			vit :=dependencies[0].GetValue().Get()

			return base+0.1*vit+(1*math.Floor(vit/5))
		},VIT)

	w.HP = attribute.NewRegeneratedAttribute("HP","当前Hp",hp,w.MaxHp.GetValue(),w.HpRecover.GetValue(),1*time.Second)

	//计算属性：怒气值回复(每秒)
	w.ApRecover =attribute.NewComputedAttribute("ApRecover","怒气值回复",0,
		func(dependencies... attribute.AttributeLike) float64{
			base:=0.1
			aware :=dependencies[0].GetValue().Get()

			return base+0.05*aware+(0.25*math.Floor(aware/7))
		},AWARE)
	w.AP = attribute.NewRegeneratedAttribute("AP","当前怒气",ap,w.MaxAP.GetValue(),w.ApRecover.GetValue(),1*time.Second)
	return w
}
