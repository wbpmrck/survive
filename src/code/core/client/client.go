package client

type ClientType uint16

const (
	MobileNative ClientType = iota //0
	MobileBrowser	//1
	PCBrowser		//2
)

type Client struct {
	//客户端ip地址
	endpointIP string
	//客户端类型
	clientType ClientType
	//所属玩家
	player Player
}
