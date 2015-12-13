package internal

import (
	"survive/server/modules/game/internal/handler"
	"survive/server/utils/register"
)

func init() {
	// 注册消息处理函数
	register.RegistStringHandler(skeleton,"UserLogin", handler.UserLogin)
	register.RegistStringHandler(skeleton,"NewAgent", handler.NewAgent)
	register.RegistStringHandler(skeleton,"CloseAgent", handler.CloseAgent)
}
