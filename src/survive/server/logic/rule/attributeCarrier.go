package rule
import "survive/server/logic/dataStructure/attribute"

/*
	描述一个带有属性的实体。
	符合本接口的实体，可以被进行属性的读取和写入、修正
 */

type AttributeCarrier interface {
	//直接获取这个对象携带的所有属性(通过指针，可以直接修改其属性)
	GetAllAttr() map[string]*attribute.Attribute
	//获取这个单位的某个属性(通过指针，可以直接修改其属性)
	GetAttr(attrKey string) *attribute.Attribute
	//给单位重新设置某个属性
	SetAttr(attrKey string, *attribute.Attribute)
	//给单位添加一个属性(如果该属性已经存在，就不做任何操作，并返回false)，新增成功返回true
	AddAttr(attrKey string, *attribute.Attribute) bool
}
