package character

//表示一个属性值
//属性是类似力量、敏捷、身高等这种，描述一个人物特点的数值。
//属性也可能产生对“能力值”的修正影响
type Attribute struct {
	Name,Desc string //属性名、属性描述
	Value int //属性值
}
