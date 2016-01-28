package builtIn
import (
	"survive/server/logic/skill/effect/bases"
	"survive/server/logic/skill/effect"
//	"fmt"
	"survive/server/logic/dataStructure/attribute"
	"bytes"
	"strconv"
)

//表示对属性进行修正（按数值修正）的一种效果
type AttributeModify struct {
	*bases.EffectBase


//	AttrKey  string //影响的属性id
//	AttrName string //影响的属性名称
//	Amount   float64  //影响到属性修正值
//	Amount   attribute.Value  //影响到属性修正值
}

//效果施加
func(self *AttributeModify) PutOn(from, target effect.EffectCarrier) bool{
	//如果该效果可以被添加到对象上
	if self.EffectBase.PutOn(from,target){
		t := target.(attribute.AttributeCarrier)
		//增加属性
		allAttr := self.GetAllAttr()
		for _,v := range allAttr{
			attr:=t.GetAttr(v.GetName())
			if attr != nil{
				//对作用对象的属性进行修正。修正的值是效果对应属性的最终值(这个值也可以被改变)
				attr.GetValue().Add(v.GetValue().Get(),self)
			}
		}
		return true
	}else{
		return false
	}
//	self.Target.Attributes[self.attrKey].GetValue.Add(self.amount)
}


//效果移除
func(self *AttributeModify) Remove() bool{
	tOld := self.Target //先保存作用对象 （因为base在remove的时候会删除）
	//如果此时效果可以被移除
	if self.EffectBase.Remove(){
//		fmt.Printf("如果此时效果可以被移除 \n")
		//先获取效果所有的加成属性信息
		allAttr := self.GetAllAttr()
		for _,v := range allAttr{
			//获取作用对象
			t := tOld.(attribute.AttributeCarrier)
			//找到作用对象的对应属性
			attr:=t.GetAttr(v.GetName())
			//对属性进行还原
			if attr != nil{
//				fmt.Printf("尝试对属性 %v 进行还原 \n",attr)
//				attr.GetValue().Add(0-v.GetValue().Get())
				attr.GetValue().UndoAllAddBy(self)
			}
		}

		return true
	}else{
		return false
	}
//	self.Target.Attributes[self.attrKey].GetValue().Add(-self.amount)

}

//配置修正值
//每次调用config,都相当于添加1或n个属性，属性参数一字排开
//例子：Config("STR","力量",10,"AGI","敏捷",20,...)
func(self *AttributeModify) Config(args ...interface{}){
	if len(args)>0{
		name :=""
		desc :=""
		raw := 0.0
		for i:=0;i<len(args);i++ {
			switch i % 3 {
			case 0:
				name = args[i].(string)  //参数0：属性名
			case 1:
				desc = args[i].(string)  //参数1：属性描述
			case 2:
				raw = args[i].(float64)  //参数2：属性原始值

				//写入效果的属性列表(如果重复存在，第一次写入的为准)
				self.AddAttr(attribute.NewAttribute(name,desc,raw))
			}
		}
	}
}
//显示效果信息
func(self *AttributeModify) GetInfo() string{
	//获取效果所有属性加成信息
	allAttr := self.GetAllAttr()
	var buffer bytes.Buffer
	for _,v := range allAttr{
		buffer.WriteString(v.GetName()) //力量
		if v.GetValue().Get() >=0{
			buffer.WriteString("+") //力量+
		}

		buffer.WriteString(strconv.FormatFloat(v.GetValue().Get(),'f',4,64)) //力量+1  或 力量-22
		buffer.WriteString(",")
	}
	if buffer.Len()>0{
		buffer.Truncate(buffer.Len()-1)
	}
	//样例：力量+1,敏捷+2,智力+3
	return buffer.String()
}
