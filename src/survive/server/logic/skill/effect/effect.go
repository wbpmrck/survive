package effect
import (
	"survive/server/logic/character"
	"survive/server/logic/dataStructure"
)

//效果是技能的基础
//定义效果接口，然后要抽象出游戏中可能出现的任意效果
//每个效果，都会定义
//todo:以后再考虑，怎么用动态语言去实现这么多效果
type Effect interface {
	PutOn(time *dataStructure.Time,from, target *character.Character) //产生效果，一个从from发动，丢给target的动作
	Config(args ...interface{}) //技能在组合效果的时候，可以给效果设置一些参数
//	Update(args ...interface{}) //进行生效处理(Update可能多次被调用)
	Remove(time *dataStructure.Time) //移除效果
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