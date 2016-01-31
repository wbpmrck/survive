package battle
import (
	"survive/server/logic/dataStructure"
	"survive/server/logic/player"
)

type PlayerReport struct {
	Player *player.Player
	IsWin bool //是否获胜
}
/*
	表示一场战斗的报告
 */
type BattleReport struct {
	StartTime dataStructure.Time
	TimeConsumed dataStructure.TimeSpan
	WinnerReport,LoserReport *PlayerReport
}
func(self *BattleReport) SetWinner(p *player.Player){
	self.WinnerReport = &PlayerReport{
		Player:p,
		IsWin:true,
	}
}
func(self *BattleReport) SetLoser(p *player.Player){
	self.LoserReport = &PlayerReport{
		Player:p,
		IsWin:false,
	}
}
func(self *BattleReport) AddTimeConsumed(ts dataStructure.TimeSpan){
	self.TimeConsumed = self.TimeConsumed.Add(ts)
}