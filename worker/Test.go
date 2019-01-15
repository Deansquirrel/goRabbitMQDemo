package worker

import "github.com/Deansquirrel/goRabbitMQDemo/common"

type testWorker struct {
}

func NewTestWorker() (testWorker, error) {
	return struct{}{}, nil
}

func (tw *testWorker) TestPublish() {
	common.PrintOrLog("=========================================================")

	common.PrintOrLog("=========================================================")
}
