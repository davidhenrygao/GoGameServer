package main

import (
	"fmt"
	"github.com/davidhenrygao/GoGameServer/utils/log"
)

func initServer(cfg initCfg) {
	err := log.InitLogger(cfg.LogLevel, cfg.LogDir, cfg.LogFileName, cfg.Logstderr)
	if err != nil {
		panic(fmt.Sprintf("Initial Logger: %v.\n", err))
	}
	logger := log.GetLogger()
	logger.Debug("Initial Logger success: %s level.", "debug")
	logger.Info("Initial Logger success: %s level.", "info")
	logger.Warn("Initial Logger success: %s level.", "warning")
	logger.Error("Initial Logger success: %s level.", "error")
	logger.Fatal("Initial Logger success: %s level.", "fatal")
	logger.Panic("Initial Logger success: %s level.", "panic")
}
