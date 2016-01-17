package rule
import "survive/server/logic/rule/event"
/*
	代表一个可以移动的单位
 */

type Moveable interface {
	event.EventEmitter
	/*
		注册一个处理函数,在实体 准备移动前调用
		该处理函数如果返回true,则实体会继续进行后续操作，如果返回false,则会让实体跳过本次操作(并触发 OnCancelMove)
	 */
	OnBeforeMove(handler event.EventHandler) int64
	/*
		注册一个处理函数，在实体移动的时候调用
	 */
	OnMove (handler event.EventHandler) int64

	/*
		注册一个处理函数，在实体移动之后调用
	 */
	OnAfterMove (handler event.EventHandler) int64
	/*
		实体的移动操作被取消之后触发一次
	 */
	OnCancelMove (handler event.EventHandler) int64
}