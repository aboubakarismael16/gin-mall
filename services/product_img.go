package services

import (
	"context"
	"gin-mall/dao"
	"gin-mall/serializer"
	"strconv"
)

type ListProductImg struct {
}

func (service *ListProductImg) ListImg(ctx context.Context, pId string) serializer.Response {
	productImgDao := dao.NewProductImgDao(ctx)
	productId, _ := strconv.Atoi(pId)
	productImgs, _ := productImgDao.ListProductImg(uint(productId))
	return serializer.BuildListResponse(serializer.BuildProductImgs(productImgs), uint(len(productImgs)))
}
