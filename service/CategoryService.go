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
	list := dao.ListCategory(category)
	data = NewRespList(len(list), list)
	return
}
