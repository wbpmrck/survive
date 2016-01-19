package rule
import (
	"survive/server/logic/skill/effect"
	"survive/server/logic/rule/event"
)
/*
	代表一个可以被施加效果的单位
 */

type EffectCarrier interface {
	event.EventEmitter

	/*
		获取该单位目前所携带的所有效果
		key:效果ID
		value:效果列表
		一个效果可能被多次叠加，那么就会有多个在一个key的slice里
	 */
	GetAllEffects() map[string][]effect.Effect
	/*
		给单位尝试加上一个效果。
		该单位应该执行以下操作：
		1、触发自身的所有 OnBeforePutOn 事件处理函数，把本次希望添加的*Effect传入
		2、此时只要有一个处理函数认为需要取消，则PutOn返回false
		3、如果没有取消，则进行效果添加操作，调用effect.PutOn
		4、PutOn成功
	 */
	PutOnEffect(effect effect.Effect,from EffectCarrier) bool
	/*
		注册一个处理函数,在实体 被施加效果前调用
		该处理函数如果返回true,则实体会继续进行后续操作，如果返回false,则会让实体跳过本次操作(并触发 OnCancelPutOn)
	 */
	OnBeforePutOnEffect(handler *event.EventHandler) event.HandlerId
	/*
		注册一个处理函数，在实体 被释放效果 之后调用
	 */
	OnAfterPutOnEffect (handler *event.EventHandler) event.HandlerId
	/*
		实体的 效果施放动作 被取消之后触发一次
	 */
	OnCancelPutOnEffect (handler *event.EventHandler) event.HandlerId


	/*
		给单位尝试取消一个效果。
		该单位应该执行以下操作：
		1、触发自身的所有 OnBeforeRemove 事件处理函数，把本次希望删除的*Effect传入
		2、此时只要有一个处理函数认为需要取消，则 Remove 返回false
		3、如果没有取消，则 Remove 成功
	 */
	RemoveEffect(effect effect.Effect) bool
	/*
		注册一个处理函数,在实体 被取消效果 前调用
		该处理函数如果返回true,则实体会继续进行后续操作，如果返回false,则会让实体跳过本次操作(并触发 OnCancelRemove)
	 */
	OnBeforeRemoveEffect(handler *event.EventHandler) event.HandlerId
	/*
		注册一个处理函数，在实体 被取消效果 之后调用
	 */
	OnAfterRemoveEffect (handler *event.EventHandler) event.HandlerId
	/*
		实体的 效果施放动作 被取消之后触发一次
	 */
	OnCancelRemoveEffect (handler *event.EventHandler) event.HandlerId
}

//提供效果携带者的默认实现
//其他实体如果没有特殊的需求，可以直接内嵌本类型实现简单的效果管理
type EffectCarrierBase struct {
	*event.EventEmitterBase

	/*
	key:效果ID
	value:效果列表
	一个效果可能被多次叠加，那么就会有多个在一个key的slice里
	 */
	allEffects map[string][]effect.Effect
}
//获取所有效果
func(self *EffectCarrierBase) GetAllEffects() map[string][]effect.Effect{
	return self.allEffects
}
/*
		给单位尝试加上一个效果。
		该单位应该执行以下操作：
		1、触发自身的所有 OnBeforePutOn 事件处理函数，把本次希望添加的*Effect传入
		2、此时只要有一个处理函数认为需要取消，则PutOn返回false
		3、如果没有取消，则PutOn成功
	 */
func(self *EffectCarrierBase) PutOnEffect(effect effect.Effect,from,target EffectCarrier) bool{
	onBeforePutOnResults:= self.Emit("Before-PutOnEffect",effect)

	//只要有一个处理函数认为需要取消，则PutOn返回false
	for _,r := range onBeforePutOnResults{
		if r !=nil && r.IsCancel{
			//执行cancel 阶段
			self.Emit("Cancel-PutOnEffect",effect)

			//返回(跳过 after 阶段)
			return false
		}
	}

	//没有被取消，则添加效果到自身效果集合
	effectSlot,exist := self.allEffects[effect.GetName()]
	if !exist{
		self.allEffects[effect.GetName()] = make([]effect.Effect,0)
		effectSlot = self.allEffects[effect.GetName()]
	}
	//记录单位具有的效果
	effectSlot = append(effectSlot,effect)

	//否则继续触发  after 阶段函数
	self.Emit("After-PutOnEffect",effect)
	return true
}
/*
	注册一个处理函数,在实体 被施加效果前调用
	该处理函数如果返回true,则实体会继续进行后续操作，如果返回false,则会让实体跳过本次操作(并触发 OnCancelPutOn)
*/
func(self *EffectCarrierBase) OnBeforePutOnEffect(handler *event.EventHandler) event.HandlerId{
	return self.On("Before-PutOnEffect",handler)
}
/*
	注册一个处理函数，在实体 被释放效果 之后调用
 */
func(self *EffectCarrierBase) OnAfterPutOnEffect (handler *event.EventHandler) event.HandlerId{
	return self.On("After-PutOnEffect",handler)
}
/*
	实体的 效果施放动作 被取消之后触发一次
 */

func(self *EffectCarrierBase)OnCancelPutOnEffect (handler *event.EventHandler) event.HandlerId{
	return self.On("Cancel-PutOnEffect",handler)
}


/*
	给单位尝试取消一个效果。
	该单位应该执行以下操作：
	1、触发自身的所有 OnBeforeRemove 事件处理函数，把本次希望删除的*Effect传入
	2、此时只要有一个处理函数认为需要取消，则 Remove 返回false
	3、如果没有取消，则 Remove 成功
 */

func(self *EffectCarrierBase)RemoveEffect(effect effect.Effect) bool{

	onBeforeRemoveResults:= self.Emit("Before-RemoveEffect",effect)

	//只要有一个处理函数认为需要取消，则 RemoveEffect 返回false
	for _,r := range onBeforeRemoveResults{
		if r !=nil && r.IsCancel{
			//执行cancel 阶段
			self.Emit("Cancel-RemoveEffect",effect)

			//返回(跳过 after 阶段)
			return false
		}
	}
	//没有被取消，则从自身效果集合移除该效果
	effectSlot,exist := self.allEffects[effect.GetName()]
	if exist{
		//根据效果类型，获取该类型的所有效果列表
		effectSlot = self.allEffects[effect.GetName()]

		//查找列表里有无指定的效果对象
		for i,_ := range effectSlot{
			//比对接口
			if effectSlot[i] == effect{
				//移除对应的效果
				effectSlot = append(effectSlot[:i], effectSlot[i+1:]...)

				//触发  after 阶段函数
				self.Emit("After-RemoveEffect",effect)
				return true
			}
		}
		return false
	}else{
		return false
	}
}


/*
	注册一个处理函数,在实体 被取消效果 前调用
	该处理函数如果返回true,则实体会继续进行后续操作，如果返回false,则会让实体跳过本次操作(并触发 OnCancelRemove)
 */

func(self *EffectCarrierBase)OnBeforeRemoveEffect(handler *event.EventHandler) event.HandlerId{
	return self.On("Before-RemoveEffect",handler)
}
/*
	注册一个处理函数，在实体 被取消效果 之后调用
 */

func(self *EffectCarrierBase)OnAfterRemoveEffect (handler *event.EventHandler) event.HandlerId{
	return self.On("After-RemoveEffect",handler)
}
/*
	实体的 效果施放动作 被取消之后触发一次
 */

func(self *EffectCarrierBase)OnCancelRemoveEffect (handler *event.EventHandler) event.HandlerId{
	return self.On("Cancel-RemoveEffect",handler)
}
func NewEffectCarrier()*EffectCarrierBase{
	return &EffectCarrierBase{
		EventEmitterBase:event.NewEventEmitter(),
		allEffects:make(map[string][]effect.Effect,0),
	}
}