package dao

import (
	. "business/common"
	"business/dao/model"
)

/**
 * 获取类别列表
 */
type ListConfigArgs struct {
	Keys []string
}

func ListConfig(args ListConfigArgs) (configList []model.Config) {
	err := DbEngine.SQL("select * from b_config where `key` in " + WhereInString(args.Keys)).Find(&configList)
	if err != nil {
		panic(NewDbErr(err))
	}
	return
}
