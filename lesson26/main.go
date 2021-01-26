package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"lesson26/middleware"
	"log"
	"net/http"
	"time"
)

// 复杂的 JWT处理
/*
1. create
2. parse
3. 加BufferTime 原理
 */

func main() {

	r := gin.Default()
	r.GET("/createToken", func(c *gin.Context) {

		j := &middleware.JWT{SigningKey: []byte("qmPlus")} // 唯一签名
		claims := middleware.CustomClaims{
			UUID:        uuid.NewV4(),
			ID:          1,
			NickName:    "andy",
			Username:    "Jack",
			AuthorityId: "888",
			BufferTime:  86400, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
			StandardClaims: jwt.StandardClaims{
				NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
				ExpiresAt: time.Now().Unix() + 604800, // 过期时间 7天  配置文件
				Issuer:    "qmPlus",                                              // 签名的发行者
			},
		}
		token, err := j.CreateToken(claims)
		if err != nil {
			log.Println("获取token失败",  err)
			c.JSON(http.StatusOK, gin.H{
				"msg": "获取token失败",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": "登入成功",
			"data": map[string]interface{}{
				"User": "UserInfo字典 数据是来源于数据库中的user信息 Module",
				"Token": token,
				"ExpiresAt": claims.StandardClaims.ExpiresAt * 1000,
			},
		})

	})

	r.GET("/parseToken", func(c *gin.Context) {
		tokenString := c.Request.Header.Get("token")
		// 解析 token
		j := middleware.NewJWT()
		claims, err := j.ParseToken(tokenString)
		if err != nil {
			if err == middleware.TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"msg": "授权已过期",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}

		c.Set("claims", claims)
		userInfo, _ := c.Get("claims")
		c.JSON(http.StatusOK, gin.H{
			"msg": "successful",
			"userInfo": userInfo,
		})
	})

	r.Run(":9090")
}
