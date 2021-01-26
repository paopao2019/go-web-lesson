package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 学习如何使用gin 返回json数据
//方法1. 使用 gin.H 自己拼接json 其实就是map
//方法2. 使用结构体 还要学会灵活使用tag对结构体字段做定制化操作
func main() {
	r := gin.Default()

	// 1. map方式 自定义数据
	r.GET("/json1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "andy",
			"age": 18,
			"height": 1.88,
		})
	})
	data := map[string]interface{} {
		"name": "andy",
		"age": 18,
		"height": 1.88,
	}
	r.GET("/json2", func(c *gin.Context) {
		c.JSON(http.StatusOK, data)
	})

	// 2. 返回结构体类型的数据
	type userInfo struct {
		Name string `json:"name"`
		Age int	`json:"age"`
		Height interface{} `json:"height"`
	}

	data2 := userInfo{
		Name: "andy",
		Age: 18,
		Height: 1.88,
	}
	r.GET("/json3", func(c *gin.Context) {
		c.JSON(http.StatusOK, data2)
	})

	// 启动http服务
	r.Run(":9090")


}
