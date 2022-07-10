package initialize

import (
	"fmt"

	"github.com/myadmin/project/global"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// 重写默认logger里 Writer接口的 Printf方法，实现自动的格式化打印日志
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.GAI_CONFIG.System.DbType {
	case "mysql":
		logZap = global.GAI_CONFIG.Mysql.LogZap
	// case "pgsql":
	// 	logZap = global.GAI_CONFIG.Pgsql.LogZap
	default:
		logZap = global.GAI_CONFIG.Mysql.LogZap
	}
	if logZap {
		global.GAI_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
