package base

type Identifiable interface {
	GetId() string
}
//表示一个可被唯一标识的消息
type IdentifiableMsg struct {
	MsgId string
}
func (self *IdentifiableMsg) GetId() string{
	return self.MsgId
}

