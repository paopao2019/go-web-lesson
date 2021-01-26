package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "GET",
	})
}

func main() {
	// 使用默认 返回引擎实例
	r := gin.Default()
	// 设置路由

	// RESTFUL 风格
	r.GET("/hello", sayHello)
	r.POST("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	// 引擎启动
	r.Run("127.0.0.1:9090")
}
