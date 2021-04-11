package service

import (
	. "business/common"
	"business/dao"
	"business/dao/model"
)

type CategoryService struct{}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (s *CategoryService) ListCategory(category *model.Category) (data *RespList) {
	count, list := dao.ListCategory(category)
	data = NewRespList(count, list)
	return
}
