package bbsapi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	XkBbsApi
}

var (
	bbsService = service.ServiceGroupApp.XkBbsServiceGroup.XkBbsService
)
