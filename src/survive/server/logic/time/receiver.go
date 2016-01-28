package time
import "survive/server/logic/dataStructure"
/*
	receiver 接口表示一个可以接受时间片并处理的对象
	注意接受者只能被动接收，所以他是暴露方法给“生产者”调用
 */
type Receiver interface {
	Receive(ts dataStructure.TimeSpan)
	Init(now dataStructure.Time) //所有的接受者的时间，都是由上游决定的，当上游开始Init的时候，要确保所有的下游都Init完成
}
