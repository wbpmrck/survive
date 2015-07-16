package timeRule

//表示一个时间片

type TimeSlice struct {
	duration int64
}
func (ts *TimeSlice) GetDuration() int64{
	return ts.duration
}
func (ts *TimeSlice) SetDuration(d int64) {
	ts.duration = d
}

func NewTimeSlice(d int64) *TimeSlice{
	return &TimeSlice{duration:d}
}