package main

import (
	"survive/server/logic/character"
	"survive/server/logic/consts/nature"
	"survive/server/logic/dataStructure/attribute"
	"survive/server/logic/skill/effect"
	"fmt"
	"survive/server/logic/battle"
"survive/server/logic/skill/effect/builtIn"
)


func main(){
	builtIn.RegBuiltInEffects()
	//创建2个角色
	ch1 := character.NewCharacter("id1","giveName1","familyName1",nil)
	ch2 := character.NewCharacter("id2","giveName2","familyName2",nil)

	//设置基本属性
	ch1.SetAttr(attribute.NewAttribute(attribute.AGI,"敏捷",10))
	ch1.SetAttr(attribute.NewAttribute(attribute.STR,"力量",10))
	ch1.SetAttr(attribute.NewAttribute(attribute.INT,"智力",10))
	ch1.SetAttr(attribute.NewAttribute(attribute.VIT,"体力",10))
	ch1.SetAttr(attribute.NewAttribute(attribute.LUCK,"运气",10))
	ch1.SetAttr(attribute.NewAttribute(attribute.AWARE,"感知",10))
	ch1.SetAttr(attribute.NewAttribute(attribute.UNDERSTAND,"悟性",10))

	warrior1 := battle.NewWarrior(ch1,nature.Physical,12,0,200,30,200,20,12,30)

	ch2.SetAttr(attribute.NewAttribute(attribute.AGI,"敏捷",10))
	ch2.SetAttr(attribute.NewAttribute(attribute.STR,"力量",10))
	ch2.SetAttr(attribute.NewAttribute(attribute.INT,"智力",10))
	ch2.SetAttr(attribute.NewAttribute(attribute.VIT,"体力",10))
	ch2.SetAttr(attribute.NewAttribute(attribute.LUCK,"运气",10))
	ch2.SetAttr(attribute.NewAttribute(attribute.AWARE,"感知",10))
	ch2.SetAttr(attribute.NewAttribute(attribute.UNDERSTAND,"悟性",10))

	warrior2 := battle.NewWarrior(ch2,nature.Physical,12,0,200,30,200,20,12,30)

	//创建一个效果
	modifyEffect := effect.Create("AttributeModify")

	//配置该效果要修正的属性值
	modifyEffect.Config(attribute.STR,20.0)
	//给角色2添加效果
	modifyEffect.PutOn(warrior1,warrior2)

	//检查属性
	if warrior2.GetAttr(attribute.STR).GetValue().Get() != 30{
		panic(fmt.Sprintf("str must be 30,but now is %v",warrior2.GetAttr(attribute.STR).GetValue().Get()))
	}
	fmt.Printf("all attr is: %v",warrior2.GetAllAttr())
}
