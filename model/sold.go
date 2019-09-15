package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Sold struct {
	ID        uint64 `gorm:"column:id"`
	ProductID uint64 `gorm:"column:product_id"`
	Quantity  int64  `gorm:"column:quantity"`
	Price     int64  `gorm:"column:price"`
	Total     int64  `gorm:"column:total"`
	Notes     string `gorm:"column:notes"`
	Created   int64  `gorm:"column:created"`
	Updated   int64  `gorm:"column:updated"`
	Product   Product
}

// TableName alisa name for sold
func (Sold) TableName() string {
	return "sold"
}

func (s *Sold) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("Updated", time.Now().Unix())
	return nil
}

func (s *Sold) BeforeCreate(scope *gorm.Scope) error {
	if s.Updated == 0 {
		scope.SetColumn("Updated", time.Now().Unix())
	}

	scope.SetColumn("Created", time.Now().Unix())
	return nil
}

func (s *Sold) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(s).Preload("Product").Find(s)
	return
}
