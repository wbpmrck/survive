package bases
import (
	"survive/server/logic/character"
	"time"
)

type EffectBase struct {
	//效果的使用者，和受众
	From, Target *character.Character
	PutOnTime time.Time //效果生效时间
	RemoveTime time.Time //效果结束时间
}
func(self *EffectBase) PutOn(from, target *character.Character){
	self.PutOnTime = time.Now()
	self.From = from
	self.Target = target
}
//效果移除
func(self *EffectBase) Remove(){
	self.RemoveTime = time.Now()
	self.From = nil
	self.Target = nil
}