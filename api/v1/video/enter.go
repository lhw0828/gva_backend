package videoApi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	XkVideoApi
}

var (
	videoService = service.ServiceGroupApp.XkVideoServiceGroup.XkVideoService
)
