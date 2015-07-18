package agent
import "code/core/timeRule"

//agent类型
type AgentType uint32

//update 2015-07-18 去掉agentType,通过不同的实现就决定了不同的行为，没有标识的必要了
//const (
//	DefaultAgent AgentType = iota //0
//)

//所有agent必须实现的接口
type Agenter interface {

	//获取Id
	GetIdentity() string
//	GetAgentType() AgentType

	GetName() string
	//获取当前agent所处的时间刻度
	GetTimeScale() timeRule.TimeScale
	/**
		返回一个管道，该管道对于外部只写。可以让agent接收一个时间片s
		该方法应该只被调用一次，并且在agent内部创建2个goroutine：
		一个负责轮询时间片chan,一个负责轮询所有的对外通道组
	 */
	Start()(s chan <- *TimeSliceMessage)
	/**
		获取agent的消息处理通道入口
	 */
	GetMessagePipe()(pipe chan <- *AgentMessage)

	//获取管理者通道入口，该通道往往只用于系统级控制消息的发送【比如GM直接kill一个boss之类的】
	GetManagerPipe()(pipe chan <- *AgentMessage)
}