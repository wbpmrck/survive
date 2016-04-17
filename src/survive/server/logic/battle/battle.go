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
	EVENT_CHARACTER_ENTER="battle-character-enter"
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

//获取某个玩家的敌对玩家的所有角色
func(self *Battle) GetEnemyWarriors(me *player.Player) []*Warrior{
	enemy := make([]*Warrior,0)
	for k,v:= range self.Players{
		if k != me.Id{
			enemy = append(enemy,self.PlayerCharactersList[v.Id]...)
		}
	}
	return enemy
}
//获取某个玩家的敌对玩家的所有角色
func(self *Battle) GetEnemyAlivedWarriors(me *player.Player) []*Warrior{
	enemy := make([]*Warrior,0)
	for k,v:= range self.Players{
		if k != me.Id{
			enemy = append(enemy,self.PlayerCharactersList[v.Id]...)
		}
	}
	for i := len(enemy) - 1; i >= 0; i-- {
		if enemy[i].Status ==WARRIOR_DEAD {
			enemy = append(enemy[:i], enemy[i + 1:]...)
		}
	}
	return enemy
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
func(self *Battle) AddWarrior(c *Warrior,pos int) bool{
//	fmt.Printf("add warrior [%s]\n",c.GetShowName())
	_,exist:= self.Players[c.Player.Id]
	if exist{
		//判断这个角色是不是已经添加过了
//		fmt.Printf("判断这个角色是不是已经添加过了 warrior [%s]\n",c.GetShowName())
		_,existWarrior := self.PlayerCharacters[c.Player.Id][c.Id]
		if !existWarrior{
//			fmt.Printf("new warrior [%s]\n",c.GetShowName())
			c.BattleIn = self
			//加入战场
			self.Field.AddCharacter(c,pos)

			//发射事件 func(目前为止总时间，加入的角色信息)
			self.Emit(EVENT_CHARACTER_ENTER,self.Report.TimeConsumed,c)

			//添加日志：
			rc := &CharacterEnterAction{
				ActionRecordBase:NewActionRecordBase(self.Report.TimeConsumed,c,ACTION_TYPE_CHARACTER_ENTER,
								fmt.Sprintf("%v 加入战斗",c.GetShowName())),
				Props:make(map[string]float64),
			}
			rc.Props[c.HP.GetName()] = c.HP.GetValue().Get()
			rc.Props[c.AP.GetName()] = c.AP.GetValue().Get()
			rc.Props[c.OP.GetName()] = c.OP.GetValue().Get()
			rc.Props["Position"] = float64(c.Position)
			self.Report.AddRecord(rc)


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

			//监听器，处理一些数值变化事件
			attributeListener := event.NewEventHandler(func (contextParams ...interface{}) (isCancel bool,handleResult interface{}){
				attr := contextParams[0].(*attribute.Attribute)
				record := &AttrChangeAction{
					ActionRecordBase:NewActionRecordBase(self.Report.TimeConsumed,c,ACTION_TYPE_ATTRIBUTE_CHANGE,
						fmt.Sprintf("%v 属性变化: [%v]",c.GetShowName(),attr)),
					AttrName:attr.GetName(),
					AttrValue:attr.GetValue().Get(),
				}
				self.Report.AddRecord(record)
				return
			})
			//订阅角色的数值变化事件，记录战斗日志
			c.AP.On(attribute.EVENT_VALUE_CHANGED,attributeListener)
			c.HP.On(attribute.EVENT_VALUE_CHANGED,attributeListener)
			c.OP.On(attribute.EVENT_VALUE_CHANGED,attributeListener)

			//订阅角色的各种事件，记录各种日志
			c.On("*",event.NewEventHandler(func (contextParams ...interface{}) (isCancel bool,handleResult interface{}){
				evtName := contextParams[0].(string)
				r:=AnalyzeReportFromEvent(self.Report.TimeConsumed,contextParams...)
				if r!=nil{
					self.Report.AddRecord(r)
				}

				//角色的死亡事件，用于计算战斗结束条件
				if evtName == EVENT_WARRIOR_DEAD{
					self.judgeEnd()
				}

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
	aliveCount := 0
	for playerId,_ = range self.PlayerCharactersList{
		for _,fighter := range self.PlayerCharactersList[playerId]{
			//如果有角色没死，则直接取消对该玩家角色的继续搜索
			if fighter.HP.GetValue().Get()>0.0{
				aliveCount++
				break
			}
		}
		if aliveCount == 0 {
			//如果能走到这里，说明这个玩家的所有角色都死掉，游戏结束
			loser := self.Players[playerId]
			self.Report.SetLoser(loser)
			winner := self.GetEnemy(loser)[0]
			self.Report.SetWinner(winner)
			self.END()
		}
		aliveCount = 0
	}

}
//战斗开始
func(self *Battle) Start(){
//	fmt.Printf("battle %v begin \n",self.Desc)
	self.Status = STATUS_ING

	//添加日志：
	rc := &BattleInitAction{
		ActionRecordBase:NewActionRecordBase(self.Report.TimeConsumed,nil,ACTION_TYPE_START_BATTLE, "战斗开始"),
		FieldLength:self.Field.GetLen(),
	}
	self.Report.AddRecord(rc)

	self.Emit(EVENT_START,self)
}
//战斗结束
func(self *Battle) END(){
//	fmt.Printf("battle %v end \n",self.Desc)
	self.Status = STATUS_OVER

	//添加日志：
	rc := NewActionRecordBase(self.Report.TimeConsumed,nil,ACTION_TYPE_END_BATTLE, "战斗结束")
	self.Report.AddRecord(rc)

	self.Emit(EVENT_END,self)
}
func(self *Battle) Init(now dataStructure.Time){
//	fmt.Printf("battle %v Init \n",self.Desc)
	self.Report.StartTime = now
}
//todo:warrior 实现 time.Receiver 接口，这样就可以得到时间片来进行行动了
func(self *Battle) Receive(ts dataStructure.TimeSpan){

	fmt.Printf("Receive:[%v] ,status is: [%v]\n",ts,self.Status)
	switch self.Status {
	case STATUS_INIT:
		self.Start()
		fallthrough
	case STATUS_ING:
//		fmt.Printf("Receive actSeq:[%v]\n",self.ActSeq.Len())
		//累加进行时间
		self.Report.AddTimeConsumed(ts)
		/*
			进行战斗:
			* 从行动队列里，取出行动顺序最高的角色，按照顺序分配动作时间片
		 */
		for oneWarrior := self.ActSeq.List.Front(); oneWarrior != nil; oneWarrior = oneWarrior.Next() {
			w := oneWarrior.Value.(*Warrior)
			if w.Status != WARRIOR_DEAD{
				//对角色进行时间片输入
				w.Receive(ts)
			}

		}

		//发射事件 func(目前为止总时间，本次流失时间)
		self.Emit(EVENT_TIME_ELAPSED,self.Report.TimeConsumed,ts)


	case STATUS_OVER:
		//如果战斗已经结束，还收到时间片，不做事情
//		self.END()
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