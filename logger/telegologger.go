package logger

import (
	"log"
	"os"

	"github.com/juhonamnam/telego"
	"github.com/juhonamnam/trading-bot-project/env"
)

type telegoLogger struct {
	debugLogger loggerInterface
	infoLogger  loggerInterface
	warnLogger  loggerInterface
	errorLogger loggerInterface
}

func (logger *telegoLogger) Debug(v ...interface{}) {
	logger.debugLogger.Println(v...)
}
func (logger *telegoLogger) Info(v ...interface{}) {
	logger.infoLogger.Println(v...)
}
func (logger *telegoLogger) Warn(v ...interface{}) {
	logger.warnLogger.Println(v...)
}
func (logger *telegoLogger) Error(v ...interface{}) {
	logger.errorLogger.Println(v...)
}

func GetTelegoLogger() telego.Logger {
	flags := log.LstdFlags
	if env.IsProduction {
		telegoLogFile, err := os.OpenFile("./logs/telego.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(err.Error())
		}
		telegoErrorLogFile, err := os.OpenFile("./logs/telego.error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(err.Error())
		}
		return &telegoLogger{
			debugLogger: &dummyLogger,
			infoLogger:  log.New(telegoLogFile, "Telego Info: ", flags),
			warnLogger:  log.New(telegoLogFile, "Telego WARN: ", flags),
			errorLogger: log.New(telegoErrorLogFile, "Telego ERROR: ", flags),
		}
	}
	return &telegoLogger{
		debugLogger: log.New(os.Stdout, "Telego Debug: ", flags),
		infoLogger:  log.New(os.Stdout, "Telego Info: ", flags),
		warnLogger:  log.New(os.Stdout, "Telego WARN: ", flags),
		errorLogger: log.New(os.Stdout, "Telego ERROR: ", flags),
	}
}
