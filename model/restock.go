package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Restock struct {
	ID               uint64 `gorm:"column:id"`
	ProductID        uint64 `gorm:"column:product_id"`
	OrderedQuantity  int64  `gorm:"column:ordered_quantity"`
	ReceivedQuantity int64  `gorm:"column:received_quantity"`
	Price            int64  `gorm:"column:price"`
	Total            int64  `gorm:"column:total"`
	ReceiptNumber    string `gorm:"column:receipt_number"`
	Notes            string `gorm:"column:notes"`
	Created          int64  `gorm:"column:created"`
	Updated          int64  `gorm:"column:updated"`
	Product          Product
}

// TableName alisa name for restock
func (Restock) TableName() string {
	return "restock"
}

func (r *Restock) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("Updated", time.Now().Unix())
	return nil
}

func (r *Restock) BeforeCreate(scope *gorm.Scope) error {
	if r.Updated == 0 {
		scope.SetColumn("Updated", time.Now().Unix())
	}

	scope.SetColumn("Created", time.Now().Unix())
	return nil
}

func (r *Restock) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(r).Preload("Product").Find(r)
	return
}
