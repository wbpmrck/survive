package attribute

//表示一个值
//游戏里的数值都有“原始值”和“修正值”
//原始值表示数值本来的属性，修正值是由于装备，技能等影响的附加属性
type Value struct {
	raw float64 //原始值
	additionVal float64 //当前累计修正值（状态、技能、光环等）
	additionPercent float32 //当前累计修正幅度
}

func (self *Value) GetRaw() float64 {
	return self.raw
}
func (self *Value) SetRaw(v float64) {
	self.raw = v
}
//注意，计算方式是，被乘系数为“原始值”+“增加值”，所以“增加值”也是很重要的
func (self *Value) Get() float64 {
	return float64((self.raw + self.additionVal) * float64(1 + self.additionPercent))
}

func (self *Value) Add(v float64) {
	self.additionVal += v
}

func (self *Value) AddByPercent(v float32) {
	self.additionPercent += v
}

