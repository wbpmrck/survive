package attribute
import "survive/server/logic/rule/event"
/*
	这是一种特殊的attr,它们除了带有自身的加成之外，还主要依靠其他属性来计算自己的"原始值"
	当其他属性改变的时候，他自己的"原始值"也会跟着改变
 */

//原始值计算函数，返回计算好的原始值
//type RawComputer func(dependencies map[string]AttributeLike) (computedRaw float64,[]AttributeLike)
type RawComputer func(dependencies... AttributeLike) (computedRaw float64)

type ComputedAttribute struct {
	*Attribute //计算属性首先自己是一个属性
	dependencies []AttributeLike //保存它所依赖的其他属性的列表
	rawComputer RawComputer //保存该属性的最终值的计算方法
	cachedRaw bool //表示当前的raw是否可以直接使用(如果是false,代表当前的raw所依赖的外部属性被修改了，在获取的时候需要重新计算一下)
}
//获取计算属性的值(注意很有意思的一点,GetValue被覆盖了，因为计算属性的属性获取方式不一样)
//每次获取计算属性的值的时候，都会重新根据依赖属性计算一下自身的原始值
func(self *ComputedAttribute) GetValue() *Value{
	//如果当前数据有脏标记
	if !self.cachedRaw{
		//重新计算raw
		self.val.SetRaw(self.rawComputer(self.dependencies))
	}
	return self.val
}

func NewComputedAttribute(name,desc string,rawValue float64,computer RawComputer,dependencies ...AttributeLike) *ComputedAttribute{
	c:= &ComputedAttribute{
		Attribute:NewAttribute(name,desc,rawValue),
		dependencies:dependencies,
		rawComputer:computer,
		cachedRaw:false,
	}
	if dependencies!=nil && len(dependencies)>0{
		for _,a:= range dependencies{
			a.On(EVENT_VALUE_CHANGED,
				//订阅所依赖属性的变化事件，第一个参数是变化后的新值
				event.NewEventHandler(func (contextParams ...interface{}) (isCancel bool,handleResult string){
					c.cachedRaw = false //只要有依赖项变化，就修改自己的脏标记
			}))
		}
	}
	return c
}
