package bbsRouter

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type XkBbsRouter struct{}

func (e *XkBbsRouter) InitXkBbsRouter(Router *gin.RouterGroup) {

	xkBbsRouterWithoutRecord := Router.Group("bbs")
	xkBbsApi := v1.ApiGroupApp.BbsApiGroup
	{
		//对外暴露的接口 http://localhost:8888/bbs/get?id=123
		xkBbsRouterWithoutRecord.GET("get", xkBbsApi.GetXkBbs)                 // 获取单一客户信息
		xkBbsRouterWithoutRecord.GET("getdetail/:id", xkBbsApi.GetXkBbsDetail) // 获取客户列表
	}
}
