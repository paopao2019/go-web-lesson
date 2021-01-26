package main

import (
	"lesson25/core"
	"lesson25/global"
	"lesson25/initialize"

)

// 小项目 todo

/*
1. GET
2. POST 增
3. PUT  改
3. DELETE 删
 */



func main() {

	global.GVA_DB = initialize.Gorm()     // gorm连接数据库
	initialize.MysqlTables(global.GVA_DB)   // 数据库表初始化
	global.GAA_TEST = 12

	//// 表迁移
	//global.GVA_DB.AutoMigrate(&TODO{})

	// http server启动 和 初始化路由
	core.RunWindowsServer()
}
