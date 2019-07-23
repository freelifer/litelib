package litelib

import (
	_ "github.com/mattn/go-sqlite3"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLoadConfigFile(t *testing.T) {
	Convey("Load a single configuration file that does exist\n", t, func() {
		litelib := NewLiteLib()
		litelib.SetConfigPath("testdata/conf.ini")

		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong111",
			})
		})
		litelib.SetGinEngine(r)

		litelib.Run()
		// c, err := LoadConfigFile("testdata/conf.ini")
		// So(err, ShouldBeNil)
		// So(c, ShouldNotBeNil)
	})
}
