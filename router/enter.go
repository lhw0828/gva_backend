package router

import (
	bbsRouter "github.com/flipped-aurora/gin-vue-admin/server/router/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/video"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	XkBbs   bbsRouter.RouterGroup
	XkVideo video.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
