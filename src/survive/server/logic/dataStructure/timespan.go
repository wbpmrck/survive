package dataStructure
import "time"

//var rate = 10 //游戏时间速度/现实时间速度 比率
/*
	表示一个游戏时间段
 */
type TimeSpan struct {
	RealSpan time.Duration //这段时间长度，在现实生活中的长度
	GameSpan time.Duration //这段时间长度，在游戏中的长度
}

//时间片增加
func(self TimeSpan) Add(ts TimeSpan)(added TimeSpan){
	self.RealSpan += ts.RealSpan
	self.GameSpan += ts.GameSpan
	return self
}

//从时间片上切下一块(减少)
func(self TimeSpan) Slice(ts TimeSpan) (remain TimeSpan){
	self.RealSpan -= ts.RealSpan
	self.GameSpan -= ts.GameSpan
	return self
}

//输入一个现实世界中流逝的时间(毫秒为单位)，产生一个包含游戏流逝时间的时间片
func NewMilliSecondSpan(realSpanInMS time.Duration,rate int) TimeSpan{
	realSpanInMicro := realSpanInMS * 1000000
	return TimeSpan{
		RealSpan:realSpanInMicro,
		GameSpan:realSpanInMicro*rate,
	}
}