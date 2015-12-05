package handler
import "github.com/name5566/leaf/gate"




func NewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

