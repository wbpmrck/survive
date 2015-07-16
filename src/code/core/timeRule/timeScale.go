package timeRule

//表示一个时间刻度
//todo:暂时没想好时间刻度如何表示，先采用和时间片一样的
type TimeScale struct {
	duration int64
}
func (ts *TimeScale) GetDuration() int64{
	return ts.duration
}
func (ts *TimeScale) SetDuration(d int64) {
	ts.duration = d
}

func NewTimeScale(d int64) *TimeScale{
	return &TimeScale{duration:d}
}