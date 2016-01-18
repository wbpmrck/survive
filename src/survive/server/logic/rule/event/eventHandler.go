package event
import "survive/server/logic/dataStructure"

//表示一个处理函数的处理结果
type HandleResult struct {
	IsCancel bool //是否进行下一个阶段的处理（这里的下一阶段，是由eventEmitter决定的）
	HandleResult string
}
//事件处理函数，第一个版本简单做，所以不支持异步的处理
type HandleFunc func(time *dataStructure.Time, contextParams ...interface{}) (isCancel bool,handleResult string)
/*
	基础的事件处理器

	这个处理器可以被外部初始化，并插入到各个实体的执行阶段里
 */
type EventHandler struct {

	//handler最多被调用的次数。-1表示无限。
	//EventEmitter 应该检查此标记，如果TTL减少到0就应该移除这个监听
	TTL int
	//能够执行事件处理
	//返回是否取消实体店本次操作，以及操作结果
	//注意，只有在beforeXXX事件中返回 true的时候，才有可能取消后续操作
	//在不同的事件中，可能会有不同的事件参数需要处理
	Func HandleFunc
}

//创建一个无限TTL的事件处理函数(只能主动通过Off方法来取消订阅)
func NewEventHandler(fn HandleFunc) *EventHandler{
	return &EventHandler{
		TTL:-1,
		Func:fn,
	}
}
//创建一个指定TTL的事件处理函数
func NewEventHandlerWithTTL(ttl int,fn HandleFunc) *EventHandler{
	return &EventHandler{
		TTL:ttl,
		Func:fn,
	}
}
