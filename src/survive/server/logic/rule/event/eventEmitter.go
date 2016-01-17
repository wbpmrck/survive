package event

/*
	* 支持通用的事件订阅、发布模型
	* 在订阅事件的时候，对EventHandler进行构造和调用
 */

type EventEmitter interface {
	On() int64 //订阅事件，并返回订阅句柄，日后可以用这个句柄取消订阅，放置内存泄露
	Once() int64 //只订阅1次
	Emit() //发射事件
	Off(token int64) bool //取消订阅事件，并返回是否取消成功(找到并成功取消订阅才返回true)

}

/*
	实现一个简单通用的事件收集、发射器
 */
type EventEmitterBase struct {


}