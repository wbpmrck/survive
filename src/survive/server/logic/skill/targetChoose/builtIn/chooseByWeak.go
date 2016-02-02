package builtIn
import (
	"sort"
	"survive/server/logic/skill/effect"
	"survive/server/logic/battle"
	"fmt"
)

//定义interface{},并实现sort.Interface接口的三个方法
type CharacterSlice []*battle.Warrior

func (c CharacterSlice) Len() int {
	return len(c)
}
func (c CharacterSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c CharacterSlice) Less(i, j int) bool {
	return c[i].HP < c[j].HP
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
func(self *ChooseByWeak) Choose(from interface{},stepName string,params ...interface{})(targets []effect.EffectCarrier,error bool) {
	from, ok := from.(battle.Warrior)
	if ok{
		battleIn := from.BattleIn
		enemy := battleIn.GetEnemy(from.Player)
		if len(enemy) > 0 {
			//这里为了简单，就取第一个对手势力
			firstEnemy := enemy[0]
			n := self.num
			//取最弱的n个,先按照hp排序，最少的放前面
			characters := (CharacterSlice)(battleIn.PlayerCharactersList[firstEnemy.Id])

			if !sort.IsSorted(characters) {
				sort.Sort(characters)
			}

			//然后返回前n个
			warriors := battleIn.PlayerCharactersList[firstEnemy.Id][0:n]

			targets = make([]effect.EffectCarrier,0,len(warriors))

			for id,_ := range warriors{
				targets = append(targets,warriors[id])
			}

			error = false
			return
		}else {
			return nil, nil
		}
	}else{
		return nil,true
	}
}