package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Product struct {
	ID       uint64 `gorm:"column:id"`
	Sku      string `gorm:"column:sku"`
	Name     string `gorm:"column:name"`
	Quantity int64  `gorm:"column:quantity"`
	Created  int64  `gorm:"column:created"`
	Updated  int64  `gorm:"column:updated"`
}

// TableName alisa name for product
func (Product) TableName() string {
	return "product"
}

func (p *Product) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("Updated", time.Now().Unix())
	return nil
}

func (p *Product) BeforeCreate(scope *gorm.Scope) error {
	if p.Updated == 0 {
		scope.SetColumn("Updated", time.Now().Unix())
	}

	scope.SetColumn("Created", time.Now().Unix())
	return nil
}
