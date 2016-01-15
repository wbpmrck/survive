package rule
/*
	fightable 表示一个可进行战斗行为的单位
 */
type Fightable interface {
	Moveable //战斗单位必须是一个 可移动 单位
	AttributeCarrier //战斗单位具有各种属性，可以进行获取和修改
}

/*
	表示一个符合战斗序列要求的单位
 */
type FightOrderable interface {


}