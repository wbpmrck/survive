package event
import "survive/server/logic/dataStructure"

/*
	* 支持通用的事件订阅、发布模型
	* 在订阅事件的时候，对EventHandler进行构造和调用
 */



//表示一个处理函数在事件处理宿主里的唯一标识，方便后续快速方便的取消订阅
type HandlerId struct {
	SlotKey string //事件所在的槽的名字
	IndexInSlot int64 //事件所在槽的处理函数索引地址
}

type EventEmitter interface {
	On(evtName string,handler *EventHandler) HandlerId //订阅事件，并返回订阅句柄，日后可以用这个句柄取消订阅，放置内存泄露
	Once(evtName string,handler *EventHandler) HandlerId //只订阅1次
	Emit(time *dataStructure.Time,evtName string,params ...interface{}) []HandleResult //发射事件,并得到所有处理函数的处理结果列表
	Off(id HandlerId) bool //取消订阅事件，并返回是否取消成功(找到并成功取消订阅才返回true)
}

/*
	实现一个简单通用的事件收集、发射器
 */
type EventEmitterBase struct {
	Slots map[string][]*EventHandler //事件槽 ：map.key是事件名 map.value是一个处理函数队列
}
func(self *EventEmitterBase) Emit(time *dataStructure.Time,evtName string,params ...interface{}) []HandleResult{
	//获取要发射的事件有多少处理函数
	slot,exist := self.Slots[evtName]
	result := make([]HandleResult)
	if exist{
		for _,handler:= range slot{
			result = append(result,handler.Func(time,params...))
		}
	}
	return result
}
func(self *EventEmitterBase) Once(evtName string,handler *EventHandler) HandlerId{
	handler.TTL =1
	return self.On(evtName,handler)
}
func(self *EventEmitterBase) On(evtName string,handler *EventHandler) HandlerId{
	slot,exist := self.Slots[evtName]
	if !exist{
		self.Slots[evtName] = make([]*EventHandler)
		slot = self.Slots[evtName]
	}
	slot = append(slot,handler)

	return HandlerId{
		SlotKey:evtName,
		IndexInSlot:len(slot)-1,
	}
}
func(self *EventEmitterBase) Off(id HandlerId) bool {
	slot,exist:=self.Slots[id.SlotKey]
	if exist{
		slot = append(slot[:id.IndexInSlot], slot[id.IndexInSlot+1:]...)
	}
}