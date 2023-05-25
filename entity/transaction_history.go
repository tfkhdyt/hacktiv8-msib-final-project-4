package entity

import "gorm.io/gorm"

type TransactionHistory struct {
	gorm.Model
	ProductID  uint
	UserID     uint
	Quantity   uint `gorm:"not null"`
	TotalPrice uint `gorm:"not null"`
}
