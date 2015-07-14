package maps

//代表地图上的一个点，是地图的最小组成单位
//可以理解为size为1,1的单元
type Block struct {
	position Position
	objectLayers []interface{} //表示上面放置的物体，从0层开始向上叠加
}
func(b *Block) getPosition() Position{
	return b.position
}
//设置block的位置
func(b *Block) setPosition(val Position) Position{
	if b.position == nil{
		b.position = Position{x:val.getX(),y:val.getY()}
	}else{
		b.position.setX(val.getX())
		b.position.setY(val.getY())
	}
}