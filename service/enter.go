package service

import (
	bbsService "github.com/flipped-aurora/gin-vue-admin/server/service/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	XkBbsServiceGroup   bbsService.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
