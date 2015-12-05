package handler
import "github.com/name5566/leaf/gate"


func CloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}
