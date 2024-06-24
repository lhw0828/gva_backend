package videoApi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/video"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 定义api接口

type XkVideoApi struct{}

func (e *XkVideoApi) CreateXkVideo(c *gin.Context) {
	var videoModel video.XkVideo
	err := c.ShouldBindJSON(&videoModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = videoService.CreateXkVideo(videoModel)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}
