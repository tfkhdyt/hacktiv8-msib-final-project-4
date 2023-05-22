package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Type              string    `gorm:"not null"`
	SoldProductAmount uint      `gorm:"default:0"`
	Products          []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
