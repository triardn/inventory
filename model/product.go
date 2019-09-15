package model

type Product struct {
	ID        uint64 `gorm:"column:id"`
	Sku       string `gorm:"column:sku"`
	Name      string `gorm:"column:name"`
	Quantity  int    `gorm:"column:quantity"`
	CreatedAt int    `gorm:"column:created_at"`
	UpdatedAt int    `gorm:"column:updated_at"`
}

// TableName alisa name for product
func (Product) TableName() string {
	return "product"
}
