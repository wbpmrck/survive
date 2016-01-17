
//技能是由若干个技能项构成
//每个技能项，是一个技能作用对象chooser，和一堆效果的集合（一个技能项，代表向一堆对象释放的效果）
//技能可以升级，要保存当前的技能等级、经验，总级数，每一级需要经验等信息
package skill
import (
	"survive/server/logic/character"
"survive/server/logic/skill/targetChoose"
)

//代表一个技能项
type SkillItem interface {
	SetTargetChooser(chooser targetChoose.TargetChooser) //设定技能的选择目标的策略
	Install(from *character.Character,targets []*character.Character) //install的过程，就是给每一个target创建effect对象，并PutOn的过程
	GetInfo() string //获取描述信息
	Update(args ...interface{}) //进行技能的生效处理(Update可能多次被调用)
	UnInstall() //技能结束处理，表示该技能释放完毕
}

//代表一个技能。
//一个技能可以含有若干个技能项
//技能的释放，就是对所有技能项目的激活
type Skill interface {
	Install(from *character.Character,targets []*character.Character)
	GetInfo() string //获取描述信息
	Update(args ...interface{}) //进行技能的生效处理(Update可能多次被调用)
	UnInstall() //技能结束处理，表示该技能释放完毕
}
