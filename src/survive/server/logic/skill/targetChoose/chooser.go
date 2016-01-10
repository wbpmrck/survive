package targetChoose
import "survive/server/logic/character"

//负责进行技能对象的选择
//可以开发不同的chooser来搭配各种各样的技能

//根据技能发动者(以及他所携带的信息),找到技能的发动目标,返回目标列表。
//如果targets=nil,err=true,则表示当前没有可命中的目标
//如果targets=nil,err=false,则可能表示这个技能是一个非指向型技能
//总之，只要err不为true,就可以调用 Install 来发动技能

type TargetChooser func(from *character.Character,params ...interface{})(targets []*character.Character,error bool)

const AllChooserFunc map[string]TargetChooser = make(map[string]TargetChooser)

