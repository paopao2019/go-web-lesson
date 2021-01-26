package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


// 参数绑定  数据解析绑定
// 适合接收接口接受的很多数据
/*
模型绑定可以将请求体绑定给一个类型，目前支持绑定的类型有 JSON, XML 和标准表单数据 (foo=bar&boo=baz)。
要注意的是绑定时需要给字段设置绑定类型的标签。比如绑定 JSON 数据时，设置 json:"fieldname"。
使用绑定方法时，Gin 会根据请求头中 Content-Type 来自动判断需要解析的类型。

生产中 只需要使用 c.ShouldBind() 函数接即可满足需求
 */

type userInfo struct {
	// 需要通过反射找到 所以字段一定要大写
	Name string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
}
func main() {
	r := gin.Default()
	var u userInfo

	// 演示 http://127.0.0.1:9090/form?name=andy&password=123456
	r.GET("/form", func(c *gin.Context) {
		// 注意1. 一定要传入指针才能改变
		err := c.ShouldBind(&u)
		if err != nil {
			fmt.Printf("数据绑定失败, err: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "bad",
				"error": err.Error(),
			})
			return
		}
		fmt.Printf("绑定的用户数据是: %#v", u) //main.userInfo{Name:"andy", Password:"123456"}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})


	// 演示json post数据 可以使用post测试
	/*
	curl --location --request POST '127.0.0.1:9090/json' \
	--header 'Content-Type: application/json' \
	--data-raw '{
	    "user": "123",
	    "password": "123456789"
	}'
	 */
	r.POST("/json", func(c *gin.Context) {
		// 注意1. 一定要传入指针才能改变
		err := c.ShouldBind(&u)
		if err != nil {
			fmt.Printf("数据绑定失败, err: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "bad",
				"error": err.Error(),
			})
			return
		}
		fmt.Printf("绑定的用户数据是: %#v", u) //main.userInfo{Name:"", Password:"123456789"}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})



	// 启动运行http服务
	r.Run(":9090")


}
