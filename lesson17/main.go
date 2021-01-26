package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("404.html")
	// 路由组
	// 1. 用于多个业务线 每个业务线一个组
	// 2. 或者用于api的版本
	// 习惯性一对{}包裹同组的路由，这只是为了看着清晰

	// 视频路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/pop", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"url": "/video/pop",
			})
		})
		videoGroup.GET("/comedy", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"url": "/video/comedy",
			})
		})
		videoGroup.GET("/free", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"url": "/video/free",
			})
		})
	}

	// 用户路由组
	userGroup := r.Group("/user")
	{
		userGroup.GET("/list", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"url": "/user/list"})
		})
		userGroup.POST("/detail", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"url": "/user/detail"})
		})
	}

	// 默认404路由 更改
	//为没有配置处理函数的路由添加处理程序，默认情况下它返回404代码，下面的代码为没有匹配到路由的请求都返回
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.html", nil)
	})
	r.Run(":9090")  // 启动http服务

}
