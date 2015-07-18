package client

type PlayerType uint16

const (
	GM PlayerType = iota //0
	NormalPlayer	//1
)

//describe a real-world player
type Player struct {
	id,name string
	playerType PlayerType
}