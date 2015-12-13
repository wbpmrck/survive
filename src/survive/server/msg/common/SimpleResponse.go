package common

import (
	"survive/server/msg/base"
	"fmt"
)

//表示一个简单的响应类消息，含有一个返回码，返回描述，和data字段
type SimpleResponseMsg struct {
	base.ResponseMsg
	RCode,RDesc string //返回码，返回描述
	Data interface{}
}

//创建一个新的反馈消息
func NewSimpleResponseMsg(responseTo base.Identifiable,code,desc string,data interface{}) *SimpleResponseMsg {
	r := new(SimpleResponseMsg)
	requestId := responseTo.GetId()
	newId := fmt.Sprintf("r_%s",requestId)
	r.MsgId = newId //反馈消息的编号
	r.RespTo.MsgId = requestId //反馈消息对应的请求消息编号
	r.RCode = code
	r.RDesc = desc
	r.Data = data
	return r
}
