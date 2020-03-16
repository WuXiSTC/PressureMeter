package util

import "log"

type LoggerConfig struct {
	Logger func(string)
	Error  func(err error)
}

var option = LoggerConfig{
	Logger: func(s string) { log.Println(s) },
	Error:  func(err error) { log.Println(err) },
}

func SetLogger(o LoggerConfig) {
	option = o
}

func Log(msg string) {
	if msg != "" {
		option.Logger(msg)
	}
}

func LogE(err error) {
	if err != nil {
		option.Error(err)
	}
}
