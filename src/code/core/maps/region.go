package maps

/**
	代表一个地图块
	一些游戏中，可能需要按块加载地图
	地图块是1个虚拟概念，通过改变地图块的参数，可以改变地图块所引用的block范围
 */
type Region struct {
	size Size //区域的大小
	leftTop Position //区域的起始坐标(左上角)
}
func(r *Region) getSize() Size{
	return r.size
}
func(r *Region) setSize(val Size){
	if r.size == nil{
		r.size = &Size{width:val.getWidth(),height:val.getHeight()}
	}else{
		r.size.setWidth(val.getWidth())
		r.size.setHeight(val.getHeight())
	}
}
func(r *Region) getLeftTop() Position{
	return r.leftTop
}
func(r *Region) setLeftTop(val Position){
	if r.leftTop == nil{
		r.leftTop = Position{x:val.getX(),y:val.getY()}
	}else{
		r.leftTop.setX(val.getX())
		r.leftTop.setY(val.getY())
	}
}
