package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title                string `gorm:"not null"`
	Price                uint   `gorm:"not null"`
	Stock                uint   `gorm:"not null"`
	CategoryID           uint
	TransactionHistories []TransactionHistory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
