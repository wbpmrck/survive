package builtIn
import (
	"survive/server/logic/skill/effect"
	"survive/server/logic/skill/effect/bases"
)

func RegBuiltInEffects(){

	effect.RegisterFactory("AttributeModify",func() effect.Effect{
		e:= &AttributeModify{}
		e.EffectBase = bases.NewBase("AttributeModify",e)
		return e
	})
	effect.RegisterFactory("AttributeDecResistance",func() effect.Effect{
		e:= &AttributeDecResistance{}
		e.EffectBase = bases.NewBase("AttributeDecResistance",e)
		return e
	})
}

