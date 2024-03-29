package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type OrderDetail struct {
	ID        uint64 `gorm:"column:id"`
	OrderID   uint64 `gorm:"column:order_id"`
	ProductID uint64 `gorm:"column:product_id"`
	Price     int64  `gorm:"column:price"`
	Quantity  int64  `gorm:"column:quantity"`
	Total     int64  `gorm:"column:total"`
	Created   int64  `gorm:"column:created"`
	Order     Order
	Product   Product
}

// TableName alisa name for order_detail
func (OrderDetail) TableName() string {
	return "order_detail"
}

func (p *OrderDetail) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Created", time.Now().Unix())
	return nil
}

func (od *OrderDetail) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(od).Preload("Product").Find(od)
	return
}
