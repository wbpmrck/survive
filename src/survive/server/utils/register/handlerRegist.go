package register
import (
	"github.com/name5566/leaf/module"
	"reflect"
)


func RegistStringHandler(skeleton *module.Skeleton,name string, handler interface{}) {
	skeleton.RegisterChanRPC(name, handler)
}


func RegistMsgHandler(skeleton *module.Skeleton,msg interface{}, handler interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(msg), handler)
}
