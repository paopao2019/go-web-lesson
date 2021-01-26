package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 中间件
/*
需要掌握:
1. 从语法上讲什么是中间件， 哪些业务需要使用中间件
	是HandlerFunc即可 , 一般是公用业务使用中间件 如 日志、登入认证、权限校验、耗时统计
2. 如何注册中间件 全局和局部路由
	r.Use(m1,...) 全局注册
	r.GET("/index", m1, ...)  路由局部注册
3. c.Next()和c.Abort() 和 c.Abort() + return的区别 用图描述清楚
 */

// 花费时间的中间件
func m1(c *gin.Context) {
	fmt.Printf("m1 in\n")
	now :=  time.Now()
	c.Next()
	cost := time.Since(now)
	fmt.Printf("cost time: %v\n", cost)
	fmt.Printf("m1 out\n")
	//c.Abort()
}

// 中间件m2
func m2(c *gin.Context) {
	fmt.Printf("m2 in\n")
	c.Next()
	//c.Abort()   // 阻止后续HandlerFunc处理 (注意不是当前HandlerFunc后面的语句, 不要理解错了)
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "中间件m2 Abort",
	//})
	//return
	fmt.Printf("m2 out\n")
}

// 中间件认证 一般会用一个闭包处理 这样里面就可以做一些其他处理
func authMiddleware() gin.HandlerFunc {
	// 连接数据库
	// 或者一些其他的准备工作
	return func(c *gin.Context) {
		// 处理认证
		// if 认证成功
		c.Set("name", "andy")  // 在gin上下文中定义变量
		c.Next()
		log.Print("认证成功")
		// else 认证失败
		//c.Abort()
	}
}

func main() {
	r := gin.Default() // 默认使用了 Logger和Recovery中间件
	//r := gin.New()   // 不使用任何中间件
	r.Use(m1, m2)   // 全局注册中间件
	r.GET("/index", func(c *gin.Context) {
		fmt.Printf("index page\n")
		c.JSON(http.StatusOK, gin.H{"msg": "index page"})
	})

	r.GET("/auth", authMiddleware(), func(c *gin.Context) {
		// 获取gin上下文中的变量
		name, ok := c.Get("name")
		if !ok {
			name = "jack default"
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "auth认证",
			"name": name,
		})
	})


	r.Run(":9090") // http启动server服务




}
