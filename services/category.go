package services

import (
	"context"
	"gin-mall/dao"
	"gin-mall/pkg/e"
	util "gin-mall/pkg/utils"
	"gin-mall/serializer"
)

type ListCategoriesService struct {
}

func (service *ListCategoriesService) CategoryList(ctx context.Context) serializer.Response {
	code := e.SUCCESS
	categoryDao := dao.NewCategoryDao(ctx)
	categories, err := categoryDao.ListCategory()
	if err != nil {
		util.LogrusObj.Infoln(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCategories(categories),
	}
}
