package dao

import (
	. "business/common"
	"business/dao/model"
)

/**
 * 获取类别列表
 */
func ListCategory(args *model.Category) (int, []model.Category) {
	var categoryList []model.Category
	count, err := DbEngine.
		Where("type = ?", args.Type).FindAndCount(&categoryList)
	if err != nil {
		panic(NewDbErr(err))
	}
	return int(count), categoryList
}
