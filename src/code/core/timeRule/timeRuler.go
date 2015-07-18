//
package timeRule

type TimeRuler interface {
	/**
	通常为被regionManager调用获取时间管道
	返回时间管道pipe
	timeRuler通过pipe，来告知agent,现在他的时间刻度可以往后走n个时间单位
	 */
	Start(agent interface{}) (pipe <-chan int)

	/**
	通常被外部管理者调用
	可以通过这个pipe,来人工触发timeRuler执行一次time的elapse
	在即时模式下，这个外部触发可能是一个脉冲
	 */
	GetTimeInPipe() (pipe chan<- int)
}
