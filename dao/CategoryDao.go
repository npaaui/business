package dao

import (
	. "business/common"
	"business/dao/model"
)

/**
 * 获取类别列表
 */
func ListCategory(args *model.Category) (categoryList []model.Category) {
	err := DbEngine.
		Where("type = ?", args.Type).Find(&categoryList)
	if err != nil {
		panic(NewDbErr(err))
	}
	return
}
