package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/myadmin/project/global"
	"github.com/myadmin/project/router"
)

// 初始化总路由
func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp.System
	// exampleRouter := router.RouterGroupApp.Example

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		// systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	// PrivateGroup := Router.Group("")
	// {
	// 	exampleRouter.InitExcelRouter(PrivateGroup)                 // 表格导入导出
	// 	exampleRouter.InitCustomerRouter(PrivateGroup)              // 客户路由
	// 	exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由

	// }
	global.GAI_LOG.Info("router register success")
	return Router
}
