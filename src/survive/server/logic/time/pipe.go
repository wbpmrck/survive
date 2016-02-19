package time
import (
	"survive/server/logic/dataStructure"
	"time"
)
/*
	时间管道,可以从上一级拿到时间片，并转发给下一级

 */

type Pipe struct{
	Rate float32 //时间比例( 现实时间流失 * rate = 游戏世界流失 )
	Now dataStructure.Time //当前管道内的时刻
	GameTimeUnit dataStructure.TimeSpan //管道的输出时间片的长度单位
	TSRemain dataStructure.TimeSpan //管道目前还剩余的时间片
	Receivers []Receiver //管道时间片的下游

//	InChan  chan dataStructure.TimeSpan //管道的时间片来源
//	OutChan chan dataStructure.TimeSpan //对外部转发时间点chan
}
//实现接受者接口，管道被初始化当前时间。这么做可以保证整个系统的各个组件进行时间同步
func(self *Pipe) Init(now dataStructure.Time){
	self.Now = now
	for _,r := range self.Receivers{
		//分发给接受者
		r.Init(now)
	}
}

func(self *Pipe) PipeUp(upstream Producer){
	if upstream != nil{
		//把自己加入上游管道的接受者中
		upstream.AppendReceiver(self)
	}
}

//实现Producer接口
//管道虽然不自己产生数据，但是当他转发数据的时候，他对下游表现的像生产者
func(self *Pipe) AppendReceiver(rec Receiver){
	self.Receivers = append(self.Receivers,rec)
}
//接受一个时间片(实现了Receiver接口)
//将时间片按照自己的单位长度，分片交给下游
//注意每交一次，都更新自己的时间
func(self *Pipe) Receive(ts dataStructure.TimeSpan){
	//接到一个时间片，首先按照自己的倍率进行处理
//	ts.GameSpan = ts.RealSpan *  time.Duration(self.Rate) //得到自己认为“已经度过的游戏时间”
	ts.GameSpan = time.Duration(float32(ts.RealSpan) * self.Rate) //得到自己认为“已经度过的游戏时间”
//	ts.RealSpan = ts.RealSpan *  time.Duration(self.Rate) //得到自己认为“已经度过的现实时间”

	self.TSRemain = self.TSRemain.Add(ts)
	//只要还有剩余时间可以分发，就分发
	for self.TSRemain.GameSpan >= self.GameTimeUnit.GameSpan{
		//扣除剩余时间片
		self.TSRemain = self.TSRemain.Slice(self.GameTimeUnit)
		for _,r := range self.Receivers{
			//分发给接受者
			r.Receive(self.GameTimeUnit)
		}
		//更新当前时刻
		self.Now = self.Now.Elapse(self.GameTimeUnit)
	}
}

func NewPipe(gameTimeUnitInMS time.Duration,timeRate float32) *Pipe{
	return &Pipe{
		Rate:timeRate,
		GameTimeUnit:dataStructure.NewMilliSecondSpanWithGameSpan(gameTimeUnitInMS,timeRate),
		TSRemain:dataStructure.NewMilliSecondSpan(0,timeRate),
		Receivers:make([]Receiver,0),
	}
}