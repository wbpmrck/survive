package handler
import (
"survive/server/msg/C2S"
"survive/server/msg/common"
"github.com/name5566/leaf/gate"
"survive/server/logger"
	"time"
	"fmt"
)

var seed=0;

func UserLogin(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*C2S.Login)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	logger.GetLogger().Debugf("game receive: [%v]", m)

	//模拟等待1s
	time.Sleep(3*time.Second)
	// 给发送者回应一个 Hello 消息
	seed++
	a.WriteMsg(common.NewSimpleResponseMsg(m,"0000","",C2S.Login{
		UserName: fmt.Sprintf("hello,client_%v",seed),
	}))
//	a.WriteMsg(&C2S.Login{
//		UserName: fmt.Sprintf("hello,client_%v",seed),
//	})
	logger.GetLogger().Debugf("game after send: [%v]", fmt.Sprintf("hello,client_%v",seed))
}
