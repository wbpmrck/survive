package effect
import (
	"survive/server/logic/character"
	"survive/server/logic/dataStructure"
	"survive/server/logic/rule"
)

/*
1 效果是技能的基础
2 定义效果接口，然后要抽象出游戏中可能出现的任意效果
3 每个效果，都会定义
4 效果并不一定只能发生在人身上，所有物体都可能会带有效果
 */
//todo:以后再考虑，怎么用动态语言去实现这么多效果
type Effect interface {
//	PutOn(time *dataStructure.Time,from, target *character) bool //产生效果，一个从from发动，丢给target的动作. 如果该对象无法被放置该效果，则返回false
	PutOn(time *dataStructure.Time,from, target rule.EffectCarrier) bool //产生效果，一个从from发动，丢给target的动作. 如果该对象无法被放置该效果，则返回false
	Config(args ...interface{}) //技能在组合效果的时候，可以给效果设置一些参数
	IsAlive() bool //每个效果都要记录自己是否仍然存活
	GetFrom() rule.EffectCarrier //获取效果的发出方
//	GetFrom() *character.Character //获取效果的发出方
	GetTarget() rule.EffectCarrier //获取效果的作用方
//	GetTarget() *character.Character //获取效果的作用方
	GetId() string //获取效果的名字
	Remove(time *dataStructure.Time) bool //移除效果,移除成功返回true
	GetInfo() string //获取描述信息
}

//定义效果创建工厂类型（这样就可以让不同的效果对象注册工厂方法，方便动态根据 效果名=>效果对象）
type EffectFactory func() *Effect


//存放所有的效果工厂
var allEffectsFactory map[string]EffectFactory = make(map[string]EffectFactory)

//注册一个效果工厂
func RegisterFactory(name string,effectFactory EffectFactory){
	allEffectsFactory[name] = effectFactory
}

//根据效果名，获取一个效果对象
func Get(name string) *Effect{
	return allEffectsFactory[name]()
}