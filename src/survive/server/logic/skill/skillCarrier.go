package skill
import "survive/server/logic/rule/event"

/*
	定义一种技能持有者的标准实现。
	该接口描述了一切可以释放技能的实体应该具备的能力。
	其他对象可以直接内嵌本结构，来具备使用技能的能力
 */

type SkillCarrier interface {
	event.EventEmitter
	//获取所有技能
	GetAllSkills() []*Skill
}

type SkillCarrierBase struct {
	*event.EventEmitterBase
	skills []*Skill
}
func(self *SkillCarrierBase) GetAllSkills() []*Skill{
	return self.skills
}


