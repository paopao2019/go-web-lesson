package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 301 重定向  这个卡看不到301日志默认gin中 直接重定向到了百度
	r.GET("/baidu", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	// 还有一种特殊的 像是重定向 其实只是跳转 内部处理函数转移

	r.GET("/a", func(c *gin.Context) {
		// 修改URL路径
		c.Request.URL.Path = "/b"
		r.HandleContext(c) // 继续后续的处理  re-enter 上下文
	})

	r.GET("/b", func(c *gin.Context) {
		c.String(http.StatusOK, "b")
	})

	// 启动服务
	r.Run(":9090")
}
