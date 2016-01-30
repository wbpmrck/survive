package player

type Player struct {
	Id string
	UserName,Password string
	VipLevel int
}
func NewPlayer(id,name,password string,vipLevel int) *Player{
	return &Player{
		Id:id,
		UserName:name,
		Password:password,
		VipLevel:vipLevel,
	}
}

