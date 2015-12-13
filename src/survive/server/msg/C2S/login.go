package C2S

import (
	"survive/server/msg/base"
)

// 一个结构体定义了一个 JSON 消息的格式
// 消息名为 Login
type Login struct {
	base.IdentifiableMsg
	UserName,Password string
}
