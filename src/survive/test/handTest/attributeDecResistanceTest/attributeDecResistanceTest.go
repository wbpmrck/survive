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

/*
	测试效果：抵抗属性减少效果
 */

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
	modifyEffect.Config(attribute.STR,"力量",-20.0)

	//模拟场景：角色1给角色2添加减力量效果
	modifyEffect.PutOn(warrior1,warrior2)

	//检查属性，看减力量效果是否生效
	if warrior2.GetAttr(attribute.STR).GetValue().Get() != -10{
		panic(fmt.Sprintf("str must be -10,but now is %v",warrior2.GetAttr(attribute.STR).GetValue().Get()))
	}

	//添加一个抵抗减属性的 效果
	decResistanceEffect := effect.Create("AttributeDecResistance")
	//所有减属性类的效果，全部抵抗10
	decResistanceEffect.Config(10.0)
	//安装效果
	decResistanceEffect.PutOn(warrior1,warrior2)
	fmt.Printf(" 1.1 看减力量效果是否生效:all attr is: %v \n",warrior2.GetAllAttr())
	//检查属性，应该有变化
	if warrior2.GetAttr(attribute.STR).GetValue().Get() != 0{
		panic(fmt.Sprintf("str must be 0,but now is %v",warrior2.GetAttr(attribute.STR).GetValue().Get()))
	}


	//todo:1.2 -- 再添加一个削弱效果少一点的效果，测试技能抵抗不能变成加成
	modifyEffect2 := effect.Create("AttributeModify")

	//智力-5，小于抵抗值，测试抵抗效果的修正
	modifyEffect2.Config(attribute.INT,"智力",-5.0)

	//模拟场景：角色1给角色2添加减力量效果
	modifyEffect2.PutOn(warrior1,warrior2)
	fmt.Printf(" 1.2 再添加一个削弱效果少一点的效果:all attr is: %v \n",warrior2.GetAllAttr())
	//检查属性，应该没有变化（被抵抗了）
	if warrior2.GetAttr(attribute.INT).GetValue().Get() != 10{
		panic(fmt.Sprintf("INT must be 10,but now is %v",warrior2.GetAttr(attribute.INT).GetValue().Get()))
	}


	//todo:2 -- 移除抵抗效果，看什么情况
	decResistanceEffect.Remove()
	fmt.Printf(" 2.1 移除抵抗效果:all attr is: %v \n",warrior2.GetAllAttr())
	//检查属性被影响
	if warrior2.GetAttr(attribute.STR).GetValue().Get() != -10{
		panic(fmt.Sprintf("str must be -10,but now is %v",warrior2.GetAttr(attribute.STR).GetValue().Get()))
	}
	if warrior2.GetAttr(attribute.INT).GetValue().Get() != 5{
		panic(fmt.Sprintf("INT must be 5,but now is %v",warrior2.GetAttr(attribute.INT).GetValue().Get()))
	}

	//todo:3 -- 移除属性加成效果，看什么情况
	modifyEffect.Remove()
	fmt.Printf(" 3.1 移除抵抗效果 modifyEffect:all attr is: %v \n",warrior2.GetAllAttr())
	if warrior2.GetAttr(attribute.STR).GetValue().Get() != 10{
		panic(fmt.Sprintf("str must be 10,but now is %v",warrior2.GetAttr(attribute.STR).GetValue().Get()))
	}
	if warrior2.GetAttr(attribute.INT).GetValue().Get() != 5{
		panic(fmt.Sprintf("INT must be 5,but now is %v",warrior2.GetAttr(attribute.INT).GetValue().Get()))
	}
	modifyEffect2.Remove()
	fmt.Printf(" 3.2 移除抵抗效果 modifyEffect2:all attr is: %v \n",warrior2.GetAllAttr())
	if warrior2.GetAttr(attribute.STR).GetValue().Get() != 10{
		panic(fmt.Sprintf("str must be 10,but now is %v",warrior2.GetAttr(attribute.STR).GetValue().Get()))
	}
	if warrior2.GetAttr(attribute.INT).GetValue().Get() != 10{
		panic(fmt.Sprintf("INT must be 10,but now is %v",warrior2.GetAttr(attribute.INT).GetValue().Get()))
	}
}
