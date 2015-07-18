package client
import "github.com/funny/link"

type ClientType uint16

const (
	MobileNative ClientType = iota //0
	MobileBrowser	//1
	PCBrowser		//2
)

type Client struct {
	//客户端session
	endPointSession link.Session
	//客户端类型
	clientType ClientType
	//client使用的鉴权策略
	author Author
	//所属玩家
	player *Player
}
/**
	进行客户端登录鉴权，设置player信息并返回
 */
func (c *Client) Login(subject interface{}) (success bool,returnCode string){
	success,p,returnCode := c.author.DoAuth(subject)
	if success{
		c.player = p
	}
	return
}
func NewClient(endPointSession link.Session,author Author,t ClientType) *Client{
	return &Client{
		clientType:t,
		author:author,
		endPointSession:t,
		player:nil,
	}
}
