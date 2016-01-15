package builtIn
import (
	"survive/server/logic/character"
	"survive/server/logic/skill/effect/bases"
	"survive/server/logic/skill/effect"
	"fmt"
	"time"
)

//表示对属性进行修正（按数值修正）的一种效果
type AttributeModify struct {
	bases.EffectBase

	attrKey string//影响的属性id
	attrName string//影响的属性名称
	amount int64 //影响到属性修正值
}

//效果施加
func(self *AttributeModify) PutOn(from, target *character.Character){
	self.EffectBase.PutOn(from,target)
	//增加属性
	self.Target.Attributes[self.attrKey].Val.ChangeAddition(self.amount)
}

//效果更新
//todo:思考何时调用Update,思考是否要传入固定的参数，比如当前阶段(判断攻击命中，判定技能命中等)
func(self *AttributeModify) Update(args ...interface{}){

}

//效果移除
func(self *AttributeModify) Remove(){
	//减少属性
	self.Target.Attributes[self.attrKey].Val.ChangeAddition(-self.amount)

	self.EffectBase.Remove()
}

//配置修正值
func(self *AttributeModify) Config(args ...interface{}){
	if len(args>0){
		self.amount = args[0].(int) //参数0：属性修正值
	}else{
		self.amount = 0
	}
}
//显示效果信息
func(self *AttributeModify) GetInfo() string{
	operator := "+";
	if self.amount <0{
		operator="-"
	}
	//显示样例：力量 + 1
	return fmt.Sprintf("%v %s %v",self.attrName,operator,self.amount)
}

func init(){
	effect.RegisterFactory("AttributeModify",func() *Effect{
		return new(AttributeModify)
	})
}