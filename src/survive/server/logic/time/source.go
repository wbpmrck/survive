package time
import (
	"survive/server/logic/dataStructure"
	"time"
	"survive/server/logger"
)

/*
	时间源,可以产生时间片，并转发给下一级
	source 需要实现生产者接口
	PS:目前只是单服务器，不考虑分布式系统的时间同步问题
	由于是单机，所以也不考虑用多个goroutine来处理时间片了
 */
const (
	STATE_RUNNING int = iota
	STATE_STOP
)



type Source struct{
	Now dataStructure.Time //当前管道内的时刻
 	//输出时间片的长度单位(这里的时间片，就是游戏主循环的一帧的间隔，由于有多级pipe的支持，无需为了最精确的游戏场景来设定
	//主循环的间隔。
	//建议的间隔时间是1s
	GameTimeUnit dataStructure.TimeSpan
	ticker *time.Ticker //计时器
	State int //时间源的状态
	Receivers []Receiver //下游接受者
}
//时间源开始不断的工作，产生一片一片时间片，并顺着接收者管道派发下去
func(self *Source) Begin(){
	self.State = STATE_RUNNING
	logger := logger.GetLogger()
	tickCount := 0 //记录生成的时间片数量
	printDuration := 10000 //每个多少次打印一下
	//建立一个计时器，计时器的轮休时间，就是单位时间片
	self.ticker = time.NewTicker(self.GameTimeUnit.RealSpan)
	logger.Infof("已创建帧间隔为 %v ms的ticker,准备进入时间生成主循环",self.GameTimeUnit.RealSpan / 1000000)

	go func() {
		for _ = range self.ticker.C {
			tickCount++
			if tickCount == printDuration{
				logger.Infof("已产生%v个时间片",tickCount)
				tickCount = 0
			}
			for _,r := range self.Receivers{
				//分发给接受者
				r.Receive(self.GameTimeUnit)
			}
		}
	}()
}
//实现Producer接口
//管道虽然不自己产生数据，但是当他转发数据的时候，他对下游表现的像生产者
func(self *Source) AppendReceiver(rec Receiver){
	self.Receivers = append(self.Receivers,rec)
}

func NewSource(now dataStructure.Time,gameTimeUnitInMs time.Duration,timeRate int) *Source{
	return &Source{
		Now:now,
		GameTimeUnit:dataStructure.NewMilliSecondSpan(gameTimeUnitInMs,timeRate),
		Receivers:make([]Receiver,0),
		State:STATE_STOP,
	}
}