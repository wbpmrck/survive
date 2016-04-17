package builtIn
import (
	"sort"
	"survive/server/logic/skill/effect"
	"survive/server/logic/battle"
	"fmt"
)

//定义interface{},并实现sort.Interface接口的三个方法
type characterSliceLessByDistance []*battle.Warrior

func (c characterSliceLessByDistance) Len() int {
	return len(c)
}
func (c characterSliceLessByDistance) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c characterSliceLessByDistance) Less(i, j int) bool {
	return c[i].SortSeed < c[j].SortSeed
}
type ChooseByAttackRegion struct {
	num int //选择几个对象(-1代表选择攻击范围内所有人)
}
func(self *ChooseByAttackRegion) GetName() string{
	return "ChooseByAttackRegion"
}
func(self *ChooseByAttackRegion) String() string{
	return fmt.Sprintf("对象:攻击者攻击范围内最靠近的%v个角色",self.num)
}
//可以设置选择的对象人数
func(self *ChooseByAttackRegion) Config(args ...interface{}){
	if len(args)>0{
		self.num = args[0].(int)
		if self.num<0{
			self.num = -1 //代表全体攻击
		}
	}
}
//选择器开始进行对象选择
func(self *ChooseByAttackRegion) Choose(fromWho interface{},stepName string,params ...interface{})(targets []effect.EffectCarrier,error bool) {
	error = true
	from, ok := fromWho.(*battle.Warrior)
	if ok {
		//首先获得攻击者的敌人
		enemy := from.BattleIn.GetEnemyAlivedWarriors(from.Player)
		if len(enemy) > 0 {

//			fmt.Printf(" before choose,has %v enemies \n",len(enemy))
			//再获得自己的攻击范围
			attackRange := from.NormalAttackSection.RangeFrom(from.Position)
			//在这些敌人中，排除攻击范围外的人
			for i := len(enemy) - 1; i >= 0; i-- {
				reached := false
				for _, r := range attackRange {
					_, dist := r.DistanceTo(enemy[i].Position)
					if dist == 0 {
						fmt.Printf(" 可以够到\n")
						reached = true
						break
					}
				}
				if !reached {
					enemy = append(enemy[:i], enemy[i + 1:]...)
				}else {
					//如果能够到，则更新其排序字段 = 与攻击者距离
					enemy[i].SortSeed = from.BattleIn.Field.GetDistanceBetween(from, enemy[i])
				}
			}

			fmt.Printf(" after choose,has %v enemies \n",len(enemy))
			n := self.num
			//离攻击者最近的放前面
			characters := (characterSliceLessByDistance)(enemy)
//			fmt.Printf(" after choose,characters %v  \n",len(characters))


			if !sort.IsSorted(characters) {
				sort.Sort(characters)
			}
//			fmt.Printf(" after sort,characters %v  \n",len(characters))

			warriors := characters
			if n<len(characters){
				warriors = characters[0:n]
			}
//			fmt.Printf(" after choose,warriors %v  \n",warriors)

			targets = make([]effect.EffectCarrier, 0, len(warriors))

			for id, _ := range warriors {
				targets = append(targets, warriors[id])
			}

			error = false
		}
	}
	return
}