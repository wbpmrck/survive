package skill
import (
"survive/server/logic/character"
)

//效果是技能的基础
//定义效果接口，然后要抽象出游戏中可能出现的任意效果
//每个效果，都会定义
//todo:以后再考虑，怎么用动态语言去实现这么多效果
type Effect interface {
	PutOn(from, target *character.Character) //产生效果，一个从from发动，丢给target的动作
	Remove(from, target *character.Character) //移除效果
}
