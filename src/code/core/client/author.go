package client

type Author interface {
	//通过输入数据subject,判断是否是安全的client请求
	//如果鉴权通过，author应该从服务器端获取一些info返回，这些info有助于client构造用户相关的业务信息
	DoAuth(subject interface{}) (authorized bool,player *Player,returnCode string)
}