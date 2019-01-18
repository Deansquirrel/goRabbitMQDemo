package common

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/Deansquirrel/go-tool"
	"github.com/Deansquirrel/goRabbitMQDemo/global"
	"github.com/Deansquirrel/goRabbitMQDemo/object"
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
