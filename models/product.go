package model

import (
	cache "gin-mall/cache"
	"gorm.io/gorm"
	"strconv"
)

// 商品模型
type Product struct {
	gorm.Model
	Name          string `gorm:"size:255;index"`
	CategoryID    uint   `gorm:"not null"`
	Title         string
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossID        int
	BossName      string
	BossAvatar    string
}

// View 获取点击数
func (product *Product) View() uint64 {
	//var ctx context.Context
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 商品游览
func (product *Product) AddView() {
	//var ctx context.Context
	// 增加视频点击数
	cache.RedisClient.Incr(cache.ProductViewKey(product.ID))
	// 增加排行点击数
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID)))
}
