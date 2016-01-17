package event
import "survive/server/logic/dataStructure"

//事件处理函数，第一个版本简单做，所以不支持异步的处理
type HandleFunc func(time *dataStructure.Time, contextParams ...interface{}) (isCancel bool,handleResult string)
/*
	基础的事件处理器

	这个处理器可以被外部初始化，并插入到各个实体的执行阶段里
 */
type EventHandler struct {

	//能够执行事件处理
	//返回是否取消实体店本次操作，以及操作结果
	//注意，只有在beforeXXX事件中返回 true的时候，才有可能取消后续操作
	//在不同的事件中，可能会有不同的事件参数需要处理
	Func HandleFunc
}

func NewEventHandler(fn HandleFunc) *EventHandler{
	return &EventHandler{
		Func:fn,
	}
}
