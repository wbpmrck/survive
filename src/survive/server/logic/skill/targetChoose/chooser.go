package targetChoose
import (
	"survive/server/logic/skill/effect"
)

//负责进行技能对象的选择
//可以开发不同的chooser来搭配各种各样的技能

//根据技能发动者(以及他所携带的信息),找到技能的发动目标,返回目标列表。
//如果targets=nil,err=true,则表示当前没有可命中的目标
//如果targets=nil,err=false,则可能表示这个技能是一个非指向型技能
//总之，只要err不为true,就可以调用 Install 来发动技能


type TargetChooser interface {
	String() string
	GetName() string //每个chooser都有名字
	Config(args ...interface{}) //可以给chooser设置一些参数
	//目标选择器，知道一个技能的发出者，要寻找效果的承受者
	// 第二个参数:stepName,让选择器能够知道当前工作在什么阶段。
	// stepName的作用：
	// 	选择器可以选择不工作，这往往是后台配置人员失误导致的
	//	有时候一个选择器，在不同阶段选取的对象也可以不一样
	Choose(from interface{},stepName string,params ...interface{})(targets []effect.EffectCarrier,error bool)
}
//type TargetChooser func(from *character.Character,params ...interface{})(targets []*character.Character,error bool)

//定义选择器创建工厂类型（这样就可以让不同的选择器对象注册工厂方法，方便动态根据 选择器名=>选择器对象）
type ChooserFactory func() TargetChooser


//存放所有的选择器工厂
var allChooserFactorys map[string]ChooserFactory

//注册一个选择器工厂
func RegisterFactory(name string, chooserFactory ChooserFactory){
	allChooserFactorys[name] = chooserFactory
}

//根据选择器名，获取一个选择器对象
func Create(name string) TargetChooser{
	return allChooserFactorys[name]()
}

func init(){
	allChooserFactorys = make(map[string]ChooserFactory)
}