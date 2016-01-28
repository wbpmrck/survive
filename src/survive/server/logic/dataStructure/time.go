package dataStructure
import "time"

/*
	表示一个游戏时刻
 */
type Time struct {
	RealTime time.Time
	GameTime time.Time
}
//让时间流逝一段，返回最新的时间
func(self Time) Elapse(ts TimeSpan) Time{
	self.RealTime = self.RealTime.Add(ts.RealSpan)
	self.GameTime = self.GameTime.Add(ts.GameSpan)
	return self
}
