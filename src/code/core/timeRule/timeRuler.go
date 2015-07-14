//define the base time controller of the game
package timeRule

type TimeRuler interface {
	/**
	获取时间管道：返回时间管道pipe
	timeRuler通过pipe，来告知agent,现在他的时间刻度可以往后走n个时间单位
	 */
	GetTimePipe(agent interface{}) (pipe <-chan int)
}
