package bases
import (
	"survive/server/logic/dataStructure"
	"survive/server/logic/time"
	"survive/server/logic/skill/effect"
//	"fmt"
	"survive/server/logic/dataStructure/attribute"
)

/*
	效果的实现基础

	1、处理效果和效果携带者之间的关系
	2、管理所有效果都需要做的事情，包括效果状态、效果发出者和接收者、效果的开始时间、结束时间等
	3、效果在生效、取消的时候，都需要和效果携带者发生交互，这些也是在这里做的
 */

type EffectBase struct {
	*attribute.AttributeCarrierBase
	Holder effect.Effect //持有者
	//效果的使用者，和受众
	From interface{}
	Target effect.EffectCarrier
	PutOnTime dataStructure.Time//效果生效时间
	RemoveTime dataStructure.Time //效果结束时间
	Alive bool //效果是否存在
	Name string //效果名
	Level int //效果等级
}

func(self *EffectBase) SetLevel(level int){
	self.Level = level
}
//获取效果的名字
func(self *EffectBase) String() string{
//	return fmt.Sprintf("%v[%v]",self.Name,self.Alive)
	return self.Holder.GetInfo()
}
//获取效果的名字
func(self *EffectBase) GetName() string{
	return self.Name
}
//尝试给对象添加一个效果
func(self *EffectBase) PutOn(from interface{}, target effect.EffectCarrier) bool{
	//如果效果还没有对象产生
	if self.Target ==nil{

		//判断接收对象是否可以放置本效果
		isTargetReceive := target.PutOnEffect(self.Holder,from)

		if isTargetReceive{
			self.PutOnTime = time.GetNow()
			self.From = from
			self.Target = target
			self.Alive = true
			return true
		}else{
			//效果不被对方接收(比如：可能对方在效果 OnBeforePutOn 事件中，被其他效果取消了)
			return false
		}

	}else{
		//效果已经被附加过了，不能被重新附加
		return false
	}
}
//效果移除
func(self *EffectBase) Remove() bool{
	//如果当前效果还有效、有接受对象
	if self.Alive && self.Target!=nil {

		//判断接收对象是否可以取消本效果
		isTargetRemove := self.Target.RemoveEffect(self.Holder)
//		fmt.Printf("如果对象答应取消该效果:%v \n",isTargetRemove)

		//如果对象答应取消该效果
		if isTargetRemove {
			self.RemoveTime = time.GetNow()
			self.From = nil
			self.Target = nil
			self.Alive = false
			return true
		}else{
			return false
		}
	}else{
		return false
	}
}



//效果是否存在
func(self *EffectBase) IsAlive() bool{
	return self.Alive
}
//获取效果的发出方
func(self *EffectBase) GetFrom() interface{} {

	return self.From
}
//获取效果的作用方
func(self *EffectBase) GetTarget() effect.EffectCarrier {
	return self.Target
}
//获取效果的开始作用时间
func(self *EffectBase) GetPutOnTime() dataStructure.Time {
	return self.PutOnTime
}
//获取效果的结束作用时间
func(self *EffectBase) GetRemoveTime() dataStructure.Time {
	return self.RemoveTime
}

//空实现，仅仅为了满足接口定义
func(self *EffectBase)Config(args ...interface{}){

}
//空实现，仅仅为了满足接口定义
func(self *EffectBase) GetInfo() string{
	return "EffectBase"
}

func NewBase(name string,holder effect.Effect) *EffectBase{
	return &EffectBase{
		Name:name,
		Alive:false,
		Holder:holder,
		AttributeCarrierBase:attribute.NewAttributeCarrier(),
	}
}