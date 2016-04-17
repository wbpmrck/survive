package battle
import (
	"survive/server/logic/skill"
	"fmt"
	"survive/server/logic/dataStructure"
	"survive/server/logic/character"
)

const (
	ACTION_TYPE_MOVE int =iota  //移动
	ACTION_TYPE_NORMAL_ATTACK	//普通攻击
	ACTION_TYPE_GET_DAMAGE	//受到伤害
	ACTION_TYPE_SKILL_RELEASE	//释放技能

	ACTION_TYPE_CHARACTER_DEAD // 角色死亡
	ACTION_TYPE_ATTRIBUTE_CHANGE //属性变化

	ACTION_TYPE_CHARACTER_ENTER //角色加入

	ACTION_TYPE_START_BATTLE //战斗开始
	ACTION_TYPE_END_BATTLE //战斗结束
)
/*
	表示一个执行了的动作，里面包含动作的关键参数，和描述信息
 */

type ActionRecord interface {
	GetWarrior() *Warrior //日志的战斗者是谁
	String() string //必须可以文本化描述
	GetType() int //动作类型
}

type ActionRecordBase struct {
	from *Warrior
	C    *character.Character
	D    dataStructure.TimeSpan //距离战斗开始的流逝时间
	T    int
	Desc string
}
func(self *ActionRecordBase) GetWarrior() *Warrior{
	return self.from
}
func(self *ActionRecordBase)String() string{
	return self.Desc
}
func(self *ActionRecordBase)GetType() int{
	return self.T
}
func NewActionRecordBase(d dataStructure.TimeSpan,from *Warrior,t int,desc string) *ActionRecordBase{
	if from != nil{
		return &ActionRecordBase{
			from:from,
			C:from.Character,
			D:d,
			T:t,
			Desc:desc,
		}
	}else{
		return &ActionRecordBase{
			from:nil,
			C:nil,
			D:d,
			T:t,
			Desc:desc,
		}
	}
}
//战斗开始 类日志
type BattleInitAction struct{
	*ActionRecordBase
	FieldLength int //战场宽度
}
//角色初始化 类日志
type CharacterEnterAction struct{
	*ActionRecordBase
	Props map[string]float64 //需要记录的属性值
}
//数值变化 类日志
type AttrChangeAction struct{
	*ActionRecordBase
	AttrName string //变化的属性名
	AttrValue float64 //变化的属性值(计算之后的)
}
//移动 类日志
type MoveAction struct{
	*ActionRecordBase
	PosFrom,PosTo int //移动位置起止
}
//技能释放 类日志
type SkillReleaseAction struct{
	*ActionRecordBase
	skill *skill.Skill
}

//根据一个warrior事件，分析出对应的日志
func AnalyzeReportFromEvent(timeElapsed dataStructure.TimeSpan,eventParams ...interface{}) (record ActionRecord){
	evtName := eventParams[0].(string)
	from := eventParams[1].(*Warrior)

	switch evtName {
	case skill.EVENT_SKILL_ITEM_RELEASE:
		sk:= eventParams[2].(*skill.SkillItem).SkillParent
		record = &SkillReleaseAction{
			ActionRecordBase:NewActionRecordBase(timeElapsed,from,ACTION_TYPE_SKILL_RELEASE,
				fmt.Sprintf("%v 使用了技能 [%v]",from.GetShowName(),sk.Name)),
			skill:sk,
		}
	case EVENT_WARRIOR_MOVE:
		fromPos := eventParams[2].(int)
		toPos := eventParams[3].(int)
		record = &MoveAction{
			ActionRecordBase:NewActionRecordBase(timeElapsed,from,ACTION_TYPE_MOVE,
				fmt.Sprintf("%v 从[%v]移动到[%v]",from.GetShowName(),fromPos,toPos)),
			PosFrom:fromPos,
			PosTo:toPos,
		}
	case EVENT_WARRIOR_NORMAL_ATTACK_PRE:
		target := eventParams[2].([]*Warrior)
		record = NewActionRecordBase(timeElapsed,from,ACTION_TYPE_NORMAL_ATTACK,
			fmt.Sprintf("%v 对[%v]个对象:[%v]发起攻击",from.GetShowName(),len(target),target))
	case EVENT_WARRIOR_GET_DAMAGE:
		damageFrom := eventParams[2].(*Warrior)
		damage := eventParams[3].(float64)
		record = NewActionRecordBase(timeElapsed,from,ACTION_TYPE_GET_DAMAGE,
			fmt.Sprintf("%v 受到来自[%v]的[%v]点伤害,剩余hp:[%v]",from.GetShowName(),damageFrom.GetShowName(),damage,from.HP))
	case EVENT_WARRIOR_DEAD:
		record = NewActionRecordBase(timeElapsed,from,ACTION_TYPE_CHARACTER_DEAD,
			fmt.Sprintf("%v 死亡",from.GetShowName()))
	default:

	}
	return
}