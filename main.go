package main

import (
	"github.com/Deansquirrel/goRabbitMQDemo/common"
	"github.com/Deansquirrel/goRabbitMQDemo/global"
	"github.com/Deansquirrel/goRabbitMQDemo/worker"
)

func main() {
	//==================================================================================================================
	err := refreshConfig()
	if err != nil {
		common.PrintAndLog("加载配置时遇到错误:" + err.Error())
		return
	}
	common.PrintOrLog("程序启动")
	defer common.PrintOrLog("程序退出")
	//==================================================================================================================
	w, err := worker.NewTestWorker()
	if err != nil {
		common.PrintAndLog(err.Error())
		return
	}
	w.TestPublish()
	//==================================================================================================================
}

func refreshConfig() error {
	config, err := common.GetSysConfig("config.toml")
	if err != nil {
		return err
	}
	global.SysConfig = config
	common.RefreshCurrConfig(config)
	return nil
}
