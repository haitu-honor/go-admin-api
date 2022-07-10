package initialize

import (
	"log"
	"os"
	"time"

	"github.com/myadmin/project/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	m := global.GAI_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	// mysql 配置
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN 数据库连接信息
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据mysql版本自动配置
	}
	// 生成 gorm 映射实例
	if db, err := gorm.Open(mysql.New(mysqlConfig), GormConf.GormConfig()); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns) //空闲中的最大连接数
		sqlDB.SetMaxOpenConns(m.MaxOpenConns) //打开到数据库的最大连接数
		return db
	}
}

/**===========================  gorm 配置  ===================================*/

type DBBASE interface {
	GetLogMode() string // 该方法在config/gorm_mysql.go 里被实现
}

var GormConf = new(_gorm)

type _gorm struct{}

// gorm 自定义配置
func (g *_gorm) GormConfig() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true} // 迁移时禁用外键约束
	// 自定义一个 logger 来替换默认的 logger
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond, // 慢查询时间
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	var logMode DBBASE
	// 选择使用哪种数据库类型的配置，并赋给 logMode
	switch global.GAI_CONFIG.System.DbType {
	case "mysql":
		logMode = &global.GAI_CONFIG.Mysql
	// case "pgsql":
	// 	logMode = &global.GAI_CONFIG.Pgsql
	// 	break
	default:
		logMode = &global.GAI_CONFIG.Mysql
	}
	// 选择gorm日志级别
	switch logMode.GetLogMode() {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}
