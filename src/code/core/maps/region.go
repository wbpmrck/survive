package maps

/**
	代表一个地图块
	地图块一般由一个goroutine来管理，块之间的数据不互通
	地图块由哪个goroutine管理，由worldMap决定
	坐标系：x向右+,y向下+
 */
type Region struct {
	size Size //区域的大小
	leftTop Position //区域的起始坐标(左上角)(相对于全局worldMap)
	data [][]*Block //地图数据
}

//判断region内部是否包含该全局坐标
func(r *Region) isContainGlobalPos(globalPos *Position) bool{
	var left,top,right,bottom int =0,0,0,0
	left = r.leftTop.getX()
	top = r.leftTop.getY()
	right = left + r.getSize().getWidth()
	bottom = top + r.getSize().getHeight()

	gx,gy := globalPos.getX(),globalPos.getY()

	return gx >= left && gx <= right && gy >= top && gy <= bottom
}
//把全局坐标转换为region内部的坐标
func(r *Region) transformGlobalPosToLocal(globalPos *Position) (localPos *Position){
	if r.isContainGlobalPos(globalPos){
		localPos = &Position{x:globalPos.getX()-r.getLeftTop().getX(),y:globalPos.getY()-r.getLeftTop().getY()}
		return localPos
	}else{
		return nil
	}
}
/**
	获取全局坐标(x,y)决定的Block
 */
func(r *Region) getDataAtGlobal(globalPos *Position) *Block{
	//如果该点在所管辖范围内，则返回
	if r.data != nil && r.isContainGlobalPos(globalPos){
		return r.getDataAtLocal(r.transformGlobalPosToLocal(globalPos))
	}else{
		return nil
	}
}
/**
	获取区域内x,y坐标的Block
 */
func(r *Region) getDataAtLocal(localPos *Position) *Block{
	if r.data != nil && localPos.getX()<r.size.getWidth() && localPos.getY()<r.size.getHeight(){
		return r.data[localPos.getX()][localPos.getY()]
	}else{
		return nil
	}
}



func(r *Region) getData() [][]*Block{
	return r.data
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
