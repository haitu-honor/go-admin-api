package system

import (
	"github.com/myadmin/project/global"
)

type JwtBlacklist struct {
	global.GAI_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
