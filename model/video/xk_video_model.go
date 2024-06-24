package video

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// XkVideo 结构体
type XkVideo struct {
	global.GVA_MODEL
	Title     string `json:"title" form:"title" gorm:"column:title;comment:标题;size:100;"`
	Cid       *int   `json:"cid" form:"cid" gorm:"column:cid;comment:分类ID;"`
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

func (XkVideo) TableName() string {
	return "xk_video"
}
