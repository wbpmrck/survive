package character
import (
	"survive/server/logic/battle"
	"survive/server/logic/player"
	"survive/server/logic/dataStructure/attribute"
	"survive/server/logic/consts/nature"
)

type Character struct {
	Id string
	Player *player.Player //所属玩家
	GivenName,FamilyName string //姓名

	//Sex,Age,Weight,Height int //性别，年龄，体重，身高
	//Str,Agi,Pcp,Int,Vit,Luk,Und int //力量 敏捷 感知 智力 体质 运气 悟性
	Attributes map[string]*attribute.Attribute //存储所有属性名
	*battle.Warrior //内嵌“战斗者”类型，实现战斗能力
	Effects map[]
}
//给角色初始化战斗属性
func(self *Character) SetWarriorData(normalAttackNature nature.Nature,size,attackFrom,attackTo,op,hp,ap,eachOpMoveDistance,maxAp float64){
	self.Warrior = battle.NewWarrior(self,normalAttackNature,size,attackFrom,attackTo,op,hp,ap,eachOpMoveDistance,maxAp)
}

//创建一个角色对象
func NewCharacter(id string,givenName,familyName string,attributes map[string]*attribute.Attribute) *Character{
	var c = &Character{
		Id:id,
		GivenName:givenName,
		FamilyName:familyName,
	}
	return c
}
