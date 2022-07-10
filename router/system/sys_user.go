package system

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// userRouter := Router.Group("user")
	// userRouterWithoutRecord := Router.Group("user")
	// baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi

}
