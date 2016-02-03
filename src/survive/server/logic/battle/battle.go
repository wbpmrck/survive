package battle
import (
	"survive/server/logic/player"
	"survive/server/logic/dataStructure"
	"survive/server/logger"
	"survive/server/logic/rule/event"
	"github.com/wbpmrck/golibs/linkedList"
	"survive/server/logic/dataStructure/attribute"
	"fmt"
	"survive/server/logic/skill/effect"
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
	EVENT_TIME_ELAPSED="battle-time-elapsed"
)

//比较函数：当一个新的节点插入的时候，根据行动顺序决定放在哪里
func compareWarriorActSeq(old, new interface {}) bool {
	if new.(*Warrior).ActSeq.GetValue().Get()>= old.(*Warrior).ActSeq.GetValue().Get() {
		return true
	}
	return false
}


//代表1场战斗
//包括战场地形、战斗双方所有单位，以及战斗情况等信息
type Battle struct {
	*event.EventEmitterBase //战斗也可能发射出各种事件
	*effect.EffectCarrierBase //战斗是一个效果携带者（一些与具体对象无关的效果，就施加在战场上，比如AOE）
	Desc string //战斗描述
	Players map[string]*player.Player //参与战斗的玩家
	PlayerCharacters map[string]map[string]*Warrior //key:player.Id value:玩家在这场战斗中投入的角色字典(key:角色id value,角色)
	PlayerCharactersList map[string][]*Warrior //key:player.Id value:玩家在这场战斗中投入的角色列表
	Field *BattleField
	Status BattleStatus
	Report *BattleReport //战斗报告
	ActSeq *linkedList.SortedLinkedList//战斗行动顺序列表(排序的链表)
}
func(self *Battle) String()string{
	return fmt.Sprintf("%v,player:[%v],status:[%v]",self.Desc,self.Players,self.Status)
}
//获取某个玩家的敌对玩家(目前，除了自己都是敌人)
func(self *Battle) GetEnemy(me *player.Player) []*player.Player{
	enemy := make([]*player.Player,0)
	for k,v:= range self.Players{
		if k != me.Id{
			enemy = append(enemy,v)
		}
	}
	return enemy
}
//加入某个角色到战斗中
func(self *Battle) AddWarrior(c *Warrior) bool{
//	fmt.Printf("add warrior [%s]\n",c.GetShowName())
	_,exist:= self.Players[c.Player.Id]
	if exist{
		//判断这个角色是不是已经添加过了
//		fmt.Printf("判断这个角色是不是已经添加过了 warrior [%s]\n",c.GetShowName())
		_,existWarrior := self.PlayerCharacters[c.Player.Id][c.Id]
		if !existWarrior{
//			fmt.Printf("new warrior [%s]\n",c.GetShowName())
			self.PlayerCharacters[c.Player.Id][c.Id] = c
			self.PlayerCharactersList[c.Player.Id] = append(self.PlayerCharactersList[c.Player.Id],c)
			//角色没有添加过，添加之，并订阅其行动顺序变化事件，来更新行动先后列表
//			fmt.Printf("before insert,ActSeq.len is %v/%v \n",self.ActSeq.Len(),self.ActSeq.Limit)
			cNode := self.ActSeq.PutOnTop(c)
//			fmt.Printf("after insert,ActSeq.len is %v/%v \n",self.ActSeq.Len(),self.ActSeq.Limit)
			c.ActSeq.On(attribute.EVENT_VALUE_CHANGED,
				//订阅所依赖属性的变化事件，第一个参数是变化的属性
				event.NewEventHandler(func (contextParams ...interface{}) (isCancel bool,handleResult interface{}){
						//先从行动顺序队列中移出
//					fmt.Printf("before remove %v/%v \n",self.ActSeq.Len(),self.ActSeq.Limit)
						self.ActSeq.Remove(cNode)
//					fmt.Printf("after remove %v/%v \n",self.ActSeq.Len(),self.ActSeq.Limit)
						//再重新插入
					cNode =self.ActSeq.PutOnTop(c)
//						fmt.Printf("变更的属性是“行动”顺序,ActSeq.len is %v/%v \n",self.ActSeq.Len(),self.ActSeq.Limit)
					return
				}))
			//订阅角色的死亡事件，用于计算战斗结束条件
			c.On(EVENT_WARRIOR_DEAD,event.NewEventHandler(func (contextParams ...interface{}) (isCancel bool,handleResult interface{}){
				self.judgeEnd()
				return
			}))

			//订阅角色的日志输出事件，记录各种日志
			c.On(EVENT_WARRIOR_NEW_ACTION_RECORD,event.NewEventHandler(func (contextParams ...interface{}) (isCancel bool,handleResult interface{}){
				record := contextParams[0].(ActionRecord)
				self.Report.AddRecord(record)

				return
			}))

			return true
		}
		return false
	}else{
		return false
	}
}
//判断战斗是否要结束，如果结束，进行结束动作
func(self *Battle) judgeEnd(){
	//判断是否有一方的角色全部死掉
	playerId := ""
	for playerId,_ = range self.PlayerCharactersList{
		for _,fighter := range self.PlayerCharactersList[playerId]{
			//如果有角色没死，则直接取消对该玩家角色的继续搜索
			if fighter.HP.GetValue().Get()>0.0{
				break
			}
		}
		//如果能走到这里，说明这个玩家的所有角色都死掉，游戏结束
		loser := self.Players[playerId]
		self.Report.SetLoser(loser)
		winner := self.GetEnemy(loser)[0]
		self.Report.SetWinner(winner)
		self.END()
	}

}
//战斗开始
func(self *Battle) Start(){
//	fmt.Printf("battle %v begin \n",self.Desc)
	self.Status = STATUS_ING
	self.Emit(EVENT_START,self)
}
//战斗结束
func(self *Battle) END(){
//	fmt.Printf("battle %v end \n",self.Desc)
	self.Status = STATUS_OVER
	self.Emit(EVENT_END,self)
}
func(self *Battle) Init(now dataStructure.Time){
//	fmt.Printf("battle %v Init \n",self.Desc)
	self.Report.StartTime = now
}
//todo:warrior 实现 time.Receiver 接口，这样就可以得到时间片来进行行动了
func(self *Battle) Receive(ts dataStructure.TimeSpan){

//	fmt.Printf("Receive:[%v] ,status is: [%v]\n",ts,self.Status)
	switch self.Status {
	case STATUS_INIT:
		self.Start()
		fallthrough
	case STATUS_ING:
//		fmt.Printf("actSeq:[%v]\n",self.ActSeq.Len())
		/*
			进行战斗:
			* 从行动队列里，取出行动顺序最高的角色，按照顺序分配动作时间片
		 */
		for oneWarrior := self.ActSeq.List.Front(); oneWarrior != nil; oneWarrior = oneWarrior.Next() {
			w := oneWarrior.Value.(*Warrior)
			//对角色进行时间片输入
			w.Receive(ts)
			//对角色进行动作，并将返回的操作日志记录
//			records :=w.Receive(ts)
//			if records!=nil && len(records)>0{
//				self.Report.AddRecords(records)
//			}
		}

		//最后:累加进行时间
		self.Report.AddTimeConsumed(ts)
		//发射事件 func(目前为止总时间，本次流失时间)
		self.Emit(EVENT_TIME_ELAPSED,self.Report.TimeConsumed,ts)


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
		Desc:"",
		EventEmitterBase:event.NewEventEmitter(),
		EffectCarrierBase:effect.NewEffectCarrier(),
		Players:make(map[string]*player.Player),
		PlayerCharacters:make(map[string]map[string]*Warrior),
		PlayerCharactersList:make(map[string][]*Warrior),
		Field:NewField(filedLen),
		Status:STATUS_INIT,
		Report:NewBattleReport(),
		//创建行动顺序链表(上限为:玩家人数*10,就是每个玩家最多10个英雄参战)
		ActSeq:linkedList.NewSortedLinkedList(len(players)*10, compareWarriorActSeq),
	}
	//初始化用户角色集合
	for _,v:= range players{
		b.Players[v.Id] = v
		b.PlayerCharacters[v.Id] = make(map[string]*Warrior)
		b.PlayerCharactersList[v.Id] = make([]*Warrior,0)
	}
	return b
}