
package skill
import (
	"survive/server/logic/skill/effect"
	"survive/server/logic/skill/targetChoose"
	"fmt"
	"survive/server/logic/rule/event"
	"survive/server/logic/math"
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
			item.Install()
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

const MaxProbability int = 1000 + 1 //表示最大可能性的概率值(1000/1000)
//表示一个技能项
type SkillItem struct {
	SkillParent *Skill //指向包含自己的技能项
	Probability int //技能项释放的概率(千分位)
	PluginStep Step //技能项插入技能持有者动作的时机(skillItem通过注册skillCarrier的这个事件,来决定何时调用effect)
	Chooser targetChoose.TargetChooser //在固定时机的基础上，chooser决定效果的作用对象
	EffectItems []effect.Effect
}
//获取描述信息
func(self *SkillItem) String() string{
	return fmt.Sprintf("作用时机:%v \r\n作用对象:%v \r\n效果项:%v",self.PluginStep.StepDesc,self.Chooser.String(),self.EffectItems)
}
func(self *SkillItem) AddEffect( e effect.Effect){
	if self.EffectItems == nil{
		self.EffectItems = make([]effect.Effect,0)
	}
	self.EffectItems = append(self.EffectItems,e)
}
func(self *SkillItem) Install(){
	//在技能持有者身上订阅触发技能的事件
	if self.SkillParent.Carrier != nil{
		//在指定的阶段
		stepName:=self.PluginStep.StepName

		//技能项在指定阶段决定是否进行效果释放处理。
		//返回信息
			//类型:*SkillItem
		self.SkillParent.Carrier.On(stepName,event.NewEventHandler(func(contextParams ...interface{}) (isCancel bool,handleResult interface{}){

			//先看概率是否释放
			thisTimeProbability := math.NextRandomInt(MaxProbability)
			if self.Probability >= thisTimeProbability{
				//概率命中了，还要看是否能找到作用对象

				//再找到指定的作用对象
				targets,err := self.Chooser.Choose(self.SkillParent.Carrier,stepName,contextParams...)
				if !err && targets!= nil && len(targets)>0{

					//有作用对象，还要看pre 阶段是否被取消
					cancel := false

					//触发技能持有者的技能释放事件(前)
					//处理函数将得到一个参数：*SkillItem
					onBeforeResults:= self.SkillParent.Carrier.Emit(EVENT_BEFORE_SKILL_ITEM_RELEASE,self)
					for _,r := range onBeforeResults{
						if r.IsCancel{
							cancel = true
							//执行cancel 阶段[技能被取消释放]
							self.SkillParent.Carrier.Emit(EVENT_CANCEL_SKILL_ITEM_RELEASE,self)
							break
						}
					}
					//没取消，进入释放
					if !cancel{
						self.SkillParent.Carrier.Emit(EVENT_SKILL_ITEM_RELEASE,self)
						for _,target := range targets{
							//对每一个效果，使用PutOn进行触发
							for _,effectItem := range self.EffectItems{
								effectItem.PutOn(self.SkillParent.Carrier,target)
							}
						}
						handleResult = self //释放成功
						//释放结束，执行 after 阶段
						self.SkillParent.Carrier.Emit(EVENT_AFTER_SKILL_ITEM_RELEASE,self)
					}

				}
			}

			return
		}))
	}
}
//定义技能项的释放时机
//后续应该初始化若干个释放时机以供选择
type Step struct {
	StepName string //插入时机名称 PS:EVENT_WARRIOR_SKILL_CHOOSE 阶段名，就表示大部分主动释放类技能应该plugin的阶段
	StepDesc string //插入时机描述
}

//创建一个技能，创建后可以添加技能项，然后使用install来调用
func NewSkill(name string,level int,exp int) *Skill{
	return &Skill{
		Name:name,
		Level:level,
		Exp:exp,
		Items:make([]*SkillItem,0),
	}
}

//创建一个技能项
func NewSkillItem(parent *Skill,plugStep Step, chooser targetChoose.TargetChooser,probability int) *SkillItem{
	return &SkillItem{
		Probability:probability,
		SkillParent:parent,
		PluginStep:plugStep,
		Chooser:chooser,
		EffectItems:make([]effect.Effect,0),
	}
}