package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取API path参数
// 注意 URL的匹配不要冲突
func main() {
	r := gin.Default()

	r.GET("/api/:name", func(c *gin.Context) {
		// 获取path参数
		nameV := c.Param("name")  // 返回string类型
		c.String(http.StatusOK, nameV)
	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"年份": year,
			"月份": month,
		})
	})

	// 启动服务
	r.Run(":9090")
}
