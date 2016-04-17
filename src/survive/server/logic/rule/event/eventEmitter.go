package event
//todo:TTL功能还未实现
const (
	ALL_CHAN string ="*"
)

/*
	* 支持通用的事件订阅、发布模型
	* 在订阅事件的时候，对EventHandler进行构造和调用
 */



//表示一个处理函数在事件处理宿主里的唯一标识，方便后续快速方便的取消订阅
type HandlerId struct {
	SlotKey string //事件所在的槽的名字
	IndexInSlot int //事件所在槽的处理函数索引地址
}

type EventEmitter interface {
	On(evtName string,handler *EventHandler) HandlerId //订阅事件，并返回订阅句柄，日后可以用这个句柄取消订阅，放置内存泄露
	Once(evtName string,handler *EventHandler) HandlerId //只订阅1次
	Emit(evtName string,params ...interface{}) []HandleResult //发射事件,并得到所有处理函数的处理结果列表
	Off(id HandlerId) bool //取消订阅事件，并返回是否取消成功(找到并成功取消订阅才返回true)
}

/*
	实现一个简单通用的事件收集、发射器
 */
type EventEmitterBase struct {
	slots map[string][]*EventHandler //事件槽 ：map.key是事件名 map.value是一个处理函数队列
}

//发射一个事件，并收集所有处理函数的返回信息
func(self *EventEmitterBase) Emit(evtName string,params ...interface{}) []HandleResult{
	//获取要发射的事件有多少处理函数
	slot,exist := self.slots[evtName]
	result := make([]HandleResult,0,len(slot))
	if exist{
		for _,handler:= range slot{
			isCancel,r := handler.Func(params ...)
			result = append(result,HandleResult{
				IsCancel:isCancel,
				HandleResult:r,
			})
		}
	}

	//再看有没有订阅了*的，有的话也调用
	subAll,existAll:= self.slots[ALL_CHAN]
	if existAll && subAll!= nil && len(subAll)>0{
		//有的话，修改输入参数，进行调用

		resultAll := make([]HandleResult,0,len(subAll))
		newParam := make([]interface{},0,len(params)+1)

		//修改后的输入参数，第一个参数是消息名
		newParam = append(newParam,evtName)
		newParam = append(newParam,params...)
		for _,handler:= range subAll{
			isCancel,r := handler.Func(newParam ...)
			resultAll = append(resultAll,HandleResult{
				IsCancel:isCancel,
				HandleResult:r,
			})
		}
		//把执行结果加到后面
		result = append(result,resultAll...)
	}

	return result
}

//订阅一个事件，只处理一次就被移除
func(self *EventEmitterBase) Once(evtName string,handler *EventHandler) HandlerId{
	handler.TTL =1
	return self.On(evtName,handler)
}

//订阅事件处理函数
func(self *EventEmitterBase) On(evtName string,handler *EventHandler) HandlerId{
	//如果该事件槽还未被订阅过，则创建事件槽
	_,exist := self.slots[evtName]
	if !exist{
		self.slots[evtName] = make([]*EventHandler,0)
	}
	//在槽中添加对应事件处理函数
	self.slots[evtName] = append(self.slots[evtName],handler)

	return HandlerId{
		SlotKey:evtName,
		IndexInSlot:len(self.slots[evtName])-1,
	}
}

//根据事件处理函数id,取消一个订阅
func(self *EventEmitterBase) Off(id HandlerId) bool {
	slot,exist:=self.slots[id.SlotKey]
	if exist{
		slot = append(slot[:id.IndexInSlot], slot[id.IndexInSlot+1:]...)
		return true
	}
	return false
}

func NewEventEmitter()*EventEmitterBase{
	return &EventEmitterBase{
		slots:make(map[string][]*EventHandler),
	}
}