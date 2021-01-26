package core

import "lesson25/initialize"

func RunWindowsServer() {
	// 初始化路由
	Router := initialize.Routers()
	// 初始化http 启动 run server
	Router.Run(":8080")
}
