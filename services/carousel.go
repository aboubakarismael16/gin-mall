package services

import (
	"context"
	"gin-mall/dao"
	"gin-mall/pkg/e"
	util "gin-mall/pkg/utils"
	"gin-mall/serializer"
)

type ListCarouselsService struct{}

func (service *ListCarouselsService) List() serializer.Response {
	code := e.SUCCESS
	carouselsCtx := dao.NewCarouselDao(context.Background())
	carousels, err := carouselsCtx.ListAddress()
	if err != nil {
		util.LogrusObj.Info("err ", err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))
}
