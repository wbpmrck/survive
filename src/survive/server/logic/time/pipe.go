package time
/*
	piper 定义了与时间相关的重要接口
	piper 目前考虑分为2类：
		activePiper ：主动的时间管道，这种时间管道有能力自己制造时间片，并pipe给下一级，但是这类管道无法从上一级得到时间片
		passivePiper: 被动的时间管道,
 */