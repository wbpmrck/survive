package battle
import "survive/server/logic/skill"

const (
	ACTION_TYPE_MOVE int =iota  //移动
	ACTION_TYPE_NORMAL_ATTACK	//普通攻击
	ACTION_TYPE_BE_NORMAL_ATTACK	//受到 普通攻击
	ACTION_TYPE_SKILL_RELEASE	//释放技能
	ACTION_TYPE_BE_SKILL_RELEASE	//受到 释放技能
	ACTION_TYPE_STATUS_CHANGE // 状态变化
	ACTION_TYPE_ATTRIBUTE_CHANGE //属性变化
)
/*
	表示一个执行了的动作，里面包含动作的关键参数，和描述信息
 */

type ActionRecord interface {
	GetWarrior() *Warrior //日志的战斗者是谁
	String() string //必须可以文本化描述
	GetType() int //动作类型
	GetTypeName() string //动作类型描述
}

type SkillReleaseAction struct{
	skill *skill.Skill

}
