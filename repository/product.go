package repository

import (
	"github.com/triardn/inventory/model"
)

type ProductRepository struct {
	RepositoryOption
}

type Result struct {
	Total int
}

type ResultFloat struct {
	Total float64
	Qty   int
}

func NewProductRepository(option RepositoryOption) *ProductRepository {
	productRepo := &ProductRepository{}
	productRepo.RepositoryOption = option

	return productRepo
}

func (pr *ProductRepository) GetAllProduct() (products []model.Product, err error) {
	err = pr.DB.
		Find(&products).
		Error
	if err != nil {
		return
	}

	return
}

func (pr *ProductRepository) GetProductDetail(id uint64) (product model.Product, err error) {
	err = pr.DB.
		Where("id = ?", id).
		Find(&product).
		Error

	if err != nil {
		return
	}

	return
}

func (pr *ProductRepository) UpdateProduct(product *model.Product, payload map[string]interface{}) (err error) {
	err = pr.DB.
		Model(&product).
		Updates(payload).
		Error

	return
}

func (pr *ProductRepository) GetProductIDBySKU(sku string) (productID uint64, err error) {
	product := model.Product{}

	err = pr.DB.
		Select("id").
		Where("sku = ?", sku).
		Find(&product).
		Error
	if err != nil {
		return
	}

	return product.ID, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (model.Product, error) {
	err := pr.DB.
		Create(&product).
		Error
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (pr *ProductRepository) GetTotalSku() (totalSku int) {
	err := pr.DB.
		Table("Product").
		Select("DISTINC sku").
		Count(&totalSku).
		Error
	if err != nil {
		return
	}

	return
}

func (pr *ProductRepository) GetTotalStock() int {
	var result Result
	err := pr.DB.
		Table("Product").
		Select("SUM(quantity) as total").
		Scan(&result).
		Error
	if err != nil {
		return 0
	}

	return result.Total
}
