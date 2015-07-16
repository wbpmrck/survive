package agent

//表示agent之间互相通信的消息
type AgentMessage struct {
	messageType uint32 //消息类型
	messageBody interface{} //消息内容
	from *Agenter
	to *Agenter
	responseChan chan *AgentMessage //消息处理方，通过这个chan回复应答
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