package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/video"
)

// 定义video的service提供xkVideo的数据curd的操作

type XkVideoService struct {
}

func (xkBbsService *XkVideoService) CreateXkVideo(xkVideo video.XkVideo) (err error) {
	err = global.GVA_DB.Create(&xkVideo).Error
	return err
}
