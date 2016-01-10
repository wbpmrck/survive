package battle
import (
	"survive/server/logic/math"
	"survive/server/logic/consts/nature"
)

//表示一个可以战斗的单位
type Fightable interface {

}

//Character可以embed这个类型，来实现战斗
type Warrior struct {
	Size int //占据长度
	NormalAttackSection math.Section //普通攻击范围
	NormalAttackNature nature.Nature //普通攻击属性
	ActSeq int //行动顺序

	OP,MaxOp,OPRecover int //行动点数，最大行动点数,行动点数恢复速度（每一个时间tick恢复量）

	EachActionCostOP int //普通攻击\技能等动作，需要消耗的Op数量
	EachOpMoveDistance int //每一个Op可以移动的距离长度(越大代表移动速度越快)

	AttackPhysical,AttackMagical int //物理、魔法攻击力
	CriticalRatePhysical,CriticalRateMagical int //物理、魔法暴击率(单位：千分位)
	DefencePhysical,DefenceMagical int //物理、魔法防御力
	FleeRatePhysical,FleeRateMagical int //物理、魔法闪避率(单位：千分位)
	HP,MaxHp int //生命值，最大生命值
	AP,MaxAP int //怒气值，最大怒气值

	//下面是在战斗中才有意义的一些属性
	Battlefield *Battle //所处的战场
	Position int //战斗中，当前所处的位置(战场为一条线，左边是0，右边为增大方向)
}
