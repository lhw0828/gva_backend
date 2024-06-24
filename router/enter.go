package router

import (
	bbsrouter "github.com/flipped-aurora/gin-vue-admin/server/router/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	XkBbs   bbsrouter.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
