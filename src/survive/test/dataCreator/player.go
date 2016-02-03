package dataCreator
import (
	"survive/server/logic/player"
	"strconv"
)

var seed_player int = 0
func GetPlayer() *player.Player{
	seed_player++
	player1 := player.NewPlayer("user"+strconv.FormatInt(int64(seed_player),10),"userName"+strconv.FormatInt(int64(seed_player),10),strconv.FormatInt(int64(seed_player),10), seed_player)
	return player1
}
func GetPlayers(n int) []*player.Player{

	players := make([]*player.Player,0,n)
	for i:=1;i<=n;i++{
		players = append(players,GetPlayer())
	}
	return players
}
