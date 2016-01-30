package main
import (
	"survive/server/logic/arena"
	"survive/server/logic/battle"
	"survive/server/logic/player"
	"survive/server/logic/time"
	"survive/server/logic/dataStructure"
	systime "time"
)


func main(){

	//暂且创建的时间速度都是和现实一样
	godTimeSrc := time.NewSource(dataStructure.Time{
		RealTime:systime.Time{},
		GameTime:systime.Time{},
	},
	1000,1)

	player1 := player.NewPlayer("user1","userName1","1",1)
	player2 := player.NewPlayer("user2","userName2","2",1)
	//create arena,time tick is 1s
	//rate 1,means like real world
	arena1 := arena.NewArena(1000,1)
	battle1 := battle.NewBattle(10,player1,player2)
	arena1.AddBattle(battle1)

	godTimeSrc.AppendReceiver(arena1)

	godTimeSrc.Begin()

	systime.Sleep(10*systime.Second)
}
