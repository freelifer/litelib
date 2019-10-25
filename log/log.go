package log

import (
	"log"
)

var DEBUG bool = false

func init() {
	log.SetPrefix("[ginplus] ")
}

func SetDebug() {
	DEBUG = true
}

func SetPrefix(prefix string) {
	log.SetPrefix(prefix)
}

func I(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func D(format string, v ...interface{}) {
	if DEBUG {
		log.Printf(format, v...)
	}
}
func E(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}
