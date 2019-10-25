package log_test

import (
	"testing"
	"github.com/freelifer/litelib/log"
)

func TestLog(t *testing.T) {
	log.SetDebug()
	log.SetPrefix("[kzhu] ")
	log.I("I am info log, %s", "TOM")
	log.D("I am debug log, %s", "TOM")
}