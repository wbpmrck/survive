package arena
import (
	"survive/server/logic/time"
	systime "time"
	"survive/server/logic/battle"
	"survive/server/logic/rule/event"
	"fmt"
)
/*
	竞技场对象
	1、竞技场玩法是脱离游戏主时间循环的一种玩法，所以需要竞技场对象接收时间片，然后分发给竞技场里的battle
	2、参与竞技场的英雄，全部是玩家英雄的一份副本
	3、Arena 是一个时间管道，他接收到时间片，然后分发给每一个战场对象(战场对象实现Receiver接口，但是不继续派发，而是直接处理)
 */
type Arena struct {
	*time.Pipe //arena 是一个时间管道
	battles	[]*battle.Battle //当前正在进行的战斗列表
}
func(self *Arena) RemoveBattle(battleIndex int){
	self.battles = append(self.battles[:battleIndex],self.battles[battleIndex+1:]...)
}
func(self *Arena) AddBattle(bt *battle.Battle){
	self.AppendReceiver(bt)
	//如果battle已经结束，则移除出
	fmt.Printf("如果battle已经结束，则移除出 \n")
	bt.On(battle.EVENT_END,event.NewEventHandler(func (contextParams ...interface{}) (isCancel bool,handleResult string){

		fmt.Printf("len is %v \n",len(contextParams))
		endBattle := contextParams[0].(*battle.Battle)
		for i,b := range self.battles{
			if b == endBattle{
				self.RemoveBattle(i)
			}
		}
		return
	}))
	return
}

//创建arena:
//第一个参数：每一帧game span
//第二个参数: 时间流逝速度 1代表和现实一样
func NewArena(timeSpanInMS systime.Duration,timeRate int)*Arena{
	return &Arena{
		Pipe:time.NewPipe(timeSpanInMS,timeRate),
	}
}