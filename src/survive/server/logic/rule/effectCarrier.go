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
	 */
	GetAllEffects() []*effect.Effect
	/*
		给单位尝试加上一个效果。
		该单位应该执行以下操作：
		1、触发自身的所有 OnBeforePutOn 事件处理函数，把本次希望添加的*Effect传入
		2、此时只要有一个处理函数认为需要取消，则PutOn返回false
		3、如果没有取消，则PutOn成功
	 */
	PutOnEffect(effect effect.Effect) bool
	/*
		注册一个处理函数,在实体 被施加效果前调用
		该处理函数如果返回true,则实体会继续进行后续操作，如果返回false,则会让实体跳过本次操作(并触发 OnCancelPutOn)
	 */
	OnBeforePutOnEffect(handler event.EventHandler) int64
	/*
		注册一个处理函数，在实体 被释放效果 之后调用
	 */
	OnAfterPutOnEffect (handler event.EventHandler) int64
	/*
		实体的 效果施放动作 被取消之后触发一次
	 */
	OnCancelPutOnEffect (handler event.EventHandler) int64


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
	OnBeforeRemoveEffect(handler event.EventHandler) int64
	/*
		注册一个处理函数，在实体 被取消效果 之后调用
	 */
	OnAfterRemoveEffect (handler event.EventHandler) int64
	/*
		实体的 效果施放动作 被取消之后触发一次
	 */
	OnCancelRemoveEffect (handler event.EventHandler) int64
}