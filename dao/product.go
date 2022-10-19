package dao

import (
	"context"
	"gin-mall/models"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

// CreateProduct 创建商品
func (dao *ProductDao) CreateProduct(product *model.Product) (err error) {
	err = dao.DB.Model(&model.Product{}).Create(&product).Error
	return
}

// ListProductByCondition 获取商品列表
func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, page model.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Preload("Category").Where(condition).
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&products).Error
	return
}

// CountProductByCondition 根据情况获取商品的数量
func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}
