package ginplus

import (
	"strconv"

	"github.com/freelifer/litelib/dao"
	"github.com/freelifer/litelib/log"
	. "github.com/freelifer/litelib/public"
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

type ControllerInterface interface {
	List(g *gin.Context)
	Get(g *gin.Context)
	Post(g *gin.Context)
	Put(g *gin.Context)
}

// 路由接口实现------------------------------------------------------------------
///----------------------------------------------------------------------------
///----------------------------------------------------------------------------
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

// Controller api输出接口----------------------------------------------------------
///----------------------------------------------------------------------------
///----------------------------------------------------------------------------

func (c *Controller) JSON(g *gin.Context, obj interface{}) {
	g.JSON(200, obj)
}

func (c *Controller) MethodNotAllowed(g *gin.Context) {
	g.JSON(405, gin.H{
		"error": "Method Not Allowed",
	})
}

func (c *Controller) BadRequestMustJson(g *gin.Context) {
	g.JSON(400, gin.H{
		"message": "Body should be a JSON object",
	})
}

func (c *Controller) BadRequestErrorJson(g *gin.Context) {
	g.JSON(400, gin.H{
		"message": "Body Problems parsing JSON",
	})
}

func (c *Controller) Forbidden(g *gin.Context, err error) {
	log.I("Forbidden, %s", err.Error())
	g.JSON(403, gin.H{
		"message": err.Error(),
	})
}

func (c *Controller) UnprocessableEntity(g *gin.Context, err error) {
	log.I("Validation Failed, %s", err.Error())
	g.JSON(422, gin.H{
		"message": "Validation Failed",
		"errors":  err.Error(),
	})
}
func (c *Controller) ServerInnerError(g *gin.Context, err error) {
	log.I("Server Inner Error, %s", err.Error())
	g.JSON(500, gin.H{
		"message": "Server Inner Error",
	})
}

// Controller 工具实现----------------------------------------------------------
///----------------------------------------------------------------------------
///----------------------------------------------------------------------------
// 获取page和pageSize
func (c *Controller) QueryPageInfo(g *gin.Context) (int, int) {
	return c.DefaultQueryForInt(g, "page", 1), c.DefaultQueryForInt(g, "pageSize", 10)
}

// DefaultQueryForInt returns the keyed url query value if it exists
func (c *Controller) DefaultQueryForInt(g *gin.Context, key string, defaultValue int) int {
	if value, ok := strconv.Atoi(g.Query(key)); ok == nil {
		return value
	}
	return defaultValue
}

// DefaultQueryForInt64 returns the keyed url query value if it exists
func (c *Controller) DefaultQueryForInt64(g *gin.Context, key string, defaultValue int64) int64 {
	if value, ok := strconv.ParseInt(g.Query(key), 10, 64); ok == nil {
		return value
	}
	return defaultValue
}

// DefaultParamForInt returns the keyed url param value if it exists
func (c *Controller) DefaultParamForInt(g *gin.Context, key string, defaultValue int) int {
	if value, ok := strconv.Atoi(g.Param(key)); ok == nil {
		return value
	}
	return defaultValue
}

// DefaultParamForInt64 returns the keyed url param value if it exists
func (c *Controller) DefaultParamForInt64(g *gin.Context, key string, defaultValue int64) int64 {
	if value, ok := strconv.ParseInt(g.Param(key), 10, 64); ok == nil {
		return value
	}
	return defaultValue
}

// Controller 数据收集----------------------------------------------------------
///----------------------------------------------------------------------------
///----------------------------------------------------------------------------

// 获取随机数
func (c *Controller) RandomString(g *gin.Context, n int) (string, bool) {
	value, err := RandomString(n)
	if err != nil {
		c.ServerInnerError(g, err)
		return "", false
	}
	return value, true
}

// 获取body体的json数据
// 统一获取数据, 并上报数据
func (c *Controller) ParseJsonParam(g *gin.Context, obj interface{}) bool {
	err := g.BindJSON(obj)
	if err != nil {
		c.BadRequestMustJson(g)
		return false
	}

	return true
}

// 空和重复检查
func (c *Controller) CheckEmptyAndExistByField(g *gin.Context, value string, bean interface{}, field string) bool {
	if !c.CheckEmptyByField(g, value, field) {
		return false
	}

	if !c.CheckAlreadExistByField(g, bean, field) {
		return false
	}
	return true
}

// 空检查
func (c *Controller) CheckEmptyByField(g *gin.Context, value string, field string) bool {
	err := dao.IsEmptyByField(value, field)
	if err != nil {
		c.UnprocessableEntity(g, err)
		return false
	}
	return true
}

// 重复检查
func (c *Controller) CheckAlreadExistByField(g *gin.Context, bean interface{}, field string) bool {
	err := dao.IsAlreadExistByField(bean, field)
	if err != nil {
		if IsErrAlreadExist(err) {
			c.UnprocessableEntity(g, err)
			return false
		}
		c.ServerInnerError(g, err)
		return false
	}
	return true
}

// ErrNotExist
// 获取并检查不存在错误
func (c *Controller) GetAndCheckNotExistByOtherField(g *gin.Context, bean interface{}, field string) bool {
	err := dao.GetBeanByOtherField(bean, field)
	if err != nil {
		if IsErrNotExist(err) {
			c.Forbidden(g, err)
			return false
		}
		c.ServerInnerError(g, err)
		return false
	}
	return true
}

// ErrNotExist
// 获取并检查不存在错误
func (c *Controller) GetAndCheckNotExist(g *gin.Context, id int64, bean interface{}, field string) bool {
	err := dao.GetBeanById(id, bean, field)
	if err != nil {
		if IsErrNotExist(err) {
			c.Forbidden(g, err)
			return false
		}
		c.ServerInnerError(g, err)
		return false
	}
	return true
}

// 创建对象
func (c *Controller) CreateBean(g *gin.Context, bean interface{}) bool {
	err := Insert(bean)

	if err != nil {
		c.ServerInnerError(g, err)
		return false
	}

	return true
}

// 检查error
func (c *Controller) ParseError(g *gin.Context, err error) bool {
	if err != nil {
		c.ServerInnerError(g, err)
		return false
	}
	return true

}
