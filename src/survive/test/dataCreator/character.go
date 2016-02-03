package dataCreator
import (
	"survive/server/logic/character"
	"survive/server/logic/dataStructure/attribute"
	"strconv"
)

var seed_character int = 0
func GetCharacter() *character.Character{
	seed_character++
	
	ch1 := character.NewCharacter("id"+strconv.FormatInt(int64(seed_character),10),"giveName"+strconv.FormatInt(int64(seed_character),10),"familyName"+strconv.FormatInt(int64(seed_character),10),nil)
	//设置基本属性
	ch1.SetAttr(attribute.NewAttribute(attribute.AGI,"敏捷",10))
	ch1.SetAttr(attribute.NewAttribute(attribute.STR,"力量",10))
	ch1.SetAttr(attribute.NewAttribute(attribute.INT,"智力",10))
	ch1.SetAttr(attribute.NewAttribute(attribute.VIT,"体力",10))
	ch1.SetAttr(attribute.NewAttribute(attribute.LUCK,"运气",10))
	ch1.SetAttr(attribute.NewAttribute(attribute.AWARE,"感知",10))
	ch1.SetAttr(attribute.NewAttribute(attribute.UNDERSTAND,"悟性",10))
	return ch1
}
func GetCharacters(n int) []*character.Character{

	characters := make([]*character.Character,0)
	for i:=1;i<=n;i++{
		characters = append(characters,GetCharacter())
	}

	return characters
}
