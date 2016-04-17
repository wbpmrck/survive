package battle
import (
	"survive/server/logic/dataStructure"
	"survive/server/logic/player"
	"fmt"
)

//type PlayerReport struct {
//	Player *player.Player
//	IsWin bool //是否获胜
//	ActionRecords []ActionRecord //从0开始，按照时间顺序进行的动作序列的行为日志
//}
//func NewPlayerReport(pl *player.Player){
//	return &PlayerReport{
//		Player:pl,
//		IsWin:false,
//		ActionRecords:make([]ActionRecord,0),
//	}
//}
/*
	表示一场战斗的报告
 */
type BattleReport struct {
	StartTime dataStructure.Time
	TimeConsumed dataStructure.TimeSpan
	Winner,Loser *player.Player
	Reports []ActionRecord //记录了玩家的角色在战斗中的表现情况
}
func(self *BattleReport)AddRecords(actionRecords []ActionRecord){

	self.Reports = append(self.Reports,actionRecords...)
	return
}
func(self *BattleReport)AddRecord(actionRecord ActionRecord){
	fmt.Printf("战斗报告(%v):\n[%v] \n",len(self.Reports)+1,actionRecord)
	self.Reports = append(self.Reports,actionRecord)
	return
}
//func(self *BattleReport)AddPlayer(pl *player.Player){
//	self.Reports = append(self.Reports,NewPlayerReport(pl))
//	return
//}
func(self *BattleReport) SetWinner(pl *player.Player){
	self.Winner = pl
	return
}
func(self *BattleReport) SetLoser(pl *player.Player){
	self.Loser = pl
	return
}

func(self *BattleReport) AddTimeConsumed(ts dataStructure.TimeSpan){
	self.TimeConsumed = self.TimeConsumed.Add(ts)
}

func NewBattleReport()*BattleReport{
	return &BattleReport{
		Reports:make([]ActionRecord,0),
	}
}