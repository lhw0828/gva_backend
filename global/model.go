package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey;comment:主键ID" json:"id" form:"id"` // 主键ID
	CreatedAt time.Time      `gorm:"type:datetime(0);autoCreateTime;comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"type:datetime(0);autoUpdateTime;comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"-"` // 删除时间
}
