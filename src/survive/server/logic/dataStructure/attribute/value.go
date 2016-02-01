package attribute
import "fmt"


//表示一个数据修正的发出方
type Adder struct {
							//表示这个对象的引用
	Ref             interface{}
	AdditionVal     float64 //当前累计修正值（状态、技能、光环等）
	AdditionPercent float32 //当前累计修正幅度
}
//表示一个值
//游戏里的数值都有“原始值”和“修正值”
//原始值表示数值本来的属性，修正值是由于装备，技能等影响的附加属性
type Value struct {
	attRef AttributeLike //持有这个值的属性的引用
	raw                  float64 //原始值
	totalAdditionVal     float64 //当前累计修正值（状态、技能、光环等）
	totalAdditionPercent float32 //当前累计修正幅度
	adders []*Adder //保存该值当前正被哪些对象修正(方便这些对象回退对值的修正)
}
func (self *Value) String() string{
	return fmt.Sprintf("{%6.3f,%6.3f,%6.3f => %6.3f}",self.raw,self.totalAdditionVal,self.totalAdditionPercent,self.Get())
}
//根据传入的修改者，在本数值项上，撤消它所作的更改 .如果找到并且撤消更改成功，返回adder
func (self *Value) UndoAllAddBy(who interface{}) *Adder{
	i,adder := self.GetAdder(who)
	if i!=-1 && adder != nil{

//		fmt.Printf("adder %v",adder)

		self.totalAdditionVal -= adder.AdditionVal
		self.totalAdditionPercent -= adder.AdditionPercent
		//删除这个adder
		self.adders = append(self.adders[:i],self.adders[i+1:]...)
		//发布属性变化事件
		if self.attRef != nil{
			self.attRef.Emit(EVENT_VALUE_CHANGED,self.attRef)
		}
	}
	return adder
}

//根据对象引用，获取对这个值的修改情况。没有的话返回空
func (self *Value) GetAdder( who interface{}) (index int,adder *Adder){
//	fmt.Printf("find %v in %v",who,self.adders)
	for i,_:= range self.adders{
		if self.adders[i].Ref == who{
			return i,self.adders[i]
		}
	}
	return -1,nil
}

func (self *Value) GetRaw() float64 {
	return self.raw
}
func (self *Value) SetRaw(v float64) {
	self.raw = v
	//发布属性变化事件
	if self.attRef != nil{
		self.attRef.Emit(EVENT_VALUE_CHANGED,self)
	}
}
//直接修改原始值(不需要进行值加成，不需要跟踪修改者，一次性修改)
func (self *Value) AddRaw(v float64) {
	self.raw += v
	//发布属性变化事件
	if self.attRef != nil{
		self.attRef.Emit(EVENT_VALUE_CHANGED,self)
	}
}
//注意，计算方式可以有2种，目前这一种，对属性修正类的技能更友好，不容易淘汰
func (self *Value) Get() float64 {
	return float64((self.raw + self.totalAdditionVal) * float64(1 + self.totalAdditionPercent))
//	return float64(self.raw* float64(1 + self.totalAdditionPercent)+ self.totalAdditionVal)
}


//对数值进行加成修正，并跟踪修改者，可以做回退
func (self *Value) Add(v float64,byWho interface{}) {
	self.totalAdditionVal += v
	//查看当前这个对象是否已经修改了自己
	_,adder := self.GetAdder(byWho)
	if adder == nil{
		//如果没有，创建一个新的adder,记录下修改数据
		adder = &Adder{
			Ref:byWho,
			AdditionVal:v,
			AdditionPercent:0.0,
		}
		self.adders = append(self.adders,adder)
	}else{
		//否则累加修改数据
		adder.AdditionVal +=v
	}
	fmt.Printf("self.attRef:%v \n",self.attRef)
	//发布属性变化事件
	if self.attRef != nil{
		fmt.Printf("changed:%v \n",self)
		self.attRef.Emit(EVENT_VALUE_CHANGED,self)
	}

}


//对数值进行加成修正，并跟踪修改者，可以做回退
func (self *Value) AddByPercent(v float32,byWho interface{}) {
	self.totalAdditionPercent += v
	//查看当前这个对象是否已经修改了自己
	_,adder := self.GetAdder(byWho)
	if adder == nil{
		//如果没有，创建一个新的adder,记录下修改数据
		adder = &Adder{
			Ref:byWho,
			AdditionVal:0.0,
			AdditionPercent:v,
		}
		self.adders = append(self.adders,adder)
	}else{
		//否则累加修改数据
		adder.AdditionPercent +=v
	}
	//发布属性变化事件
	if self.attRef != nil{
		self.attRef.Emit(EVENT_VALUE_CHANGED,self)
	}
}

//----------下面是之前的实现---------------
//
////表示一个值
////游戏里的数值都有“原始值”和“修正值”
////原始值表示数值本来的属性，修正值是由于装备，技能等影响的附加属性
//type Value struct {
//	raw float64 //原始值
//	additionVal float64 //当前累计修正值（状态、技能、光环等）
//	additionPercent float32 //当前累计修正幅度
//}
//
//func (self *Value) GetRaw() float64 {
//	return self.raw
//}
//func (self *Value) SetRaw(v float64) {
//	self.raw = v
//}
////注意，计算方式是，被乘系数为“原始值”+“增加值”，所以“增加值”也是很重要的
//func (self *Value) Get() float64 {
//	return float64((self.raw + self.additionVal) * float64(1 + self.additionPercent))
//}
//
//func (self *Value) Add(v float64) {
//	self.additionVal += v
//}
//
//func (self *Value) AddByPercent(v float32) {
//	self.additionPercent += v
//}
//
