package worker

import (
	"github.com/Deansquirrel/go-tool"
	"github.com/Deansquirrel/goRabbitMQDemo/common"
	"github.com/Deansquirrel/goRabbitMQDemo/global"
	"github.com/Deansquirrel/goRabbitMQDemo/object"
	"time"
)

func RefreshCurrConfig(config *object.SysConfig) error {
	global.SysConfig = config
	global.IsDebug = config.Total.IsDebug
	if global.RabbitMQ != nil {
		global.RabbitMQ.Close()
	}
	global.RabbitMQ = nil
	global.RabbitMQ = go_tool.NewRabbitMQ(
		config.RabbitMQ.User,
		config.RabbitMQ.Pwd,
		config.RabbitMQ.Server,
		config.RabbitMQ.Port,
		config.RabbitMQ.VirtualHost, time.Second*10, time.Millisecond*500, 3, time.Second*5)
	err := rabbitMqInit()
	if err != nil {
		return err
	}
	return nil
}

func rabbitMqInit() error {
	r := global.RabbitMQ
	var err error

	conn, err := r.GetConn()
	if err != nil {
		return err
	}
	defer func() {
		_ = conn.Close()
	}()
	//==================================================================================================================
	_ = r.ExchangeDelete(conn, "exchangeFanOut", true, true)
	_ = r.ExchangeDelete(conn, "exchangeDirect", true, true)
	_ = r.ExchangeDelete(conn, "exchangeTopic", true, true)
	_ = r.ExchangeDelete(conn, "exchangeHeaders", true, true)

	err = r.ExchangeDeclare(conn, "exchangeFanOut", go_tool.ExchangeFanout, true, false, false, false)
	if err != nil {
		return err
	}
	err = r.ExchangeDeclare(conn, "exchangeDirect", go_tool.ExchangeDirect, true, false, false, false)
	if err != nil {
		return err
	}
	err = r.ExchangeDeclare(conn, "exchangeTopic", go_tool.ExchangeTopic, true, false, false, false)
	if err != nil {
		return err
	}
	err = r.ExchangeDeclare(conn, "exchangeHeaders", go_tool.ExchangeHeaders, true, false, false, false)
	if err != nil {
		return err
	}
	//==================================================================================================================
	_ = r.QueueDelete(conn, "A", true, true, true)
	_ = r.QueueDelete(conn, "B", true, true, true)
	_ = r.QueueDelete(conn, "C", true, true, true)
	_ = r.QueueDelete(conn, "D", true, true, true)

	err = r.QueueDeclare(conn, "A", true, false, false, false)
	if err != nil {
		return err
	}

	err = r.QueueDeclare(conn, "B", true, false, false, false)
	if err != nil {
		return err
	}

	err = r.QueueDeclare(conn, "C", true, false, false, false)
	if err != nil {
		return err
	}

	err = r.QueueDeclare(conn, "D", true, false, false, false)
	if err != nil {
		return err
	}
	//==================================================================================================================
	err = r.QueueBind(conn, "A", "", "exchangeFanOut", false)
	if err != nil {
		return err
	}
	err = r.QueueBind(conn, "B", "", "exchangeFanOut", false)
	if err != nil {
		return err
	}
	err = r.QueueBind(conn, "C", "", "exchangeFanOut", false)
	if err != nil {
		return err
	}
	//==================================================================================================================
	err = r.AddProducer("")
	if err != nil {
		return err
	}
	//==================================================================================================================
	err = r.AddConsumer("qA", "A", handleA)
	if err != nil {
		return err
	}

	err = r.AddConsumer("qB", "B", handleB)
	if err != nil {
		return err
	}

	err = r.AddConsumer("qC", "C", handleC)
	if err != nil {
		return err
	}

	err = r.AddConsumer("qD", "D", handleD)
	if err != nil {
		return err
	}
	//==================================================================================================================
	return nil
}

func handleA(msg string) {
	common.PrintOrLog(go_tool.GetDateTimeStr(time.Now()) + " " + "A" + " " + msg)
}

func handleB(msg string) {
	common.PrintOrLog(go_tool.GetDateTimeStr(time.Now()) + " " + "B" + " " + msg)
}

func handleC(msg string) {
	common.PrintOrLog(go_tool.GetDateTimeStr(time.Now()) + " " + "C" + " " + msg)
}

func handleD(msg string) {
	common.PrintOrLog(go_tool.GetDateTimeStr(time.Now()) + " " + "D" + " " + msg)
}
