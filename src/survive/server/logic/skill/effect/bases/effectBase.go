package bases
import (
	"survive/server/logic/dataStructure"
	"survive/server/logic/rule"
)

/*
	效果的实现基础

	1、处理效果和效果携带者之间的关系
	2、管理所有效果都需要做的事情，包括效果状态、效果发出者和接收者、效果的开始时间、结束时间等
	3、效果在生效、取消的时候，都需要和效果携带者发生交互，这些也是在这里做的
 */

type EffectBase struct {
	//效果的使用者，和受众
	From, Target rule.EffectCarrier
	PutOnTime *dataStructure.Time//效果生效时间
	RemoveTime *dataStructure.Time //效果结束时间
	Alive bool //效果是否存在
	Id string //效果名
}
//获取效果的名字
func(self *EffectBase) GetId() string{
	return self.Id
}
//尝试给对象添加一个效果
func(self *EffectBase) PutOn(time *dataStructure.Time,from, target rule.EffectCarrier) bool{
	//如果效果还没有对象产生
	if self.Target ==nil{

		//判断接收对象是否可以放置本效果
		isTargetReceive := self.Target.PutOnEffect(self)

		if isTargetReceive{
			self.PutOnTime = time
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
func(self *EffectBase) Remove(time *dataStructure.Time) bool{
	//如果当前效果还有效、有接受对象
	if self.Alive && self.Target!=nil {

		//判断接收对象是否可以取消本效果
		isTargetRemove := self.Target.RemoveEffect(self)

		//如果对象答应取消该效果
		if isTargetRemove {
			self.RemoveTime = time
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
func(self *EffectBase) GetFrom() rule.EffectCarrier {
	return self.From
}
//获取效果的作用方
func(self *EffectBase) GetTarget() rule.EffectCarrier {
	return self.Target
}
//获取效果的开始作用时间
func(self *EffectBase) GetPutOnTime()*dataStructure.Time {
	return self.PutOnTime
}
//获取效果的结束作用时间
func(self *EffectBase) GetRemoveTime()*dataStructure.Time {
	return self.RemoveTime
}

//空实现，仅仅为了满足接口定义
func(self *EffectBase)Config(args ...interface{}){

}
//空实现，仅仅为了满足接口定义
func(self *EffectBase) GetInfo() string{
	return "EffectBase"
}

func NewBase(id string) *EffectBase{
	return &EffectBase{
		Id:id,
		Alive:false,
	}
}