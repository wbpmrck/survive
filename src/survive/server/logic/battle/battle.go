package battle
import (
	"survive/server/logic/player"
	"survive/server/logic/dataStructure"
	"logProxy/logger"
	"survive/server/logic/rule/event"
)

//战斗状态类型
type BattleStatus int
func(self BattleStatus) String()string{
	switch self {
	case STATUS_ING:
		return "正在进行"
	case STATUS_INIT:
		return "初始化"
	case STATUS_OVER:
		return "已经结束"
	default:
		return "未知"
	}
}
const (
	STATUS_INIT BattleStatus = iota //初始化
	STATUS_ING  //进行中
	STATUS_OVER //结束
)

const (
	EVENT_START="battle-start"
	EVENT_END="battle-end"
)

//代表1场战斗
//包括战场地形、战斗双方所有单位，以及战斗情况等信息
type Battle struct {
	*event.EventEmitterBase //战斗也可能发射出各种事件
	Players map[string]*player.Player //参与战斗的玩家
	PlayerCharacters map[string]map[string]*Warrior //key:player.Id value:玩家在这场战斗中投入的角色字典(key:角色id value,角色)
	PlayerCharactersList map[string][]*Warrior //key:player.Id value:玩家在这场战斗中投入的角色列表
	Field *BattleField
	Status BattleStatus
	Report *BattleReport //战斗报告
}
//获取某个玩家的敌对玩家(目前，除了自己都是敌人)
func(this *Battle) GetEnemy(me *player.Player) []*player.Player{
	enemy := make([]*player.Player,0)
	for k,v:= range this.Players{
		if k != me.Id{
			enemy = append(enemy,v)
		}
	}
	return enemy
}
//加入某个角色到战斗中
func(this *Battle) AddCharacter(c *Warrior) bool{
	_,exist:=this.Players[c.Player.Id]
	if exist{
		_,existWarrior := this.PlayerCharacters[c.Player.Id][c.Id]
		if !existWarrior{
			this.PlayerCharacters[c.Player.Id][c.Id] = c
			this.PlayerCharactersList = append(this.PlayerCharactersList,c)
			return true
		}
		return false
	}else{
		return false
	}
}
//战斗开始
func(self *Battle) Start(){
	self.Status = STATUS_ING
	self.Emit(EVENT_START,self)
}
//战斗结束
func(self *Battle) END(){
	self.Status = STATUS_OVER
	self.Emit(EVENT_END,self)
}

//todo:warrior 实现 time.Receiver 接口，这样就可以得到时间片来进行行动了
func(self *Battle) Receive(ts dataStructure.TimeSpan){

	switch self.Status {
	case STATUS_INIT:
		self.Start()
		fallthrough
	case STATUS_ING:
		/*
			进行战斗:
			* 从行动队列里，取出行动顺序最高的角色，按照顺序分配动作时间片
		 */




		//最后:累加进行时间
		self.Report.AddTimeConsumed(ts)


	case STATUS_OVER:
		//如果战斗已经结束，还收到时间片，尝试再次结束战斗
		self.END()
	default:
		logger.GetLogger().Errorf("wrong status :%v",self.Status)
	}
}
//创建1场战斗，选定作战多方
func NewBattle( filedLen int,players ...*player.Player) *Battle{
	b := &Battle{
		Players:make(map[string]*player.Player),
		PlayerCharacters:make(map[string]map[string]*Warrior),
		PlayerCharactersList:make(map[string][]*Warrior),
		Field:NewField(filedLen),
		Status:STATUS_INIT,
		Report:&BattleReport{},
	}
	//初始化用户角色集合
	for _,v:= range players{
		b.PlayerCharacters[v.Id] = make(map[string]*Warrior)
	}
	return b
}