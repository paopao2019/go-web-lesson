package main

import (
	"fmt"
	"lesson30/core"
	"lesson30/global"
	"time"
)

/*
1. vipper 初始  v.SetConfigFile(config)
2. 和定义的结构体对应起来使用  -- 实际应用中别  v.Unmarshal(&global.GVA_CONFIG)
 */
func main() {
	global.GVA_VP = core.Viper()          // 初始化Viper

	// 1. 使用viper获取配置文件 如果存在与分隔的键路径匹配的键
	redisAddress := global.GVA_VP.GetString("redis.addr")
	fmt.Printf("%v\n", redisAddress)   //10.108.26.93:6379

	// 2. 使用结构 - 一般都是使用map结构体
	// 使用map的好处就是有提示 对于ide比较友好
	redisAddressMap := global.GVA_CONFIG.Redis.Addr
	fmt.Printf("%v\n", redisAddressMap)   //10.108.26.93:6379

	time.Sleep(time.Second * 120)

	// 运行过程中 更改下配置文件 查看配置文件是否变化
	fmt.Printf("%v\n", global.GVA_CONFIG.Redis.DB)   // 0->1
}
