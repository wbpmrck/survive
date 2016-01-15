package bases
import (
	"survive/server/logic/character"
	"survive/server/logic/dataStructure"
)

type EffectBase struct {
	//效果的使用者，和受众
	From, Target *character.Character
	PutOnTime *dataStructure.Time//效果生效时间
	RemoveTime *dataStructure.Time //效果结束时间
}
func(self *EffectBase) PutOn(time *dataStructure.Time,from, target *character.Character){
	self.PutOnTime = time
	self.From = from
	self.Target = target
}
//效果移除
func(self *EffectBase) Remove(time *dataStructure.Time,){
	self.RemoveTime = time
	self.From = nil
	self.Target = nil
}