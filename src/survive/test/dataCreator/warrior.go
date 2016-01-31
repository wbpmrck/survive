package dataCreator
import (
	"survive/server/logic/battle"
	"survive/server/logic/consts/nature"
)

var seed_warrior int = 0
func GetWarrior() *battle.Warrior{
	seed_warrior++
	ch1 := GetCharacter()
	var warrior1 *battle.Warrior
	if seed_warrior %2 == 0 {
		warrior1 = battle.NewWarrior(ch1,nature.Physical,12,0,200,30,200,20,12,30)
	}else{
		warrior1 = battle.NewWarrior(ch1,nature.Magical,12,50,400,30,200,20,12,30)
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
