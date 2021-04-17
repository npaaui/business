package service

import (
	. "business/common"
	"business/dao"
	"business/dao/model"
)

func (s *ConfigService) ListCategory(category *model.Category) (data *RespList) {
	count, list := dao.ListCategory(category)
	data = NewRespList(count, list)
	return
}

func (s *ConfigService) InfoCategory(category *model.Category) bool {
	has := category.Info()
	return has
}
