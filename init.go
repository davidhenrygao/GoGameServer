package main

import (
	"fmt"
	"github.com/davidhenrygao/GoGameServer/module/log"
)

func initServer(cfg initCfg) {
	err := log.InitLogger(cfg.LogLevel, cfg.LogDir, cfg.LogFileName, cfg.Logstderr)
	if err != nil {
		panic(fmt.Sprintf("Initial Logger: %v.\n", err))
	}
	logger := log.GetLogger()
	logger.Info("Init Server...")
	logger.Info("Init logger success.")

}
