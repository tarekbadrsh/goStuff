package main

import (
	"github.com/tarekbadrshalaan/goStuff/logger/zapPackaging/logger"
)

func main() {
	logger.InitiatLogger()
	defer logger.Sync()
	logger.Info("hi")
	logger.Infof("Hi from %v", "my computer")
	logger.Warn("WarnWarnWarn")
	logger.Error("ErrorErrorError")
	logger.Fatal("FatalFatalFatal")
	logger.Infof("%v bye ", "my computer")
}
