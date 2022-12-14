package dao

import (
	"context"
	"gin-mall/models"
	"gorm.io/gorm"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{NewDBClient(ctx)}
}

func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}

// CreateProductImg 创建商品图片
func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) error {
	return dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error
}

// ListProductImgByProductId 根据商品id获取商品图片
func (dao *ProductImgDao) ListProductImgByProductId(pId uint) (products []*model.ProductImg, err error) {
	err = dao.DB.Model(&model.ProductImg{}).
		Where("product_id=?", pId).Find(&products).Error
	return
}

func (dao *ProductImgDao) ListProductImg(productId uint) (productImg []*model.ProductImg, err error) {
	err = dao.DB.Model(&model.ProductImg{}).Where("product_id=?", productId).Find(&productImg).Error
	return
}
