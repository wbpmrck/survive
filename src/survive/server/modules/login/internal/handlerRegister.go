package internal

import (
	"survive/server/msg/C2S"
	"survive/server/modules/login/internal/handler"
	"survive/server/utils/register"
)

func init() {
	register.RegistMsgHandler(skeleton,&C2S.Login{}, handler.Login) //注册消息处理函数
}

