package router

import (
	"github.com/gin-gonic/gin"
	"lesson25/api/v1"
)

// 初始化业务路由
func InitTODORouter(Router *gin.RouterGroup) {
	TODORouter := Router.Group("/v1/todo")
	{
		TODORouter.GET("/", v1.TODOGet)
		TODORouter.POST("/", v1.TODOPost)
		TODORouter.PUT("/:id", v1.TODOPut)
		TODORouter.DELETE("/:id", v1.TODODelete)
	}
}

