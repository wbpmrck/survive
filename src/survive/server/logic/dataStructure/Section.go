package dataStructure
import "survive/server/logic/dataStructure/attribute"

//表示一个距离范围，可以形容攻击范围、作用范围等等
type Section struct {
	From,To *attribute.Attribute
}
