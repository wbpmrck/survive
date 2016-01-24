
package skill
import (
	"survive/server/logic/skill/effect"
	"survive/server/logic/skill/targetChoose"
	"fmt"
)
/*
	关于技能的设计：
	1、技能其实是效果、和作用对象的封装
	2、技能负责产生效果，但它无法撤消效果。效果是真正实现技能的核心，它自己决定自己何时被消灭，当然，也可以被主动消灭
	3、技能还有一个主要的目的，是管理效果的参数，比如技能可以有等级和熟练度的概念，借此来影响内部效果的配置值
 */

//代表一个技能
//一个技能，是由若干个效果项构成的
type Skill struct {
	Carrier SkillCarrier //这个技能所属的对象(是对象的技能)
	Level int //当前技能等级
	Exp int //当前技能熟练度
	Items []*SkillItem
	Name string
}
//技能安装
func(self *Skill) Install(from SkillCarrier){
	self.Carrier = from
	//初始化每个效果
	if self.Items != nil && len(self.Items)>0{
		for _,item := range self.Items{
			item.Install(from)
		}
	}
}
//获取技能信息(暂时就返回名字+技能列表)
func(self *Skill) GetInfo() string{
	return fmt.Sprintf("[%v]%v",self.Name,self.Items)
}
//获取描述信息
func(self *Skill) AddSkillItem(item *SkillItem) {
	if item != nil{
		if self.Items != nil{
			self.Items = make([]*SkillItem,0)
		}
		self.Items = append(self.Items,item)
	}
}


//表示一个技能项
type SkillItem struct {
	SkillParent *Skill //指向包含自己的技能项
	PluginStep Step //技能项插入技能持有者动作的时机(skillItem通过注册skillCarrier的这个事件,来决定何时调用effect)
	Chooser targetChoose.TargetChooser //在固定时机的基础上，chooser决定效果的作用对象
	EffectItems []effect.Effect
}
//获取描述信息
func(self *SkillItem) String() string{
	return fmt.Sprintf("作用时机:%v \r\n作用对象:%v \r\n效果项:%v",self.PluginStep.StepDesc,self.Chooser.String(),self.EffectItems)
}
func(self *SkillItem) Install(){
	//在技能持有者身上订阅触发技能的事件
	if self.SkillParent.Carrier != nil{
		//在指定的阶段
		stepName:=self.PluginStep.StepName
		self.SkillParent.Carrier.On(stepName,func(contextParams ...interface{}) (isCancel bool,handleResult string){
			//找到指定的作用对象
			targets,err := self.Chooser.Choose(self.SkillParent.Carrier,stepName,contextParams...)
			if err == nil && targets!= nil && len(targets)>0{
				for _,target := range targets{
					//对每一个效果，使用PutOn进行触发
					for _,effectItem := range self.EffectItems{
						effectItem.PutOn(self.SkillParent.Carrier,target)
					}
				}
			}
		})
	}
}
//定义技能项的释放时机
//后续应该初始化若干个释放时机以供选择
type Step struct {
	StepName string //插入时机名称
	StepDesc string //插入时机描述
}
