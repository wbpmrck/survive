package attribute

//表示一个属性值
//属性是类似力量、敏捷、身高等这种，描述一个人物特点的数值。
//属性也可能产生对“能力值”的修正影响
type Value struct {
	raw int64 //原始值
	additionVal int64 //当前累计修正值（状态、技能、光环等）
	additionPercent float32 //当前累计修正幅度
}

func (self *Value) GetRaw() int64 {
	return self.raw
}
func (self *Value) SetRaw(v int64) {
	self.raw = v
}
//注意，计算方式是，被乘系数为“原始值”+“增加值”，所以“增加值”也是很重要的
func (self *Value) GetValue() int64 {
	return int64((self.raw + self.additionVal) * (1 + self.additionPercent))
}

func (self *Value) ChangeAddition(v int64) {
	self.additionVal += v
}

func (self *Value) ChangeAdditionPercent(v float32) {
	self.additionPercent += v
}

