package ginplus

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Route(group gin.IRoutes, path string, ci ControllerInterface) {
	group.GET(fmt.Sprintf("/%s", path), ci.List)
	group.GET(fmt.Sprintf("/%s/:id", path), ci.Get)
	group.POST(fmt.Sprintf("/%s", path), ci.Post)
	group.PUT(fmt.Sprintf("/%s/:id", path), ci.Put)
}
