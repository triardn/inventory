package repository

import (
	"github.com/triardn/inventory/model"
)

type ProductRepository struct {
	RepositoryOption
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
