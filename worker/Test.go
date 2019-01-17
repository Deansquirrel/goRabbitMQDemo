package worker

import (
	"github.com/Deansquirrel/go-tool"
	"github.com/Deansquirrel/goRabbitMQDemo/common"
	"github.com/Deansquirrel/goRabbitMQDemo/global"
	"time"
)

type testWorker struct {
}

func NewTestWorker() (*testWorker, error) {
	return nil, nil
}

func (tw *testWorker) TestPublish() {
	common.PrintOrLog("=========================================================")
	chQ := make(chan struct{})
	time.AfterFunc(time.Second*15, func() {
		chQ <- struct{}{}
	})
	common.PrintOrLog("=========================================================")
	err := global.RabbitMQ.AddConsumer("testC,", "A", handler, time.Second, time.Second)
	if err != nil {
		common.PrintOrLog(err.Error())
	}
	common.PrintOrLog("=========================================================")
round:
	for {
		go func() {
			err := global.RabbitMQ.Publish("testP", "exchangeDirect", "", go_tool.GetDateTimeStr(time.Now()))
			if err != nil {
				common.PrintOrLog(err.Error())
				return
			}
			common.PrintOrLog("done")
		}()
		time.Sleep(time.Second)
		select {
		case <-chQ:
			break round
		default:
		}
	}
	common.PrintOrLog("=========================================================")
}

func handler(msg string) {
	common.PrintOrLog(msg)
}
