package main

import (
	"github.com/tarekbadrshalaan/goStuff/logger/zapPackaging/logger"
)

func main() {
	logger.InitiatLogger()
	defer logger.Sync()
	logger.Info("hi")
}
