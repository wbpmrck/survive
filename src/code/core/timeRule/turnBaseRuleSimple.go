
package timeRule
//实现一个简单的回合制时间规则
//每接收到inPipe的一个脉冲，rule自身累加当前时间刻度
//当接收outPipe的一个时间请求时，rule判断agent获取到的时间刻度，是否小于自身刻度，是的话可以执行
//agent必须在获取到时间片的时候，更新自己的时间刻度