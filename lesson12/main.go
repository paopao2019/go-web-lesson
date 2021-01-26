package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


// 获取表单参数通过 PostForm 方法获取
func main() {
	r := gin.Default()
	// 解析html文件(模板)
	r.LoadHTMLFiles("login.html", "upload.html")

	r.GET("/login", func(c *gin.Context) {
		// 渲染html(模板)文件
		c.HTML(http.StatusOK, "login.html", "ok" )
	})

	r.POST("/login", func(c *gin.Context) {
		// 获取表单post数据 如果key没有 返回空字符串
		username := c.PostForm("username")
		password := c.PostForm("password")

		// 如果表达没有提交post参数则使用默认
		//email := c.DefaultPostForm("email", "xxx@google.com")
		// 或者使用GetPostForm 得到两个返回值 如果key不存在布尔值为false
		email, ok := c.GetPostForm("email")
		if !ok {
			email = "xxx@test.com"
			fmt.Printf("未提供email参数\n")
		}
		fmt.Printf("用户: %s, 密码:%s, 邮箱:%s", username, password, email)

		// 渲染html(模板)文件
		c.HTML(http.StatusOK,"upload.html", gin.H{
			"username": username,
			"password": password,
			"email": email,
		})

	})
	// 启动服务
	r.Run(":9090")
}
