package rule
import (
	"survive/server/logic/dataStructure"
	"survive/server/logic/rule/event"
)

/*
	timeVariable 表示一个可随着时间变化，进行自我改变的实体
	这样的实体会随着游戏时间变化而做出主动的行为
 */
type TimeVariable interface {
	event.EventEmitter
	/*
		注册一个处理函数,在实体被更新之前调用
		该处理函数如果返回true,则实体会继续进行后续更新操作，如果返回false,则会让实体跳过本次更新(并触发 OnCancelUpdate)
	 */
	OnBeforeUpdate (handler event.EventHandler) int64
	/*
		注册一个处理函数，在实体被更新的时候调用
		该函数可以操作实体的任何内容
	 */
	OnUpdate (handler event.EventHandler) int64

	/*
		注册一个处理函数，在实体被更新之后调用
	 */
	OnAfterUpdate (handler event.EventHandler) int64
	/*
		实体的更新操作被取消之后触发一次，可以用来做一些兜底操作
	 */
	OnCancelUpdate (handler event.EventHandler) int64

	/*
		update的入口，由时间管理者输入当前时间
	 */
	Update(time *dataStructure.Time)
}