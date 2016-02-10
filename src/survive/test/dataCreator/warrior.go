package dataCreator
import (
	"survive/server/logic/battle"
	"survive/server/logic/consts/nature"
	"survive/server/logic/skill/targetChoose"
	"survive/server/logic/skill/targetChoose/builtIn"
)
func init(){
	builtIn.RegBuiltInChoosers()
}
var seed_warrior int = 0
func GetWarrior() *battle.Warrior{

	targetChooser := targetChoose.Create("ChooseByAttackRegion")
	targetChooser.Config(2) //设置可以攻击2人

	seed_warrior++
	ch1 := GetCharacter()
	var warrior1 *battle.Warrior
	if seed_warrior %2 == 0 {
		warrior1 = battle.NewWarrior(ch1,nature.Physical,targetChooser,12,0,20,30,200,20,1,30)
	}else{
		warrior1 = battle.NewWarrior(ch1,nature.Magical,targetChooser,12,0,20,30,200,20,1,30)
	}

	return warrior1
}
func GetWarriors(n int) []*battle.Warrior{

	warriors := make([]*battle.Warrior,0)
	for i:=1;i<=n;i++{
		warriors = append(warriors,GetWarrior())
	}

	return warriors
}
