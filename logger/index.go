package logger

import (
	"log"
	"os"

	"github.com/juhonamnam/trading-bot-project/env"
)

type logger struct {
	Debug loggerInterface
	Info  loggerInterface
	Error loggerInterface
}

type loggerInterface interface {
	Println(v ...any)
	Printf(format string, v ...any)
}

type dummyLoggerStruct struct{}

func (d *dummyLoggerStruct) Println(_ ...any)          {}
func (d *dummyLoggerStruct) Printf(_ string, _ ...any) {}

var dummyLogger = dummyLoggerStruct{}

var VBS logger

func init() {
	flags := log.LstdFlags
	if env.IsProduction {
		vbsLogFile, err := os.OpenFile("./logs/vbs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(err.Error())
		}
		vbsErrorLogFile, err := os.OpenFile("./logs/vbs.error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(err.Error())
		}
		VBS = logger{
			Debug: &dummyLogger,
			Info:  log.New(vbsLogFile, "VBS Info: ", flags),
			Error: log.New(vbsErrorLogFile, "VBS Error: ", flags),
		}

	} else {
		VBS = logger{
			Debug: log.New(os.Stdout, "VBS Debug: ", flags),
			Info:  log.New(os.Stdout, "VBS Info: ", flags),
			Error: log.New(os.Stdout, "VBS Error: ", flags),
		}
	}
}
