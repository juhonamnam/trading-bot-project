package logger

import (
	"log"
	"os"
)

type logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
}

type loggerStruct struct {
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

func (logger *loggerStruct) Debug(v ...interface{}) {
	logger.debugLogger.Println(v...)
}
func (logger *loggerStruct) Info(v ...interface{}) {
	logger.infoLogger.Println(v...)
}
func (logger *loggerStruct) Warn(v ...interface{}) {
	logger.warnLogger.Println(v...)
}
func (logger *loggerStruct) Error(v ...interface{}) {
	logger.errorLogger.Println(v...)
}

var VBS logger

func init() {
	flags := log.LstdFlags
	VBS = &loggerStruct{
		debugLogger: log.New(os.Stdout, "VBS Debug: ", flags),
		infoLogger:  log.New(os.Stdout, "VBS Info: ", flags),
		warnLogger:  log.New(os.Stdout, "VBS Warn: ", flags),
		errorLogger: log.New(os.Stdout, "VBS Error: ", flags),
	}
}
