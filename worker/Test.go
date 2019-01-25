package worker

import (
	"fmt"
	"github.com/Deansquirrel/go-tool"
	"time"
)

type testWorker struct {
}

func NewTestWorker() (*testWorker, error) {
	return nil, nil
}

func (tw *testWorker) TestPublish() {

	//for i := 0;i<1000;i++ {
	//	err := global.RabbitMQ.Publish("","exchangeFanOut","","Test Msg")
	//	if err != nil {
	//		common.PrintOrLog(err.Error())
	//	}
	//}

	//_ = global.RabbitMQ.Publish("","exchangeTopic","*.*.*","exchangeTopic message")
	//_ = global.RabbitMQ.Publish("","exchangeTopic","*.B.C","exchangeTopic booking message")
	//err := global.RabbitMQ.Publish("","exchangeTopic","A.B.*","exchangeTopic create message")
	//if err != nil {
	//	common.PrintOrLog(err.Error())
	//}

	r := go_tool.NewRabbitMQ("sa","Zl84519741","47.99.146.112",5672,"goTest",time.Second * 30,time.Millisecond * 500,3,time.Second * 10)

	err := r.AddConsumer("","TktCreateYwdetail",subHandler)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func subHandler(msg string){
	fmt.Println(msg)
}
