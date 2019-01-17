package object

type Producer struct {
}

func (p *Producer) StartPublish(exChange string, key string) {

	//chT := time.Tick(time.Millisecond * 100)
	//go func(){
	//
	//	for t :=  range chT{
	//		ch,err := global.RabbitMQ.GetChannel()
	//		if err != nil {
	//			common.PrintOrLog(err.Error())
	//		}
	//		defer func(){
	//			ch.Cancel()
	//		}()
	//		err = global.RabbitMQ.Publish(ch,exChange,key,go_tool.GetDateTimeStr(t))
	//		if err != nil {
	//			common.PrintOrLog(err.Error())
	//		}
	//	}
	//}()
}
