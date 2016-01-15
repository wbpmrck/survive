package attribute

//描述一类实体，他们可以被当作属性看待
type AttributeLike interface {
	GetName() string //可以获取属性名
	SetName(name string) //可以设置属性名

	GetDesc() string //可以获取属性描述
	SetDesc(desc string) //可以设置属性描述

	GetValue() *Value //可以获取值
	SetValue(val *Value) //可以设置值
}

//表示一个属性值
//属性是类似力量、敏捷、身高等这种，描述一个人物特点的数值。
//属性也可能产生对“能力值”的修正影响
type Attribute struct {
	name,desc string //属性名、属性描述
	val *Value //属性值
}
func(self *Attribute) GetName() string{
	return self.name
}

func(self *Attribute) SetName(name string) {
	self.name = name
}
func(self *Attribute) GetDesc() string{
	return self.desc
}

func(self *Attribute) SetDesc(desc string) {
	self.desc = desc
}
func(self *Attribute) GetValue() *Value{
	return self.val
}
func(self *Attribute) SetValue(val *Value) {
	self.val = val
}
//创建一个属性(名称、描述、原始值)
func NewAttribute(name,desc string,rawValue float64) *Attribute{
	return &Attribute{
		name:name,
		desc:desc,
		val:&Value{
			raw:rawValue,
			additionVal:0,
			additionPercent:0,
		},
	}
}