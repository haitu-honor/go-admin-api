package main

import (
	"github.com/myadmin/project/core"
	"github.com/myadmin/project/global"
	"github.com/myadmin/project/initialize"
	"go.uber.org/zap"
)

func main() {
	global.GAI_VP = core.Viper() // 初始化Viper
	global.GAI_LOG = core.Zap()  // 初始化zap日志库
	// zap 库自己提供的全局的 logger 是zap.S() 和 zap.L()
	zap.ReplaceGlobals(global.GAI_LOG) // 将全局的 logger 替换为我们通过配置定制的 logger
	global.GAI_DB = initialize.Gorm()  // gorm连接数据库
	if global.GAI_DB != nil {
		initialize.RegisterTables(global.GAI_DB) //初始化表
		db, _ := global.GAI_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
