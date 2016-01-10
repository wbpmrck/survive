package character
import (
	"survive/server/logic/consts/nature"
	"survive/server/logic/battle"
	"survive/server/logic/player"
)

type Character struct {
	Player *player.Player //所属玩家
	GivenName,FamilyName string //姓名

	//Sex,Age,Weight,Height int //性别，年龄，体重，身高
	//Str,Agi,Pcp,Int,Vit,Luk,Und int //力量 敏捷 感知 智力 体质 运气 悟性
	Attributes map[string]Attribute //存储所有属性名
	battle.Warrior //内嵌“战斗者”类型，实现战斗能力
}


//创建一个角色对象
func New() *Character{
	var c = &Character{
		NormalAttackNature:nature.Magical,
	}
	return c
}
