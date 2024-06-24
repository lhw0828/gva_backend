package bbsapi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// 定义api接口
type XkBbsApi struct{}

func (e *XkBbsApi) GetXkBbs(c *gin.Context) {
	var xkBbs bbs.XkBbs
	err := c.ShouldBindJSON(&xkBbs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data, err := bbsService.GetXkBbs(xkBbs.ID)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithDetailed(data, "创建成功", c)
}

func (e *XkBbsApi) GetXkBbsDetail(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	id := c.Param("id")
	// 这个是用来获取?age=123
	//age := c.Query("age")
	parseUint, _ := strconv.ParseUint(id, 10, 64)
	data, err := bbsService.GetXkBbs(uint(parseUint))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}
