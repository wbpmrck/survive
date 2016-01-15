package builtIn
import (
	"survive/server/logic/character"
	"sort"
	"survive/server/logic/skill/targetChoose"
)

//定义interface{},并实现sort.Interface接口的三个方法
type CharacterSlice []*character.Character

func (c CharacterSlice) Len() int {
	return len(c)
}
func (c CharacterSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c CharacterSlice) Less(i, j int) bool {
	return c[i].HP < c[j].HP
}

//一种选择对象的策略：按照虚弱程度选择
//params[0]:设置最弱的n个人
func ChooseTargets(from *character.Character,params ...interface{})(targets []*character.Character,error bool)  {
	battleField := from.Battlefield
	enemy := battleField.GetEnemy(from.Player)
	if len(enemy)>0{
		//这里为了简单，就取第一个对手势力
		firstEnemy :=enemy[0]
		n:= int(params[0])
		//取最弱的n个,先按照hp排序，最少的放前面
		characters := (CharacterSlice)(battleField.PlayerCharacters[firstEnemy.Id])
		if !sort.IsSorted(characters) {
			sort.Sort(characters)
		}

		//然后返回前n个
		return battleField.PlayerCharacters[firstEnemy.Id][0:n]
	}else{
		return nil,nil
	}
}

func init(){
	targetChoose.Register("ByWeak",ChooseTargets)
}