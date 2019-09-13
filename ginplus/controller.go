package ginplus

import (
	. "github.com/freelifer/litelib/public"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Page     int
	PageSize int
}

type ControllerInterface interface {
	List(g *gin.Context)
	Get(g *gin.Context)
	Post(g *gin.Context)
	Put(g *gin.Context)
}

func (c *Controller) List(g *gin.Context) {
	g.JSON(405, gin.H{
		"error": "Method Not Allowed",
	})
}

func (c *Controller) Get(g *gin.Context) {
	g.JSON(405, gin.H{
		"error": "Method Not Allowed",
	})
}

func (c *Controller) Post(g *gin.Context) {
	g.JSON(405, gin.H{
		"error": "Method Not Allowed",
	})
}

func (c *Controller) Put(g *gin.Context) {
	g.JSON(405, gin.H{
		"error": "Method Not Allowed",
	})
}

// 获取page和pageSize
func (c *Controller) QueryDefaultParam(g *gin.Context) {
	c.Page = DefaultQueryForInt(g, "page", 1)
	c.PageSize = DefaultQueryForInt(g, "pageSize", 10)
}
