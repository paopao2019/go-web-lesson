package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// querystring
// GET请求 URL 参数通过 DefaultQuery 或 Query 方法获取
func main() {
	r := gin.Default()
	// 获取query参数
	// 如URL: http://127.0.0.1:9090/web?name=abc&age=10
	r.GET("/web", func(c *gin.Context) {
		// 从url中接收参数的函数也有很多

		//name := c.Query("name")  // 获取url中的query参数key为name,返回string类型数据, 如果没有key则返回空""
		name := c.DefaultQuery("name", "杨超越")  // 没有参数，则提供默认值
		age := c.Query("age")
		height,ok := c.GetQuery("height") // 返回两个值 不存在key布尔值为false
		if !ok {
			fmt.Printf("height key未传入\n")
			height = "1.80"
		}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age": age,
			"height": height,
		})
	})


	// 获取表单参数通过 PostForm 方法获取 -> 第12课
	// 获取api path参数 api 参数通过Context的Param方法来获取 -> 第13课
	// 启动服务
	r.Run(":9090")
}
