package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

// gin中的模板渲染
func main() {
	r := gin.Default()

	// 自定义函数 模板中可以使用的函数 如定义一个 去除默认的转义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})
	// 静态文件处理 就是增加路由
	r.Static("/static", "./static")


	// 解析模板
	//r.LoadHTMLFiles("templates/home/index.tmpl", "templates/index/index.tmpl")
	//如果模板中没有使用define定义模板名字。 默认的名字是base filename， 使用LoadHTMLGlob
	//可以看到 Loaded HTML Templates 的name名字
	r.LoadHTMLGlob("templates/**/*")

	// 渲染模板
	r.GET("/index/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
			"name": "我是index页面",
			"link": "<a href=http://www.liwenzhou.com>李文周站点</a>",
		})
	})

	r.GET("/home/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.tmpl", gin.H{
			"name": "我是home页面",
		})
	})

	// 启动服务
	err := r.Run(":9090")
	if err != nil {
		fmt.Printf("gin start server failed, err: %v\n", err)
	}


}
