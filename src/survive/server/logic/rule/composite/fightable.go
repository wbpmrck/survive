package composite
import "survive/server/logic/rule"
/*
	fightable 表示一个可进行战斗行为的单位
 */
type Fightable interface {
	rule.Moveable //战斗单位必须是一个 可移动 单位
	rule.AttributeCarrier //战斗单位具有各种属性，可以进行获取和修改
	rule.TimeVariable //战斗单位也会随着时间变化
	rule.EffectCarrier //战斗单位同时也可以被加上各种效果
}

/*
	表示一个符合战斗序列要求的单位
 */
type FightOrderable interface {


}