package builtIn

import (
	"testing"
	"survive/server/logic/character"
	"survive/server/logic/consts/nature"
	"survive/server/logic/dataStructure/attribute"
	"survive/server/logic/skill/effect"
)


func TestModifyEffect(t *testing.T){
	//创建2个角色
	ch1 := character.NewCharacter("id1","giveName1","familyName1",nil)
	ch2 := character.NewCharacter("id2","giveName2","familyName2",nil)

	//设置基本属性
	ch1.SetAttr(attribute.NewAttribute(attribute.AGI,"",10))
	ch1.SetAttr(attribute.NewAttribute(attribute.STR,"",10))
	ch1.SetAttr(attribute.NewAttribute(attribute.INT,"",10))
	ch1.SetAttr(attribute.NewAttribute(attribute.VIT,"",10))
	ch1.SetAttr(attribute.NewAttribute(attribute.LUCK,"",10))
	ch1.SetAttr(attribute.NewAttribute(attribute.AWARE,"",10))
	ch1.SetAttr(attribute.NewAttribute(attribute.UNDERSTAND,"",10))

	ch2.SetAttr(attribute.NewAttribute(attribute.AGI,"",10))
	ch2.SetAttr(attribute.NewAttribute(attribute.STR,"",10))
	ch2.SetAttr(attribute.NewAttribute(attribute.INT,"",10))
	ch2.SetAttr(attribute.NewAttribute(attribute.VIT,"",10))
	ch2.SetAttr(attribute.NewAttribute(attribute.LUCK,"",10))
	ch2.SetAttr(attribute.NewAttribute(attribute.AWARE,"",10))
	ch2.SetAttr(attribute.NewAttribute(attribute.UNDERSTAND,"",10))

	//初始化战斗属性
	ch1.SetWarriorData(nature.Physical,12,0,200,30,200,20,12,30)

	//创建一个效果
	modifyEffect := effect.Create("AttributeModify")

	modifyEffect.Config()
	//给角色添加效果
	modifyEffect.PutOn(ch1,ch2)

	//检查属性
}
