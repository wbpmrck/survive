package builtIn
import (
	"survive/server/logic/skill/effect/bases"
	"survive/server/logic/skill/effect"
	"fmt"
	"survive/server/logic/dataStructure"
	"survive/server/logic/rule/composite"
)

//表示对属性进行修正（按数值修正）的一种效果
type AttributeModify struct {
	bases.EffectBase

	AttrKey  string //影响的属性id
	AttrName string //影响的属性名称
	Amount   float64  //影响到属性修正值
}

//效果施加
func(self *AttributeModify) PutOn(time *dataStructure.Time,from, target composite.AttributeAndEffectCarrier) bool{
	//如果该效果可以被添加到对象上
	if self.EffectBase.PutOn(time,from,target){
		//增加属性
		target.GetAttr(self.AttrKey).GetValue().Add(self.Amount)
		return true
	}else{
		return false
	}
//	self.Target.Attributes[self.attrKey].GetValue.Add(self.amount)
}


//效果移除
func(self *AttributeModify) Remove(time *dataStructure.Time) bool{
	//如果此时效果可以被移除
	if self.EffectBase.Remove(time){
		//减少属性
		target := self.Target.(composite.AttributeAndEffectCarrier)
		target.GetAttr(self.AttrKey).GetValue().Add(-self.Amount)
		return true
	}else{
		return false
	}
//	self.Target.Attributes[self.attrKey].GetValue().Add(-self.amount)

}

//配置修正值
func(self *AttributeModify) Config(args ...interface{}){
	if len(args>0){
		self.Amount = args[0].(float64) //参数0：属性修正值
	}else{
		self.Amount = 0
	}
}
//显示效果信息
func(self *AttributeModify) GetInfo() string{
	operator := "+";
	if self.Amount <0{
		operator="-"
	}
	//显示样例：力量 + 1
	return fmt.Sprintf("%v %s %v",self.AttrName,operator,self.Amount)
}

func init(){
	effect.RegisterFactory("AttributeModify",func() *effect.Effect{
		return &AttributeModify{
			bases.EffectBase:bases.NewBase("AttributeModify"),
		}
	})
}