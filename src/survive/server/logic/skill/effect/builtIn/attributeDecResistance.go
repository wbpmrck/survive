package builtIn
import (
	"survive/server/logic/skill/effect/bases"
	"fmt"
	"survive/server/logic/skill/effect"
	"survive/server/logic/dataStructure"
	"survive/server/logic/rule/composite"
	"math"
	"survive/server/logic/rule/event"
)

/*
	抵抗类效果：
	抵抗 属性减少 类效果

	说明：
	1、可以指定抵抗某种属性的减少效果
	2、可以指定抵抗数值

	作用阶段：
	1、其他效果在实体上的 BeforePutOn
	2、修改其他效果的Amount

	PS:开发本效果主要为了测试整个效果激活体系的健壮性
 */
type AttributeDecResistance struct {
	bases.EffectBase

	AttrKey  string //影响的属性id
	AttrName string //影响的属性名称
	Amount   float64  //抵抗值(在config的时候，判断抵抗值不能小于0) PS:当然，如果想用这一个效果来制作多个效果出来(比如加深削弱)，也是可以的
}

//效果施加
func(self *AttributeDecResistance) PutOn(from, target composite.AttributeAndEffectCarrier) bool{
	//如果该效果可以被添加到对象上
	if self.EffectBase.PutOn(from,target){
		//注册一个处理事件，到对象的 效果生效前 阶段
		target.OnBeforePutOnEffect(event.NewEventHandler(func (contextParams ...interface{}) (isCancel bool,handleResult string){
			isCancel = false //默认不取消

			//看看即将生效的效果，是不是 "属性修正效果"
			effectWantPutOn,ok := contextParams[0].(*AttributeModify)
			if ok{
				//如果是，则检查修正值是否为负数
				if effectWantPutOn.Amount<0{
					//对修正值进行抵抗处理
					effectWantPutOn.Amount += self.Amount

					//抵抗处理并不能变成加强处理，所以如果修正后>=0,则直接取消本次效果加成好了
					if effectWantPutOn.Amount >=0{
						isCancel = true
					}
				}
			}
			return
		}))
		return true
	}else{
		return false
	}
}


//效果移除
func(self *AttributeDecResistance) Remove() bool{
	//如果此时效果可以被移除
	if self.EffectBase.Remove(){
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
func(self *AttributeDecResistance) Config(args ...interface{}){
	if len(args>0){
		self.Amount = math.Max(args[0].(float64),0) //参数0：抵抗值
	}else{
		self.Amount = 0
	}
}
//显示效果信息
func(self *AttributeDecResistance) GetInfo() string{
	operator := "+";
	if self.Amount <0{
		operator="-"
	}
	//显示样例：力量 削弱抵抗 + 1
	return fmt.Sprintf("%v 削弱抵抗%s %v",self.AttrName,operator,self.Amount)
}

func init(){
	effect.RegisterFactory("AttributeDecResistance",func() effect.Effect{
		return &AttributeDecResistance{
			bases.EffectBase:bases.NewBase("AttributeDecResistance"),
		}
	})
}
