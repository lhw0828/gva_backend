package service

import (
	bbsService "github.com/flipped-aurora/gin-vue-admin/server/service/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/video"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	XkBbsServiceGroup   bbsService.ServiceGroup
	XkVideoServiceGroup video.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
