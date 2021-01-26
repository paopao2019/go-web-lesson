package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"

	"net/http"
)

func main() {

	// 1. 数据来自 csv文件列表
	//e := casbin.NewEnforcer("./resource/rbac_model.conf", "./resource/policy.csv")

	// 数据来自mysql  https://github.com/casbin/gorm-adapter
	// 学习资料来自: https://casbin.org/docs/zh-CN/overview
	a, _ := gormadapter.NewAdapter("mysql", "root:p@ss1234@anji@tcp(10.108.26.60:3307)/test_db", true) // Your driver and data source.
	e, _ := casbin.NewEnforcer("./resource/rbac_model.conf", a)
	// Load the policy from DB.
	e.LoadPolicy()

	sub := "alice" // 想要访问资源的用户。
	obj := "data1" //要访问的资源。
	act := "read" // 用户对资源执行的操作。

	// 判断策略中是否存在
	success, _ := e.Enforce(sub, obj, act)

	// 添加策略 -> 会加入到数据中
	added, err := e.AddPolicy("eve", "data3", "read")

	if added {
		fmt.Printf("添加策略成功\n")
	} else {
		fmt.Printf("添加策略失败 err: %v\n", err)
	}

	// 批量添加 二维切片
	rules := [][] string {
		[]string {"jack", "data4", "read"},
		[]string {"katy", "data4", "write"},
		[]string {"leyo", "data4", "read"},
		[]string {"ham", "data4", "write"},
	}
	e.AddPolicies(rules)

	// 删除策略
	// removed, _ := e.RemovePolicy("alice", "data1", "read")
	added, err = e.AddPolicy("alice", "data1", "read")
	// 0 是固定的 不需要知道啥意思  内部  RemoveFilteredNamedPolicy 和这个是一样的
	removed, _ := e.RemoveFilteredPolicy(0, "alice")
	//
	if removed {
		log.Println("移除策略成功")
	} else {
		log.Println("移除策略失败")
	}

	if success {
		// 允许alice读取data1
		fmt.Printf("允许:%s 访问%s 资源\n", sub, obj)
	} else {
		// 拒绝请求，显示错误
		fmt.Printf("拒绝:%s 访问%s 资源\n", sub, obj)
	}

	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		fmt.Printf("URL is:%v\n", c.Request.URL)
		fmt.Printf("Path is:%v\n", c.Request.URL.Path)
		fmt.Printf("RequestURI is:%v\n",  c.Request.URL.RequestURI())

		c.String(http.StatusOK, "ok")
	})

	r.Run(":9090")
}
