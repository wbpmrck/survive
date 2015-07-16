package agent
import (
	"code/core/timeRule"
	"fmt"
)

//表示agent之间互相通信的消息
type AgentMessage struct {
	messageType uint32 //消息类型
	messageBody interface{} //消息内容
	from *Agenter
	to *Agenter
	responseChan chan *AgentMessage //消息处理方，通过这个chan回复应答
}
func (msg *AgentMessage) GetResponseChan()chan *AgentMessage{
	return msg.responseChan
}
func (msg *AgentMessage) GetMessageType() uint32{
	return msg.messageType
}
func (msg *AgentMessage) GetMessageBody() interface{}{
	return msg.messageBody
}
func (msg *AgentMessage) GetSrc() *Agenter{
	return msg.from
}
func (msg *AgentMessage) GetTarget() *Agenter{
	return msg.to
}

func (msg AgentMessage) String()string{
	return fmt.Sprintf("type: %v,body: %v",msg.messageType,msg.messageBody)
}

type TimeSliceMessage struct {
	ts *timeRule.TimeSlice//消息内容:时间片大小
	from *Agenter
	to *Agenter
	responseChan chan *AgentMessage //消息处理方，通过这个chan回复应答
}

func (msg *TimeSliceMessage) GetResponseChan()chan *AgentMessage{
	return msg.responseChan
}
func (msg *TimeSliceMessage) GetTimeSlice() *timeRule.TimeSlice{
	return msg.ts
}
func (msg *TimeSliceMessage) GetSrc() *Agenter{
	return msg.from
}
func (msg *TimeSliceMessage) GetTarget() *Agenter{
	return msg.to
}

func (msg TimeSliceMessage) String()string{
	return fmt.Sprintf("timeSlice.duration = %v",msg.ts.GetDuration())
}

/**
	创建请求类：普通消息
	此类消息会自动创建一个cap为0的responseChan,用于接受对方的回答
 */
func CreateRequestAgentMessage(messageType uint32,messageBody interface{},from *Agenter, to *Agenter) *AgentMessage{
	msg := &AgentMessage{
		messageType:messageType,
		messageBody:messageBody,
		from:from,
		to:to,
		responseChan:MakeChanWithAgentMessage(0),
	}
	return msg
}

/**
	创建请求类：时间片消息
 */
func CreateRequestTSMessage(ts *timeRule.TimeSlice,from *Agenter, to *Agenter) *TimeSliceMessage{
	msg := &TimeSliceMessage{
		ts:ts,
		from:from,
		to:to,
		responseChan:MakeChanWithAgentMessage(0),
	}
	return msg
}

/**
	创建通知类：时间片消息
 */
func CreateNotifyTSMessage(ts *timeRule.TimeSlice,from *Agenter, to *Agenter) *TimeSliceMessage{
	msg := &TimeSliceMessage{
		ts:ts,
		from:from,
		to:to,
		responseChan:nil,
	}
	return msg
}

/**
	创建通知类消息
	此类消息不含responseChan(=nil),通常用于单方面通知对方的时候使用
 */
func CreateNotifyAgentMessage(messageType uint32,messageBody interface{},from *Agenter, to *Agenter) *AgentMessage{
	msg := &AgentMessage{
		messageType:messageType,
		messageBody:messageBody,
		from:from,
		to:to,
		responseChan:nil,
	}
	return msg
}

/**
	辅助方法，快速创建一个上限为cap的chan *AgentMessage
 */
func MakeChanWithAgentMessage(cap int) chan *AgentMessage {
	channel := make(chan *AgentMessage, cap)
	return channel
}


//2015-07-16：废弃了，对双工通信设想的过于复杂
///**
//	定义消息管道对
//	每一次外界请求与agent通信，都会生成这样的一对pipe
// */
//type MessagePipePair struct {
//	InPipe,OutPipe chan *AgentMessage
//	//结构未来被存放在连接池的下标，在移除的时候增加查找效率
//	IndexInPool uint64
//	//表示这对管道是否在被使用【主要由于动态select比较难写且耗费性能，所以要初始化规定的连接池，并根据此标识来控制
//	// 不会有2个外部对象同时使用这组通道】
//	IsUsing bool
//}
//
//func NewMessagePipePair(indexInPool uint64,isUsing bool) *MessagePipePair{
//	p := &MessagePipePair{InPipe:make(chan *AgentMessage),OutPipe:make(chan *AgentMessage),indexInPool:indexInPool,isUsing:isUsing}
//	return p
//}