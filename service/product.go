package service

import (
	"github.com/triardn/inventory/common"
	"github.com/triardn/inventory/model"
	"github.com/triardn/inventory/repository"
)

type ProductService struct {
	repository repository.IProductRepository
	ServiceOption
}

func NewProductService(productRepository repository.IProductRepository, logger *common.APILogger) *ProductService {
	productService := &ProductService{}
	productService.Logger = logger
	productService.repository = productRepository
	return productService
}

func (prs *ProductService) GetAllProduct() (products []model.Product, err error) {
	return prs.repository.GetAllProduct()
}

func (prs *ProductService) GetProductDetail(id uint64) (oroduct model.Product, err error) {
	return prs.repository.GetProductDetail(id)
}

func (prs *ProductService) UpdateProduct(product *model.Product, payload map[string]interface{}) (err error) {
	return prs.repository.UpdateProduct(product, payload)
}

func (prs *ProductService) GetProductIDBySKU(sku string) (uint64, error) {
	return prs.repository.GetProductIDBySKU(sku)
}

func (prs *ProductService) CreateProduct(product model.Product) (model.Product, error) {
	return prs.repository.CreateProduct(product)
}
