package dao

import (
	. "business/common"
	"business/dao/model"
	"strings"
)

/**
 * 获取类别列表
 */
type ListConfigArgs struct {
	Keys []string
}

func ListConfig(args ListConfigArgs) (configList []model.Config) {
	keyIn := "('" + strings.Join(args.Keys, "','") + "') "
	err := DbEngine.SQL("select * from b_config where `key` in " + keyIn).Find(&configList)
	if err != nil {
		panic(NewDbErr(err))
	}
	return
}
