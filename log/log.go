package log

import (
	"log"
)

func I(format string, v ...interface{}) {
	log.Printf(format, v)
}

func E(format string, v ...interface{}) {
	log.Fatalf(format, v)
}
