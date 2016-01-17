package battle
import (
	"survive/server/logic/character"
	"survive/server/logic/player"
)

//代表1场战斗
//包括战场地形、战斗双方所有单位，以及战斗情况等信息
type Battle struct {
	Players map[string]*player.Player //参与战斗的玩家
	PlayerCharacters map[string]map[string]*character.Character //key:player.Id value:玩家在这场战斗中投入的角色字典
	Field *BattleField
}
//获取某个玩家的敌对玩家(目前，除了自己都是敌人)
func(this *Battle) GetEnemy(me *player.Player) []*player.Player{
	enemy := make([]*player.Player)
	for k,v:= range this.Players{
		if k != me.Id{
			enemy = append(enemy,v)
		}
	}
	return enemy
}
//加入某个角色到战斗中
func(this *Battle) AddCharacter(c *character.Character,*player.Player) bool{
	_,exist:=this.Players[c.Id]
	if exist{
		this.PlayerCharacters[c.Id] = c
		return true
	}else{
		return false
	}
}

//创建1场战斗，选定作战多方
func NewBattle( filedLen int,players ...*player.Player) *Battle{
	b := &Battle{
		Players:make(map[string]*player.Player),
		PlayerCharacters:make(map[string]map[string]*character.Character),
		Field:NewField(filedLen),
	}
	//初始化用户角色集合
	for _,v:= range players{
		b.PlayerCharacters[v.Id] = make(map[string]*character.Character)
	}
	return b
}