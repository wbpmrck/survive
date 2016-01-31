package main
import (
	"survive/server/logic/arena"
	"survive/server/logic/battle"
	"survive/server/logic/time"
	"survive/server/logic/dataStructure"
	systime "time"
	"survive/test/dataCreator"
	"fmt"
	"survive/server/logic/dataStructure/attribute"
)


func main(){

	fmt.Println("---------test begin---------")

	//暂且创建的时间速度都是和现实一样
	godTimeSrc := time.NewSource(dataStructure.Time{
		RealTime:systime.Time{},
		GameTime:systime.Time{},
	},
	1000,1)

	//创建 player
	players := dataCreator.GetPlayers(2)
	//create arena,time tick is 1s
	//rate 1,means like real world
	arena1 := arena.NewArena(1000,1)
	battle1 := battle.NewBattle(10,players...)
	battle1.Desc="test battle"
	arena1.AddBattle(battle1)

	//让 arena 可以接收时间片
	godTimeSrc.AppendReceiver(arena1)

	//添加双方 warriors (一方5个)
	warriors := dataCreator.GetWarriors(len(players)*5)


	for i:=0;i<len(players)*5;i++{
		//前5个 属于player1
		if i<5{
			warriors[i].SetPlayer(players[0])
		}else{
			//后5个，属于player2
			warriors[i].SetPlayer(players[1])
		}
		battle1.AddWarrior(warriors[i])
	}
	//微调，让第5个 最快行动
	warriors[4].GetAttr(attribute.AGI).GetValue().Add(10,nil)


	//时间开始
	godTimeSrc.Begin()
	fmt.Println("-----时间开始---------")
	systime.Sleep(4*systime.Second)

	fmt.Println("-----降低敏捷---------")
	warriors[4].GetAttr(attribute.AGI).GetValue().Add(-11,nil)

	fmt.Println("-----Sleep---------")
	systime.Sleep(5*systime.Second)

	fmt.Println("-----再增加敏捷---------")
	warriors[4].GetAttr(attribute.AGI).GetValue().Add(21,nil)

	fmt.Println("-----Sleep---------")
	systime.Sleep(5*systime.Second)
}
