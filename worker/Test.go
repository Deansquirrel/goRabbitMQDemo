package worker

type testWorker struct {
}

func NewTestWorker() (*testWorker, error) {
	return nil, nil
}

func (tw *testWorker) TestPublish() {
	//_ = global.RabbitMQ.Publish("","exchangeTopic","*.*.*","exchangeTopic message")
	//_ = global.RabbitMQ.Publish("","exchangeTopic","*.B.C","exchangeTopic booking message")
	//err := global.RabbitMQ.Publish("","exchangeTopic","A.B.*","exchangeTopic create message")
	//if err != nil {
	//	common.PrintOrLog(err.Error())
	//}
}
