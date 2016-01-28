package attribute

/*
	描述一个带有属性的实体。
	符合本接口的实体，可以被进行属性的读取和写入、修正

	注意，这里定义的属性携带者，携带的是属性接口，而不是具体的属性。
	具体获取到的属性可能是普通属性，也可能是计算属性，但是它们表现出来的行为都是一样的
 */
type AttributeCarrier interface {
	//直接获取这个对象携带的所有属性(通过指针，可以直接修改其属性)
	GetAllAttr() map[string]AttributeLike
	//获取这个单位的某个属性(通过指针，可以直接修改其属性)
	GetAttr(attrKey string) AttributeLike
	//给单位重新设置某个属性
//	SetAttr(attrKey string, attr AttributeLike)
	SetAttr(attr AttributeLike)
	//给单位添加一个属性(如果该属性已经存在，就不做任何操作，并返回false)，新增成功返回true
//	AddAttr(attrKey string, attr AttributeLike) bool
	AddAttr(attr AttributeLike) bool
	//判断对象是否存在一个属性
	HasAttr(name string) bool
}

/*
	一个基本的属性携带者的实现。
	其他复杂对象如果没有特殊需求的，可以组合这个类型来满足 AttributeCarrier 接口的需求
 */
type AttributeCarrierBase struct {
	attributes map[string]AttributeLike
}
//返回是否存在某属性
func(self *AttributeCarrierBase) HasAttr(name string) bool{
	_,exist:= self.attributes[name]
	return exist
}
//直接获取这个对象携带的所有属性(通过指针，可以直接修改其属性)
func(self *AttributeCarrierBase) GetAllAttr() map[string]AttributeLike{
	return self.attributes
}
//获取这个单位的某个属性(通过指针，可以直接修改其属性)
func(self *AttributeCarrierBase) GetAttr(attrKey string) AttributeLike{
	return self.attributes[attrKey]
}
//给单位重新设置某个属性
func(self *AttributeCarrierBase) SetAttr(attr  AttributeLike){
	if attr != nil{
		self.attributes[attr.GetName()] = attr
	}
}
//给单位添加一个属性(如果该属性已经存在，就不做任何操作，并返回false)，新增成功返回true
func(self *AttributeCarrierBase) AddAttr(attr AttributeLike) bool{
	if attr != nil{
		_,exist := self.attributes[attr.GetName()]
		if !exist{
			self.attributes[attr.GetName()] = attr
			return true
		}
	}
	return false
}

//创建一个空的属性携带者
func NewAttributeCarrier() *AttributeCarrierBase{
	return &AttributeCarrierBase{
		attributes:make(map[string]AttributeLike),
	}
}
//创建一个带初始属性的携带者
func NewAttributeCarrierWithAttr(initAttr map[string]AttributeLike) *AttributeCarrierBase{
	if initAttr == nil{
		initAttr =make(map[string]AttributeLike)
	}
	return &AttributeCarrierBase{
		attributes:initAttr,
	}
}