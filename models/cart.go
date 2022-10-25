package model

import "gorm.io/gorm"

// Cart 购物车模型
type Cart struct {
	gorm.Model
	UserID    uint `gorm:"not nul"`
	ProductID uint `gorm:"not null"`
	BossID    uint `gorm:"not nul"`
	Num       uint `gorm:"not nul"`
	MaxNum    uint `gorm:"not nul"`
	Check     bool
}
