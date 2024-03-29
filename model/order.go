package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	ID      uint64 `gorm:"column:id"`
	Invoice string `gorm:"column:invoice"`
	Total   int64  `gorm:"column:total"`
	Notes   string `gorm:"column:notes"`
	Created int64  `gorm:"column:created"`
}

// TableName alisa name for order
func (Order) TableName() string {
	return "order"
}

func (p *Order) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Created", time.Now().Unix())
	return nil
}
