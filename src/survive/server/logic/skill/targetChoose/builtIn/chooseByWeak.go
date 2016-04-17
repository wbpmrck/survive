package builtIn
import (
	"sort"
	"survive/server/logic/skill/effect"
	"survive/server/logic/battle"
	"fmt"
)

//定义interface{},并实现sort.Interface接口的三个方法
type characterSliceLessByHP []*battle.Warrior

func (c characterSliceLessByHP) Len() int {
	return len(c)
}
func (c characterSliceLessByHP) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c characterSliceLessByHP) Less(i, j int) bool {
	return c[i].HP.GetValue().Get() < c[j].HP.GetValue().Get()
}

type ChooseByWeak struct {
	num int //选择几个对象
}
func(self *ChooseByWeak) GetName() string{
	return "ChooseByWeak"
}
func(self *ChooseByWeak) String() string{
	return fmt.Sprintf("对象:技能使用者的敌对势力中最虚弱的%v个角色",self.num)
}
//可以设置选择的对象人数
func(self *ChooseByWeak) Config(args ...interface{}){
	if len(args)>0{
		self.num = args[0].(int)
		if self.num<0{
			self.num = 1 //最少选择1个对象
		}
	}
}
//选择器开始进行对象选择
//按照敌人虚弱程度进行选择，通常
func(self *ChooseByWeak) Choose(fromWho interface{},stepName string,params ...interface{})(targets []effect.EffectCarrier,error bool) {

	error = true
	from, ok := fromWho.(*battle.Warrior)
	if ok{
		//首先获得攻击者的敌人
		enemy := from.BattleIn.GetEnemyWarriors(from.Player)
		if len(enemy) > 0 {
			n := self.num
			//取最弱的n个,先按照hp排序，最少的放前面
			characters := (characterSliceLessByHP)(enemy)

			if !sort.IsSorted(characters) {
				sort.Sort(characters)
			}

			//然后返回前n个
			warriors := characters[0:n]

			targets = make([]effect.EffectCarrier,0,len(warriors))

			for id,_ := range warriors{
				targets = append(targets,warriors[id])
			}

			error = false
		}
	}
	return
}