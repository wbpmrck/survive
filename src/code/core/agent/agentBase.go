package agent
import (
	"code/core/timeRule"
	"time"
//	"errors"
//	"fmt"
	"code/utils"
)


type AgentStatus uint32

const (
	STATUS_NONE AgentStatus = iota
	STATUS_ACTIVATE
	STATUS_PAUSE
)
//能够和agent建立的连接数上限，默认是4个，也即同时一个agent只能与4个外部agent通信【主要是因为动态的chan数组的select监听实现起来很麻烦，而且官方不推荐】
//const MAX_CONN_COUNT uint64 = 4

/**
	agent 的逻辑处理接口，自定义的agent如果要使用agentBase的机制，就必须实现这个接口
	一定要注意，这里的所有逻辑处理，务必要在内存中进行，否则会占据过多的agent轮询goroutine的时间
 */
type AgentLogicHandler interface {
	//handle timeSlice event
	HandleTimeSlice(ts *TimeSliceMessage)

	//handle message event
	HandleMessage(msg *AgentMessage)
}


//定义基础的agent
type AgentBase struct {
	//唯一标识
	identity string
	//类型
	agentType AgentType
	//名字
	name string


	//当前所在的时间刻度
	currentTimeScale *timeRule.TimeScale
	//当前状态
	status AgentStatus
	//上一次状态变化的时间
	lastStatusChangeTime time.Time


	//agent的逻辑处理者
	logicHandler AgentLogicHandler



	//和当前agent建立联系的实体数量
//	connectionsCount uint64

//	和当前agent建立联系的实体所使用的管道对【包括消息通道、管理通道、时间片通道】
//	connections [] *MessagePipePair

	inMessagePipe chan *AgentMessage
	messagePipeSize uint64

	//管理者和agent通信的专用通道
//	managerPipes *MessagePipePair
	inManageCommandPipe chan *AgentMessage
	managePipeSize uint64

	//时间片管道。agent从这个管道，获取行动时间片
	//在一个时间片内，会告知agent获得了多少行动时间
	timeInChan chan *TimeSliceMessage
}

/**
	获取agent 状态
 */
func(a *AgentBase) GetStatus() AgentStatus{
	return a.status
}
/**
	设置agent 状态(private)
 */
func(a *AgentBase) setStatusAndLastChangeTime(s AgentStatus) {
	a.status = s
	a.lastStatusChangeTime = time.Now()
}
/**
	获取agent唯一标识
 */
func(a *AgentBase) GetIdentity() string{
	return a.identity
}
/**
	获取agentType
 */
func(a *AgentBase) GetAgentType() AgentType{
	return a.agentType
}

/**
	获取agent name
 */
func(a *AgentBase) GetName() string{
	return a.name
}
/**
	获取当前agent所处的时间刻度
 */
func(a *AgentBase) GetTimeScale() timeRule.TimeScale{
	return *a.currentTimeScale
}
/**
	获取管理者专属的消息通道
 */
func (a *AgentBase) GetManagerPipe()(pipe chan <- *AgentMessage){
	return a.inManageCommandPipe
}

///**
//	初始化所有的消息通道池
// */
//func (a *AgentBase) InitConnectionPipesPool(){
//	a.connections = make([] *MessagePipePair,MAX_CONN_COUNT,MAX_CONN_COUNT)
//	for idx,_ := range a.connections{
//		//创建新的未使用的管道对，放入连接池
//		a.connections[idx]=NewMessagePipePair(idx,false)
//	}
//}

///**
//	回收之前获取到的通信管道
// */
//func (a *AgentBase) RecycleMessagePipes(pipePairId uint64){
//	if pipePairId<0 || pipePairId>len(a.connections){
//		panic(errors.New(fmt.Sprintf("RecycleMessagePipes error,pipePairId is %v,connections pool's len is %v ",
//			pipePairId,len(a.connections))))
//	}else{
//		//重置标记
//		a.connections[pipePairId].IsUsing = false
//		//减少count(为下一次分配做准备)
//		a.connectionsCount -=1
//	}
//}


/**
获取agent的消息处理通道入口
 */
func(a *AgentBase) GetMessagePipe()(pipe chan <- *AgentMessage){
	return a.inMessagePipe
}
/**
	启动agent
	agent启动之后，返回自己的时间片chan给外部
 */
func(a *AgentBase) Start() chan <-*TimeSliceMessage{

	//设置状态为激活的
	a.setStatusAndLastChangeTime(STATUS_ACTIVATE)

	//开启3个 goroutine,分别进行【被动消息】、【时间片】、【管理消息管道】的轮询
	go a.startTimeChanPolling()
	go a.startMessageChanPolling()
	go a.startManagerMessageChanPolling()
	return a.timeInChan
}

/**
	开始时间片管道的轮询
 */
func (a *AgentBase) startTimeChanPolling(){
	for{
		timeSlice := <-a.timeInChan
		if a.logicHandler != nil{
			a.logicHandler.HandleTimeSlice(timeSlice)
		}
	}
}

/**
	开始普通消息管道的轮询
 */
func (a *AgentBase) startMessageChanPolling(){
	for{
		msg := <-a.inMessagePipe
		if a.logicHandler != nil{
			a.logicHandler.HandleMessage(msg)
		}
	}
}
/**
	开始Manager消息管道的轮询
 */
func (a *AgentBase) startManagerMessageChanPolling(){
	for{
		msg := <-a.inManageCommandPipe
		if a.logicHandler != nil{
			a.logicHandler.HandleMessage(msg)
		}
	}
}



/**
	创建一个 AgentBase
 */
func CreateAgentBase(name string,logicHandler AgentLogicHandler,managePipeSize,messagePipeSize uint64) *AgentBase{
	//todo:创建的时候，需要指定类型、名字,logichandler 等，做一些初始化操作
	agent := &AgentBase{
		identity:utils.GetGuid(),
		name : name,
		agentType:DefaultAgent,

		status:STATUS_NONE,
		currentTimeScale:timeRule.NewTimeScale(0),
		lastStatusChangeTime:time.Now(),
		logicHandler:logicHandler,
		messagePipeSize:messagePipeSize,
		managePipeSize:managePipeSize,
		inMessagePipe:make(chan *AgentMessage,messagePipeSize),
		inManageCommandPipe:make(chan *AgentMessage,managePipeSize),
		timeInChan:make(chan *TimeSliceMessage),
	}
	return agent
}