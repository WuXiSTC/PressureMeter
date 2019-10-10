package util

import "log"

func Log(msg string) {
	if msg != "" {
		log.Println(msg)
	}
}

func LogE(err error) {
	if err != nil {
		log.Println(err)
	}
}
