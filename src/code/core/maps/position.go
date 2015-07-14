package maps

//表示一个地图上的位置
type Position struct {
	x,y int64
}
func(p *Position) getX() int64{
	return p.x
}
func(p *Position) setX(val int64) {
	p.x = val
}
func(p *Position) getY() int64{
	return p.y
}

func(p *Position) setY(val int64) {
	p.y = val
}