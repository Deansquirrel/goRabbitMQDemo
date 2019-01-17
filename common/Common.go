package common

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/Deansquirrel/go-tool"
	"github.com/Deansquirrel/goRabbitMQDemo/global"
	"github.com/Deansquirrel/goRabbitMQDemo/object"
	"time"
)

func PrintAndLog(msg string) {
	fmt.Println(msg)
	if global.IsDebug {
		err := go_tool.Log(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func PrintOrLog(msg string) {
	if global.IsDebug {
		err := go_tool.Log(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(msg)
	}
}

func GetSysConfig(fileName string) (*object.SysConfig, error) {
	path, err := go_tool.GetCurrPath()
	if err != nil {
		return nil, err
	}
	var config object.SysConfig
	_, err = toml.DecodeFile(path+"\\"+fileName, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func RefreshCurrConfig(config *object.SysConfig) error {
	global.SysConfig = config
	global.IsDebug = config.Total.IsDebug
	global.RabbitMQ.Close()
	global.RabbitMQ = nil
	global.RabbitMQ = go_tool.NewRabbitMQ(
		config.RabbitMQ.User,
		config.RabbitMQ.Pwd,
		config.RabbitMQ.Server,
		config.RabbitMQ.Port,
		config.RabbitMQ.VirtualHost)
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

	err = r.QueueBind(conn, "A", "", "exchangeDirect", false)
	if err != nil {
		return err
	}
	err = r.QueueBind(conn, "B", "", "exchangeDirect", false)
	if err != nil {
		return err
	}

	err = r.AddProducer("testP", time.Second, time.Millisecond*100, time.Second)
	if err != nil {
		return err
	}
	return nil
}

func GetWxLastNo(cardNo string) (int, error) {
	if len(cardNo) != 15 {
		return 0, errors.New("号码长度应为15位")
	}
	b := []byte(cardNo)
	var r int
	for i := 0; i < len(b); i++ {
		r = r + int(b[i])
	}
	return r % 10, nil
}
