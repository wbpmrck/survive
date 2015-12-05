package internal

import (
	"survive/server/modules/game/internal/handler"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	regist("UserLogin", handler.UserLogin)
	regist("NewAgent", handler.NewAgent)
	regist("CloseAgent", handler.CloseAgent)
}

func regist(name string, h interface{}) {
	skeleton.RegisterChanRPC(name, h)
}
