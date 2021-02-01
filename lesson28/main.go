package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"log"
	"net/http"
)

// https://github.com/mojocn/base64Captcha
//https://blog.csdn.net/zhanyia/article/details/105311794
//DriverDigit config for captcha-engine-digit.  // 数据是存在内存中的，所有可以校验前端发来的验证码
var store = base64Captcha.DefaultMemStore

// base64Captcha.DefaultMemStore是默认的:10240 个 过期时间10分钟

/*
创建验证码驱动
分为5种不同类型的图形验证码
分别是什么样子如: https://captcha.mojotv.cn/
dight 数字验证码  -> 演示这种
audio 语音验证码
string 字符验证码
math 数学验证码(加减乘除)
chinese中文验证码-有bug
*/
func Captcha(c *gin.Context) {
	//字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		log.Println("获取验证码失败..", err)
		c.JSON(http.StatusOK, gin.H{"message": "获取验证码成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":   "获取验证码成功",
			"CaptchaId": id,
			"PicPath":   b64s,
		})

	}
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		a := 1
		b := 2
		var e int
		e = a + b
		fmt.Printf("a+b: %d\n", e)
		c.String(http.StatusOK, "ok")
	})
	r.POST("/base/captcha", Captcha)

	// 如发送 http://127.0.0.1:9090/base/verify?captchaId=Fh970kCUVnQHaVg6B0sE&captcha=634647
	r.GET("/base/verify", func(c *gin.Context) {
		CaptchaId := c.Query("captchaId")
		Captcha := c.Query("captcha")

		fmt.Printf("id:%s 字符:%s\n", CaptchaId, Captcha)
		// http://tool.chinaz.com/tools/imgtobase/ 还原图片测试这个接口
		// 同时在内存清理掉这个图片
		if store.Verify(CaptchaId, Captcha, true) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "验证码成功",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": "验证码错误",
			})
		}
	})
	r.Run(":9090")
}
