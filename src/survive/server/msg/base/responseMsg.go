package base

//表示这是一个回复另外一个消息的消息
type ResponseMsg struct {
	IdentifiableMsg
	RespTo IdentifiableMsg
}

