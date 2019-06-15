package public

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

var (
	DB struct {
		Engine *xorm.Engine
		Tables []interface{}
	}

	Gin struct {
		g *gin.Engine
	}
)
