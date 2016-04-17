package attribute
import (
//	"fmt"
	"time"
	"fmt"
)
/*
	这是一种特殊的attr,它们除了带有自身的加成之外，还可以随着时间变化而恢复自己的值
	这些属性通常有上限 、 恢复速度
	PS:恢复属性恢复的是Value的raw,而不是对属性进行加成。
 */


type RegeneratedAttribute struct {
	*Attribute                           //恢复属性首先自己是一个属性
	Max                    *Value        //该属性值恢复的上限
	RegenerateCountPerTick *Value        //每次恢复属性的多少
	tickDuration           time.Duration //表示2次恢复间隔的时间长短
	totalDuration          time.Duration //表示已经积累的恢复时间
}

func(self *RegeneratedAttribute)String()string{
	return fmt.Sprintf("[%s|%s = %v/%v,回复速度:%v/%v]",self.name,self.desc,self.val,self.Max,self.RegenerateCountPerTick,self.tickDuration)
}
func(self *RegeneratedAttribute)TimeAcquire(ts time.Duration){
	//如果当前属性满,不处理。不满才可以处理
	if self.val.GetRaw() < self.Max.Get(){
		self.totalDuration += ts //先新增积累的时间
		//判断时间是否达到恢复间隔
		if self.totalDuration >= self.tickDuration {
			//判断回复量
			cur := self.val.GetRaw()
			reg := self.RegenerateCountPerTick.Get()
			max := self.Max.Get()
			if cur+reg>=max{
				self.val.SetRaw(max)
			}else{
				self.val.AddRaw(reg)
			}

			self.totalDuration -= self.tickDuration
		}
	}
}
func NewRegeneratedAttribute(name,desc string,rawValue float64,max,regCountPerTick *Value,tickDuration time.Duration) *RegeneratedAttribute {
	c:= &RegeneratedAttribute{
		Attribute:NewAttribute(name,desc,rawValue),
		Max:max,
		RegenerateCountPerTick:regCountPerTick,
		tickDuration:tickDuration,
		totalDuration:0,
	}

	return c
}
