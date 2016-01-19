package character
import (
	"survive/server/logic/battle"
	"survive/server/logic/player"
	"survive/server/logic/dataStructure/attribute"
	"survive/server/logic/consts/nature"
	"survive/server/logic/rule/event"
	"survive/server/logic/rule"
)

type Character struct {
	*event.EventEmitterBase //角色是一个事件发生器
   //Sex,Age,Weight,Height int //性别，年龄，体重，身高
   //Str,Agi,Pcp,Int,Vit,Luk,Und int //力量 敏捷 感知 智力 体质 运气 悟性
	*rule.AttributeCarrierBase //角色也是一个属性携带者
	*rule.EffectCarrierBase //角色是一个效果携带者
	Id string
	Player *player.Player //所属玩家
	GivenName,FamilyName string //姓名

//	Attributes map[string]*attribute.Attribute //存储所有属性名  （不再单独实现，直接内嵌 AttributeCarrierBase）
	*battle.Warrior //内嵌“战斗者”类型，实现战斗能力
}
//给角色初始化战斗属性
func(self *Character) SetWarriorData(normalAttackNature nature.Nature,size,attackFrom,attackTo,op,hp,ap,eachOpMoveDistance,maxAp float64){
	self.Warrior = battle.NewWarrior(self,normalAttackNature,size,attackFrom,attackTo,op,hp,ap,eachOpMoveDistance,maxAp)
}

//创建一个角色对象
func NewCharacter(id string,givenName,familyName string,attributes map[string]*attribute.AttributeLike) *Character{
	var c = &Character{
		Id:id,
		GivenName:givenName,
		FamilyName:familyName,
		EventEmitterBase:event.NewEventEmitter(),
		AttributeCarrierBase:rule.NewAttributeCarrier(),
	}
	return c
}
