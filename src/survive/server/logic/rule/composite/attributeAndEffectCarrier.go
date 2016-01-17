package composite
import "survive/server/logic/rule"

/*
	同时具备属性、效果的单位
 */
type AttributeAndEffectCarrier interface {
	rule.AttributeCarrier
	rule.EffectCarrier
}
