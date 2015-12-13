package msg

import (
	"github.com/name5566/leaf/network/json"
	"survive/server/msg/C2S"
	"survive/server/msg/common"
)

// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
var Processor = json.NewProcessor()

func init() {
	// 这里我们注册了一个 JSON 消息 Hello
	Processor.Register(&C2S.Login{})
	Processor.Register(&common.SimpleResponseMsg{})
}
