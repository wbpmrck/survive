package battle
import "survive/server/logic/dataStructure"
/*
	表示一场战斗的报告
 */
type BattleReport struct {
	TimeConsumed dataStructure.TimeSpan
}
func(self *BattleReport) AddTimeConsumed(ts dataStructure.TimeSpan){
	self.TimeConsumed = self.TimeConsumed.Add(ts)
}