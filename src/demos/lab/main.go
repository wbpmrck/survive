package main


import (
	"code/core/agent"
	"fmt"
	"time"
	"code/core/timeRule"
)

//type AgentLogicHandler interface {
//	//handle timeSlice event
//	HandleTimeSlice(ts *TimeSliceMessage)
//
//	//handle message event
//	HandleMessage(msg *AgentMessage)
//}
type TestHandler struct {

}
func (t *TestHandler) HandleTimeSlice(ts *agent.TimeSliceMessage){
	fmt.Printf("HandleTimeSlice，ts is:%v \n",ts.GetTimeSlice())
}

func (t *TestHandler) HandleMessage(msg *agent.AgentMessage){
	fmt.Printf("HandleMessage\n")
	body := fmt.Sprintf("pong! i got your msg:%s",msg.GetMessageBody())
	resp := agent.CreateNotifyAgentMessage(2,body,nil,nil)
	msg.GetResponseChan()<-resp
}

func main(){
	var handler agent.AgentLogicHandler = &TestHandler{}
	var aagent = agent.CreateAgentBase("a",handler,2,2)
	fmt.Println("agent start! \n")
	tsChan := aagent.Start()

	countSendTs := 10
	countSendMsg := 10

	for c:=0;c<countSendTs;c++{
		tsMsg := agent.CreateNotifyTSMessage(timeRule.NewTimeSlice(int64(c)),nil,nil)
		fmt.Printf("send %s to tsChan!\n",tsMsg)
		tsChan <- tsMsg
	}

	msgChan := aagent.GetMessagePipe()
	for c:=0;c<countSendMsg;c++{
		msg := agent.CreateRequestAgentMessage(0,fmt.Sprintf("ping:%v",c),nil,nil)
		fmt.Printf("send %s to msgChan!\n",msg)
		msgChan <- msg
		resp := <-msg.GetResponseChan()
		fmt.Printf("get resp from msgChan! %s \n",resp.GetMessageBody())
	}
	time.Sleep(10*time.Second)
}
