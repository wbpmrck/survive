package battle
import (
	"survive/server/logic/consts/nature"
	"survive/server/logic/dataStructure"
	"survive/server/logic/dataStructure/attribute"
)


//Character可以embed这个类型，来实现战斗
type Warrior struct {
	Size *attribute.Attribute //占据长度
	NormalAttackSection *dataStructure.Section //普通攻击范围
	NormalAttackNature nature.Nature //普通攻击属性
	ActSeq *attribute.Attribute //行动顺序

	OP,MaxOp,OPRecover *attribute.Attribute //行动点数，最大行动点数,行动点数恢复速度（每一个时间tick恢复量）

	EachActionCostOP *attribute.Attribute //普通攻击\技能等动作，需要消耗的Op数量
	EachOpMoveDistance *attribute.Attribute //每一个Op可以移动的距离长度(越大代表移动速度越快)

	AttackPhysical,AttackMagical *attribute.Attribute //物理、魔法攻击力
	CriticalRatePhysical,CriticalRateMagical *attribute.Attribute //物理、魔法暴击率(单位：千分位)
	DefencePhysical,DefenceMagical *attribute.Attribute //物理、魔法防御力
	FleeRatePhysical,FleeRateMagical *attribute.Attribute //物理、魔法闪避率(单位：千分位)
	HP,MaxHp *attribute.Attribute //生命值，最大生命值
	AP,MaxAP *attribute.Attribute //怒气值，最大怒气值

	//下面是在战斗中才有意义的一些属性
	Battlefield *Battle //所处的战场
	Position *attribute.Attribute //战斗中，当前所处的位置(战场为一条线，左边是0，右边为增大方向)
}
