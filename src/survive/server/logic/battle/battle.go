package battle
import (
	"survive/server/logic/character"
	"survive/server/logic/player"
)

//代表1场战斗
//包括战场地形、战斗双方所有单位，以及战斗情况等信息
type Battle struct {
	Players map[int]*player.Player //参与战斗的玩家
	PlayerCharacters map[int][]*character.Character //key:player.Id value:玩家在这场战斗中投入的角色列表
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
