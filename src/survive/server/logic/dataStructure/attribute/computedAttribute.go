package attribute
/*
	这是一种特殊的attr,它们除了带有自身的加成之外，还主要依靠其他属性来计算自己的"原始值"
	当其他属性改变的时候，他自己的"原始值"也会跟着改变
 */

//原始值计算函数，返回计算好的原始值
type RawComputer func(dependencies map[string]AttributeLike) float64

type ComputedAttribute struct {
	*Attribute //计算属性首先自己是一个属性
	dependencies map[string]AttributeLike //保存它所依赖的其他属性的列表
	rawComputer RawComputer //保存该属性的最终值的计算方法
}
//获取计算属性的值(注意很有意思的一点,GetValue被覆盖了，因为计算属性的属性获取方式不一样)
//每次获取计算属性的值的时候，都会重新根据依赖属性计算一下自身的原始值
func(self *ComputedAttribute) GetValue() *Value{
	self.val.SetRaw(self.rawComputer(self.dependencies))
	return self.val
}

func NewComputedAttribute(name,desc string,rawValue float64,dependencies map[string]AttributeLike,computer RawComputer) *ComputedAttribute{
	return &ComputedAttribute{
		Attribute:NewAttribute(name,desc,rawValue),
		dependencies:dependencies,
		rawComputer:computer,
	}
}
