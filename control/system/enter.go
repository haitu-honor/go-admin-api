package system

import "github.com/myadmin/project/service"

type ApiGroup struct {
	JwtApi
	BaseApi
}

var (
	jwtService  = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
