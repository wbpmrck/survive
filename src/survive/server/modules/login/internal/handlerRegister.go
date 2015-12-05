package internal

import (
	"reflect"
"survive/server/msg"
	"survive/server/modules/login/internal/handler"
)

func regist(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	regist(&msg.Hello{}, handler.Login) //注册消息处理函数
}

