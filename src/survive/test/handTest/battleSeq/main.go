package main
import (
	"survive/server/logic/arena"
	"survive/server/logic/battle"
	"survive/server/logic/time"
	"survive/server/logic/dataStructure"
	systime "time"
	"survive/test/dataCreator"
	"fmt"
//	"survive/server/logic/dataStructure/attribute"
	"survive/server/logic/rule/event"
	"os"
"encoding/json"
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
	//rate 10,means 10X real world speed
	arena1 := arena.NewArena(500,20) //arena为了尽快完成比赛结果，时间流逝速度为20倍，每个时间片为500ms

	fieldLen := 900
	battle1 := battle.NewBattle(fieldLen,players...)
	battle1.Desc="test battle"
	arena1.AddBattle(battle1)

	//让 arena 可以接收时间片
	godTimeSrc.AppendReceiver(arena1)

	//添加双方 warriors (一方5个)
	warriors := dataCreator.GetWarriors(len(players)*2)


	//对第一个角色做特殊加强处理
	w1,w2 := warriors[0],warriors[1]

	w1.EachOpMoveDistance.GetValue().Add(1,nil)
	w1.GetAttr(attribute.AGI).GetValue().Add(490,nil)
	w1.GetAttr(attribute.VIT).GetValue().Add(290,nil)

	w2.GetAttr(attribute.AGI).GetValue().Add(-4,nil)
	w2.GetAttr(attribute.STR).GetValue().Add(-5,nil)
	w2.GetAttr(attribute.VIT).GetValue().Add(-5,nil)


	//对第二个角色做特殊削弱处理

	for i:=0;i<len(players)*2;i++{
		//前5个 属于player1
		if i<2{
			warriors[i].SetPlayer(players[0])
			battle1.AddWarrior(warriors[i],0)
		}else{
			//后5个，属于player2
			warriors[i].SetPlayer(players[1])
			battle1.AddWarrior(warriors[i],fieldLen-1)
		}
	}

	//订阅日志
	battle1.On(battle.EVENT_END,event.NewEventHandler(func (contextParams ...interface{}) (isCancel bool,handleResult interface{}){
		//写入文件
		userFile := "record.json"
		fout,err := os.Create(userFile)
		defer fout.Close()
		if err != nil {
			fmt.Println(userFile,err)
			return
		}
		b, err := json.Marshal(battle1.Report)
		if err != nil {
			fmt.Println("json err:", err)
		}

//		fmt.Println(string(b))

		fout.WriteString(string(b))

		fmt.Println("write over!")
		return
	}))

	//时间开始
	godTimeSrc.Begin()

	fmt.Println("-----Sleep---------")
	systime.Sleep(300*systime.Second)
}
