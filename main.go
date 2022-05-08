package main

import (
	"jackson-gin/global"
	"jackson-gin/initialize"
)

func main() {
	// 初始化配置
	initialize.InitializeConfig()

	// 初始化日志
	global.App.Log = initialize.InitializeLog()

	// 初始化数据库
	global.App.DB = initialize.InitializeDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 初始化验证器
	initialize.InitializeValidator()

	// 初始化Redis
	global.App.Redis = initialize.InitializeRedis()

	// 初始化文件系统
	initialize.InitializeStorage()

	// 初始化计划任务
	initialize.InitializeCron()

	// 启动服务器
	initialize.RunServer()
}
