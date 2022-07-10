package initialize

import (
	"os"

	"github.com/myadmin/project/global"
	"github.com/myadmin/project/model/example"
	"github.com/myadmin/project/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Gorm 选择连接的数据库类型，初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.GAI_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	// case "pgsql":
	// 	return GormPgSql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表
// Author SliverHorn
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// 系统模块表
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},

		// 示例模块表
		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},
	)
	if err != nil {
		global.GAI_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GAI_LOG.Info("register table success")
}
