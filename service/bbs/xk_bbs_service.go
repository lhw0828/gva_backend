package bbsservice

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// 定义bbs的service提供xkbbs的数据curd的操作
type XkBbsService struct {
}

//@author: lhw
//@function: CreateXkBbs
//@description: 创建文章
//@param: xkBbs bbs.XkBbs
//@return: err error

func (xkBbsService *XkBbsService) CreateXkBbs(xkBbs *bbs.XkBbs) (err error) {
	err = global.GVA_DB.Create(&xkBbs).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除客户
//@param: e model.ExaCustomer
//@return: err error

func (xkBbsService *XkBbsService) DeleteXkBbs(xkBbs *bbs.XkBbs) (err error) {
	err = global.GVA_DB.Delete(&xkBbs).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateExaCustomer
//@description: 更新客户
//@param: e *model.ExaCustomer
//@return: err error

func (xkBbsService *XkBbsService) UpdateXkBbs(xkBbs *bbs.XkBbs) (err error) {
	err = global.GVA_DB.Save(xkBbs).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExaCustomer
//@description: 获取客户信息
//@param: id uint
//@return: customer model.ExaCustomer, err error

func (xkBbsService *XkBbsService) GetXkBbs(id uint) (xkBbs *bbs.XkBbs, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&xkBbs).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomerInfoList
//@description: 分页获取客户列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (xkBbsService *XkBbsService) GetXkBbsInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&bbs.XkBbs{})

	var XkBbsList []bbs.XkBbs
	err = db.Count(&total).Error

	if err != nil {
		return XkBbsList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&XkBbsList).Error
	}
	return XkBbsList, total, err
}
