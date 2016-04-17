package dataStructure
import (
	"survive/server/logic/dataStructure/attribute"
	"fmt"
)

type SectionValueRange struct {
	From,To int //这里的range,代表在battleField上的具体位置
}
//返回一个点到这个距离范围的差值，如果刚好落在里面，就返回0
//第一个参数为 -1，则表示目标在位置区间的左侧
//第一个参数为 0，则表示目标在位置区间的中间
//第一个参数为 1，则表示目标在位置区间的右侧
func(self *SectionValueRange)DistanceTo(pos int) (side int,dist int){
	if pos < self.From {
		//目标位置在范围的 左侧
		side= -1
		dist =self.From - pos

	}else if pos >= self.From && pos <= self.To{
		//目标位置在范围的 内部
		side = 0
		dist = 0
	}else if pos > self.To{
		//目标位置在范围的 右侧
		side =1
		dist=pos - self.To
	}
	fmt.Printf("范围[%v,%v] 与目标[%v]的距离:{%v,%v} \n",self.From,self.To,pos,side,dist)
	return
}

//表示一个距离范围，可以形容攻击范围、作用范围等等
type Section struct {
	From,To *attribute.Attribute
}
//将该范围在1维直线上的某点进行双向发散[0,n],返回在2个方向上的范围列表
func(self *Section) RangeFrom(from int) (resultRange []*SectionValueRange){
	resultRange = make([]*SectionValueRange,2)
	thisFrom := int(self.From.GetValue().Get())
	thisTo := int(self.To.GetValue().Get())
	leftRange := &SectionValueRange{
		From:from-thisTo,
		To:from-thisFrom,
	}
	rightRange := &SectionValueRange{
		From:from + thisFrom,
		To:from + thisTo,
	}
	resultRange[0] = leftRange
	resultRange[1] = rightRange

	return
}

func NewSection(name,desc string,from,to float64) *Section{
	return &Section{
		From:attribute.NewAttribute(name+"-from",desc+"起",from),
		To:attribute.NewAttribute(name+"-to",desc+"止",to),
	}
}