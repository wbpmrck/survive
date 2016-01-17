package battle
import (
	"survive/server/logic/character"
	"survive/server/logic/dataStructure"
)

/*
	代表战场的地形
 */
type BattleField struct {
	field []map[string]*character.Character
}
//添加一个角色进入战场地形
func(self *BattleField) AddCharacter(c *character.Character,pos dataStructure.BattlePos){
	self.field[pos][c.Id] = c
	c.Position = pos
}

//移动一个角色在战场的位置
func(self *BattleField) MoveCharacter(c *character.Character,posTo dataStructure.BattlePos){
	if(c.Position!=posTo){
		//从旧的位置删除
		delete(self.field[c.Position],c.Id)
		//移动到新位置
		c.Position = posTo
		self.field[posTo][c.Id] = c
	}

}
//创建战场
func NewField(length int) *BattleField {
	field := make([]map[string]*character.Character,length)

	//初始化每个格子
	for k,_:= range field{
		field[k] = make(map[string]*character.Character)
	}

	return &BattleField{
		field:field,
	}
}