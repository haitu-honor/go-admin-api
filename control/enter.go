package v1

import (
	"github.com/myadmin/project/control/example"
	"github.com/myadmin/project/control/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
