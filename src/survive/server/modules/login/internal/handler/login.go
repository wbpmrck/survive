package handler
import (
"survive/server/msg"
"survive/server/logger"
"survive/server/modules/game"
"github.com/name5566/leaf/gate"
)

func Login(args []interface{}) {
	m := args[0].(*msg.Hello)
	a := args[1].(gate.Agent)

	//1个简单的登录逻辑验证，如果发送的accId不在配置文件配置的范围内，则登录失败
	//	if len(m.AccID) < gamedata.AccIDMin || len(m.AccID) > gamedata.AccIDMax {
	//		a.WriteMsg(&msg.S2C_Auth{Err: msg.S2C_Auth_AccIDInvalid})
	//		return
	//	}

	logger.GetLogger().Debugf("login prepare send: [%v]",m.Name)
	// login
	game.ChanRPC.Go("UserLogin", m,a)
	logger.GetLogger().Debugf("login after send: [%v]",m.Name)

	//	a.WriteMsg(&msg.H{Err: msg.S2C_Auth_OK})
}

