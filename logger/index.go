package logger

import (
	"log"
	"os"
)

type logger struct {
	Debug *log.Logger
	Info  *log.Logger
	Error *log.Logger
}

var VBS logger

func init() {
	flags := log.LstdFlags
	VBS = logger{
		Debug: log.New(os.Stdout, "VBS Debug: ", flags),
		Info:  log.New(os.Stdout, "VBS Info: ", flags),
		Error: log.New(os.Stdout, "VBS Error: ", flags),
	}
}
