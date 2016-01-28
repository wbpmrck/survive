package builtIn
import (
	"survive/server/logic/skill/effect/bases"
	"fmt"
	"survive/server/logic/skill/effect"
	"math"
	"survive/server/logic/rule/event"
	"survive/server/logic/dataStructure/attribute"
)

/*
	抵抗类效果：
	抵抗 属性减少 类效果

	说明：
	1、为了首个版本简单可用，本效果对所有属性减少效果都有效
	2、可以指定抵抗数值 （抵抗百分比 后面再实现）

	作用阶段：
	1、其他效果在实体上的 BeforePutOn
		也就是说，当一个单个的效果施加进来的时候，该抵抗效果会发生作用
	2、修改其他效果的 value

	PS:开发本效果主要为了测试整个效果激活体系的健壮性
 */

//一个标记，如果一个 属性修正效果上，被标记了这个属性，那么表示这个效果是一个削弱效果，并且被抵抗过。属性值就是抵抗的值
type AttributeDecResistance struct {
	*bases.EffectBase

	HandlerId event.HandlerId
	Amount   float64  //抵抗值(在config的时候，判断抵抗值不能小于0) PS:当然，如果想用这一个效果来制作多个效果出来(比如加深削弱)，也是可以的
}

//效果施加
func(self *AttributeDecResistance) PutOn(from, target effect.EffectCarrier) bool{
	//如果该效果可以被添加到对象上
	if self.EffectBase.PutOn(from,target){

		//功能1：对于在本效果释放之后再释放的效果，直接进行抵抗
		//注册一个处理事件，到对象的 效果生效前 阶段
		self.HandlerId = target.OnBeforePutOnEffect(event.NewEventHandler(func (contextParams ...interface{}) (isCancel bool,handleResult string){
			isCancel = false //默认不取消

			//看看即将生效的效果，是不是 "属性修正效果"
			effectWantPutOn,ok := contextParams[0].(*AttributeModify)
			if ok{
				//如果是，则获取效果修正的所有属性
				allAttr := effectWantPutOn.GetAllAttr()
				for _,v := range allAttr{
					// 检查修正值是否为负数,并且进行处理
					if v.GetValue().Get()<0.0{
						//获取削减值的绝对值
						decVal := -v.GetValue().Get()

						//如果抵消之后，最终削弱效果<=0,则只抵消 被削弱的部分
						resistVal := math.Min(decVal,self.Amount)
						//进行抵消(对修改效果的值，进行抵消加成)
						v.GetValue().Add(resistVal,self)
					}
				}
			}
			return
		}))

		//功能2：如果在此之前，已经中了减少属性的效果，则立刻进行修正(这是因为，属性修正的效果是在putOn就已经触发了的)
		allModifyEffectsBeforeMe := target.GetAllEffects()["AttributeModify"]
		if allModifyEffectsBeforeMe !=nil{
//			fmt.Printf(" 在此之前，已经中了减少属性的效果，则立刻进行修正 %v , %v \n",allModifyEffectsBeforeMe,len(allModifyEffectsBeforeMe))
			for i,_:= range allModifyEffectsBeforeMe{
				//转化为 "属性修正效果"对象
				e,ok := allModifyEffectsBeforeMe[i].(*AttributeModify)
				if ok{
//					fmt.Printf(" 如果是,则获取效果修正的所有属性 \n")

					//如果是，则获取效果修正的所有属性
					allAttr := e.GetAllAttr()
					for _,v := range allAttr {
						// 检查修正值是否为负数,并且进行处理(这样，当这个效果消失的时候，他会正确的恢复用户属性，而不会多加)
						if v.GetValue().Get() < 0.0 {
							//获取削减值的绝对值
							decVal := -v.GetValue().Get()

							//如果抵消之后，最终削弱效果<=0,则只抵消 被削弱的部分
							resistVal := math.Min(decVal,self.Amount)
							//进行抵消(对修改效果的值，进行抵消加成)
							v.GetValue().Add(resistVal,self)

							//改变了修正效果之后，修改用户当前属性“被属性修改效果改变的大小”
							t := target.(attribute.AttributeCarrier)
							attr:=t.GetAttr(v.GetName())
							if attr != nil{
//								fmt.Printf(" 修正被削弱的属性:%v 加成 %v \n",attr.GetName(),resistVal)
								attr.GetValue().Add(resistVal,e)
							}
						}
					}
				}
			}
		}

		return true
	}else{
		return false
	}
}


//效果移除
func(self *AttributeDecResistance) Remove() bool{
	targetOld := self.Target //保存之前作用的效果对象
	//取消事件监听（意味着对新投放的效果失去首次抵抗作用）
	targetOld.Off(self.HandlerId)
	//如果此时,抵抗效果可以被移除
	if self.EffectBase.Remove(){

		//获取对象当前已经中了的减少属性的效果，则立刻进行逆修正（抵抗效果消失了，之前的抵抗逻辑要回退）
		allModifyEffectsBeforeMe := targetOld.GetAllEffects()["AttributeModify"]
		for i,_:= range allModifyEffectsBeforeMe{
			//转化为 "属性修正效果"对象
			e,ok := allModifyEffectsBeforeMe[i].(*AttributeModify)
			if ok {
//				fmt.Printf(" 该效果是削弱效果，并且被标记了抵抗，进行逆向处理 \n")
				//如果是，则获取效果修正的所有属性
				allAttr := e.GetAllAttr()
				for _,v := range allAttr {
					// 检查值是否被自己修改过,并且回退这些修改
					addedByMe := v.GetValue().UndoAllAddBy(self)
					if addedByMe != nil{
						//确实有过修改，则继续修改这个效果对人物属性的修改
//						fmt.Printf(" 确实有过修改，%v\n",addedByMe)

						t := targetOld.(attribute.AttributeCarrier)
						attr:=t.GetAttr(v.GetName())
						if attr != nil{
							//修改 modify 效果作用在对象上的值，还原
//							fmt.Printf(" 修改 modify 效果作用在对象上的值，%v, 修改值:%v\n",attr,-addedByMe.AdditionVal)
							attr.GetValue().Add(-addedByMe.AdditionVal,e)
//							fmt.Printf(" 修改 modify 效果作用在对象上的值后，%v\n",attr)
						}
					}
				}
			}
		}
		return true
	}else{
		return false
	}
}


//配置修正值
func(self *AttributeDecResistance) Config(args ...interface{}){
	if len(args)>0{
		self.Amount = math.Max(args[0].(float64),0) //参数1：抵抗值
	}else{
		self.Amount = 0
	}
}
//显示效果信息
func(self *AttributeDecResistance) GetInfo() string{
	operator := "";
	if self.Amount >=0{
		operator="+"
	}
	//显示样例：属性削弱效果抵抗+1
	return fmt.Sprintf("属性削弱效果抵抗%s%v",operator,self.Amount)
}
