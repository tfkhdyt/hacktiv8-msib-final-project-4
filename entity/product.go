package entity

import (
	"fmt"

	"gorm.io/gorm"

	"hacktiv8-msib-final-project-4/pkg/errs"
)

type Product struct {
	gorm.Model
	Title                string `gorm:"not null"`
	Price                uint   `gorm:"not null"`
	Stock                uint   `gorm:"not null"`
	CategoryID           uint
	TransactionHistories []TransactionHistory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (p *Product) CheckStock(quantity uint) errs.MessageErr {
	if p.Stock < quantity {
		return errs.NewBadRequest(
			fmt.Sprintf(
				"Insufficient product stock. There are only %d items left in stock",
				p.Stock,
			),
		)
	}

	return nil
}
