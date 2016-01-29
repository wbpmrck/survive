package time

/*
	表示一个可以制造时间片的对象
 */
type Producer interface  {
	//生产者可以为多个接受者提供数据
	AppendReceiver(rec Receiver)
}
