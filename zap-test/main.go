package main

import (
	"zap-test/logger"
	"zap-test/setting"

	"go.uber.org/zap"
)

func main() { // Initialize the logger
	if err := setting.ReadConfigFile(); err != nil {
		panic("Failed to read configuration file: " + err.Error())
	}
	var logConfig setting.ZapConfig = *setting.Conf
	logger.InitZap(&logConfig)
	defer logger.Logger.Sync() // Ensure all buffered log entries are flushed
	// fmt.Printf("%+v\n", setting.Conf)
	// fmt.Printf("%+v\n", setting.Conf.LogFile)
	var port int = 8000
	mode := "debug"
	logger.Logger.Info("Configuration loaded successfully", zap.Int("port", port), zap.String("mode", mode))
	// Start the server (not implemented in this example)
	// StartServer()
}
