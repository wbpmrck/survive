package attribute
import (
	"fmt"
	"survive/server/logic/rule/event"
)
const(
	EVENT_VALUE_CHANGED string ="EVENT_VALUE_CHANGED" //对应处理函数参数:func(newVal *Value)
)

const(
	STR string = "str"
	AGI string = "agi"
	INT string = "int"
	VIT string = "vit"
	LUCK string = "luk"
	AWARE string = "awr"
	UNDERSTAND string = "und"

)
//描述一类实体，他们可以被当作属性看待
//有了这一层抽象，普通属性和“计算属性” 就可以被一视同仁了
type AttributeLike interface {
	event.EventEmitter
	GetName() string //可以获取属性名
	SetName(name string) //可以设置属性名

	GetDesc() string //可以获取属性描述
	SetDesc(desc string) //可以设置属性描述

	GetValue() *Value //可以获取值
	SetValue(val *Value) //可以设置值
	String() string
}

//表示一个属性值
//属性是类似力量、敏捷、身高等这种，描述一个人物特点的数值。
//属性也可能产生对“能力值”的修正影响
type Attribute struct {
	*event.EventEmitterBase //属性需要用事件的方式通知外部自己被修改了
	name,desc string //属性名、属性描述
	val *Value //属性值
}
func(self *Attribute) String() string{
	return fmt.Sprintf("[%s|%s = %v]",self.name,self.desc,self.val)
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
	 a:=&Attribute{
		name:name,
		desc:desc,
		val:&Value{
			raw:rawValue,
			totalAdditionVal:0,
			totalAdditionPercent:0,
			adders:make([]*Adder,0),
		},
	 	EventEmitterBase:event.NewEventEmitter(),

	 }
	a.val.attRef = a
	return a
}