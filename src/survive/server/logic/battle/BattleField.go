package battle
import (
	"survive/server/logic/math"
	"fmt"
)

/*
	代表战场的地形
 */
type BattleField struct {
	field []map[string]*Warrior
}
func(self *BattleField) GetLen() int{
	return len(self.field)
}
//获取战场上2个角色之间的距离
func(self *BattleField) GetDistanceBetween(from,to *Warrior) int{
	//由于战场坐标为[0,n],所以可以直接求绝对值
	return math.AbsInt(from.Position - to.Position)
}

//找到在战场上，攻击者attacker 攻击范围内外的最近的一个敌人
func(self *BattleField) FindNearestAttackTarget(attacker *Warrior) (side int,distance int,target *Warrior){
	target = nil
	//首先获得攻击者的敌人
	enemies := attacker.BattleIn.GetEnemyAlivedWarriors(attacker.Player)

	//再获得自己的攻击范围
	attackRange := attacker.NormalAttackSection.RangeFrom(attacker.Position)
	var minRange int = -1

	//再遍历敌人，看谁的距离最靠近攻击范围边缘
	for _,enemyWarrior:= range enemies{
		for _,r := range attackRange{
			s,thisDistance := r.DistanceTo(enemyWarrior.Position)
			if minRange == -1 || minRange > thisDistance{
				//找到更近的一个敌人
				minRange = thisDistance
				side = s
				target = enemyWarrior
			}
		}
	}

	fmt.Printf("%v 找到最靠近的敌人:[%v] \n",attacker,target)

	//返回
	return side,minRange,target
}
//添加一个角色进入战场地形
func(self *BattleField) AddCharacter(c *Warrior,pos int){
	self.field[pos][c.Id] = c
	c.Position = pos
}

//移动一个角色在战场的位置
func(self *BattleField) MoveCharacter(c *Warrior,posTo int)  (moved int){
	moved = 0

	if(c.Position!=posTo){
		moved = math.AbsInt((posTo - c.Position))
		//从旧的位置删除
		delete(self.field[c.Position],c.Id)
		//移动到新位置
		c.Position = posTo
		self.field[posTo][c.Id] = c
	}

	return moved
}
//创建战场
func NewField(length int) *BattleField {
	field := make([]map[string]*Warrior,length)

	//初始化每个格子
	for k,_:= range field{
		field[k] = make(map[string]*Warrior)
	}

	return &BattleField{
		field:field,
	}
}