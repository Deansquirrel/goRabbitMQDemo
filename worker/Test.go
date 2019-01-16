package worker

import (
	"github.com/Deansquirrel/go-tool"
	"github.com/Deansquirrel/goRabbitMQDemo/common"
	"github.com/Deansquirrel/goRabbitMQDemo/global"
	"time"
)

type testWorker struct {
}

func NewTestWorker() (testWorker, error) {
	return struct{}{}, nil
}

func (tw *testWorker) TestPublish() {
	var err error
	common.PrintOrLog("=========================================================")
	for i := 0; i < 50; i++ {
		err = global.RabbitMQ.Publish("exchangeFanOut", "", "hello word")
		if err != nil {
			common.PrintOrLog(err.Error())
		}
	}
	common.PrintOrLog("=========================================================")
	for i := 0; i < 50; i++ {
		err = global.RabbitMQ.Publish("exchangeFanOut", "", "hello word222")
		if err != nil {
			common.PrintOrLog(err.Error())
		}
	}
	common.PrintOrLog("=========================================================")
	time.Sleep(time.Second * 5)
	common.PrintOrLog("=========================================================")
	go func() {
		for {
			val, ok, err := global.RabbitMQ.Consume("A")
			if err != nil {
				common.PrintAndLog(err.Error())
			} else if ok {
				common.PrintOrLog(go_tool.GetDateTimeStr(time.Now()) + "-" + "A" + "-" + val)
			} else {
				common.PrintOrLog(go_tool.GetDateTimeStr(time.Now()) + "-" + "A" + "-" + "null")
				time.Sleep(time.Second * 15)
			}
		}
	}()
	common.PrintOrLog("=========================================================")
	go func() {
		for {
			val, ok, err := global.RabbitMQ.Consume("B")
			if err != nil {
				common.PrintAndLog(err.Error())
			} else if ok {
				common.PrintOrLog(go_tool.GetDateTimeStr(time.Now()) + "-" + "B" + "-" + val)
			} else {
				common.PrintOrLog(go_tool.GetDateTimeStr(time.Now()) + "-" + "B" + "-" + "null")
				time.Sleep(time.Second * 15)
			}
		}
	}()
	common.PrintOrLog("=========================================================")
}

func handler(body string) {
	common.PrintOrLog(body)
}
