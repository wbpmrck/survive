package builtIn
import (
	"survive/server/logic/skill/targetChoose"
)

func RegBuiltInChoosers(){

	targetChoose.RegisterFactory("ChooseByWeak",func() targetChoose.TargetChooser{
		chooser := &ChooseByWeak{
		}
		return chooser
	})
	targetChoose.RegisterFactory("ChooseByAttackRegion",func() targetChoose.TargetChooser{
		chooser := &ChooseByAttackRegion{
		}
		return chooser
	})
}

