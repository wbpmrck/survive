package skill
import "survive/server/logic/rule/event"

//事件
const (
	EVENT_BEFORE_SKILL_ITEM_RELEASE ="BEFORE_SKILLITEM_RELEASE" //技能项释放（前）
	EVENT_SKILL_ITEM_RELEASE ="SKILLITEM_RELEASE" //技能项释放（ING）
	EVENT_AFTER_SKILL_ITEM_RELEASE ="AFTER_SKILLITEM_RELEASE" //技能项释放（后）
	EVENT_CANCEL_SKILL_ITEM_RELEASE ="CANCEL_SKILLITEM_RELEASE" //技能项被取消释放
)
/*
	定义一种技能持有者的标准实现。
	该接口描述了一切可以释放技能的实体应该具备的能力。
	其他对象可以直接内嵌本结构，来具备使用技能的能力
 */

type SkillCarrier interface {
	event.EventEmitter
	//获取所有技能
	GetAllSkills() []*Skill
//	/*
//		订阅一个事件，在技能释放者释放前触发
//		PS:在每一个SkillItem的效果释放前就触发，且只触发一次
//	 */
//	OnBeforeReleaseSkillItem(handler *event.EventHandler) event.HandlerId
//	/*
//		订阅一个事件，在技能释放者释放后触发
//		PS:在每一个SkillItem的效果释放后就触发，且只触发一次
//	 */
//	OnAfterReleaseSkillItem (handler *event.EventHandler) event.HandlerId
}

//skillCarrierBase 并不内嵌eventEmitter,而是让其他受体直接内嵌Emitter来完成事件操作
type SkillCarrierBase struct {
	skills []*Skill
}
func(self *SkillCarrierBase) GetAllSkills() []*Skill{
	return self.skills
}

func NewSkillCarrierBase() *SkillCarrierBase{
	return &SkillCarrierBase{
		skills:make([]*Skill,0),
	}
}


