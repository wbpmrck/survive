package dataCreator
import "survive/server/logic/player"

var seed int = 0
func GetPlayer() *player.Player{
	seed++
	player1 := player.NewPlayer("user"+string(seed),"userName"+string(seed),string(seed),seed)
	return player1
}
func GetPlayers(int n) []*player.Player{

	players := make([]*player.Player,n)
	for i:=1;i<=n;i++{
		players = append(GetPlayer())
	}
	return players
}
