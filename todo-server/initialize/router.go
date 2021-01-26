package initialize

import (
	"github.com/gin-gonic/gin"
	"lesson25/middleware"
	"lesson25/router"
	"net/http"
)


func Routers() *gin.Engine {
	r := gin.Default()

	// 跨域的全局中间件
	r.Use(middleware.Cors())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "基于 gin Gorm的小项目 TODO待办事项")
	})

	PublicGroup := r.Group("")
	PublicGroup.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "公共路由，不需要做鉴权的路由")
	})

	PrivateGroup := r.Group("")
	{
		router.InitTODORouter(PrivateGroup)
	}

	return r


}
