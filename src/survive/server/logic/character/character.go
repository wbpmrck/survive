package character
import (
	"survive/server/logic/player"
	"survive/server/logic/dataStructure/attribute"
	"survive/server/logic/rule/event"
	"survive/server/logic/skill/effect"
	"survive/server/logic/skill"
)

type Character struct {
	Player *player.Player //所属玩家
	*event.EventEmitterBase //角色是一个事件发生器
   //Sex,Age,Weight,Height int //性别，年龄，体重，身高
   //Str,Agi,Pcp,Int,Vit,Luk,Und int //力量 敏捷 感知 智力 体质 运气 悟性
	*attribute.AttributeCarrierBase //角色也是一个属性携带者
	*effect.EffectCarrierBase //角色是一个效果携带者
	*skill.SkillCarrierBase //角色 同时也是一个技能携带者
	Id string //角色Id
	GivenName,FamilyName string //姓名

//	Attributes map[string]*attribute.Attribute //存储所有属性名  （不再单独实现，直接内嵌 AttributeCarrierBase）
//	*battle.Warrior //内嵌“战斗者”类型，实现战斗能力
}

//创建一个角色对象
func NewCharacter(id string,givenName,familyName string,attributes map[string]*attribute.AttributeLike) *Character{
	var c = &Character{
		Id:id,
		GivenName:givenName,
		FamilyName:familyName,
		EventEmitterBase:event.NewEventEmitter(),
		AttributeCarrierBase:attribute.NewAttributeCarrier(),
		EffectCarrierBase:effect.NewEffectCarrier(),
	}
	return c
}
