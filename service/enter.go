package service

import (
	"github.com/myadmin/project/service/example"
	"github.com/myadmin/project/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
