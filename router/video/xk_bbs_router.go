package video

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type XkVideoRouter struct{}

func (e *XkVideoRouter) InitXkVideoRouter(Router *gin.RouterGroup) {

	xkVideoRouterWithoutRecord := Router.Group("video")
	xkVideoApi := v1.ApiGroupApp.VideoApiGroup.XkVideoApi
	{
		//对外暴露的接口 http://localhost:8888/video/createXkVideo
		xkVideoRouterWithoutRecord.POST("createXkVideo", xkVideoApi.CreateXkVideo) // 新增video
	}
}
