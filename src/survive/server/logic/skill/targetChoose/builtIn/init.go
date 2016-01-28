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
}

