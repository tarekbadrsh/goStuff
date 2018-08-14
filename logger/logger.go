package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/op/go-logging"
)

// Logger var to use
var Logger = logging.MustGetLogger("logger")

var logFileformat = logging.MustStringFormatter(
	`%{level:.4s},%{time:15:04:05.0},%{shortfile},%{message}`,
)

var logTerminalformat = logging.MustStringFormatter(
	`%{color}%{level:.4s},%{time:15:04:05.0},%{shortfile},%{message}%{color:reset}`,
)

func setTerminalLogger(level logging.Level) *logging.LeveledBackend {
	terminalLogger := logging.NewLogBackend(os.Stderr, "", 0)
	terminalformatted := logging.NewBackendFormatter(terminalLogger, logTerminalformat)
	terminallevelled := logging.AddModuleLevel(terminalformatted)
	terminallevelled.SetLevel(level, "")
	return &terminallevelled
}

func setFileLogger(level logging.Level, filePath string) *logging.LeveledBackend {
	fileName := fmt.Sprintf(filePath)
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}
	fileLogger := logging.NewLogBackend(f, "", 0)
	fileformatted := logging.NewBackendFormatter(fileLogger, logFileformat)
	filelevelled := logging.AddModuleLevel(fileformatted)
	filelevelled.SetLevel(level, "")
	return &filelevelled
}

// SetLogger :
func SetLogger(loggerType string, loggerLevel int64, logdir string) error {
	switch loggerType {
	case "f":
		if _, errIsNotExist := os.Stat(logdir); os.IsNotExist(errIsNotExist) {
			return errIsNotExist
		}
		go func() {
			for {
				now := time.Now()
				now.Add(time.Hour)
				day := now.Day()
				month := now.Month()
				year := now.Year()
				fileName := fmt.Sprintf("%v/%d-%d-%d.csv", logdir, year, month, day)
				fileLogger := setFileLogger(logging.Level(loggerLevel), fileName)
				logging.SetBackend(*fileLogger)
				waitFor := 24 - now.Hour()
				<-time.After(time.Duration(waitFor) * time.Hour)
			}
		}()
	case "t":
		terminalLogger := setTerminalLogger(logging.Level(loggerLevel))
		logging.SetBackend(*terminalLogger)
	case "a":
		terminalLogger := setTerminalLogger(logging.Level(loggerLevel))
		if _, errIsNotExist := os.Stat(logdir); os.IsNotExist(errIsNotExist) {
			return errIsNotExist
		}
		go func(terminalLogger *logging.LeveledBackend) {
			for {
				now := time.Now()
				now.Add(time.Hour)
				day := now.Day()
				month := now.Month()
				year := now.Year()
				fileName := fmt.Sprintf("%v/%d-%d-%d.csv", logdir, year, month, day)
				fileLogger := setFileLogger(logging.Level(loggerLevel), fileName)
				logging.SetBackend(*fileLogger, *terminalLogger)
				waitFor := 24 - now.Hour()
				<-time.After(time.Duration(waitFor) * time.Hour)
			}
		}(terminalLogger)
	case "n":
		logging.SetBackend()
	}
	return nil
}
