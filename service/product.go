package service

import (
	"strconv"

	"github.com/triardn/inventory/common"
	"github.com/triardn/inventory/model"
	"github.com/triardn/inventory/repository"
)

type ProductService struct {
	repository  repository.IProductRepository
	restockRepo repository.IRestockRepository
	ServiceOption
}

func NewProductService(productRepository repository.IProductRepository, restockRepository repository.IRestockRepository, logger *common.APILogger) *ProductService {
	productService := &ProductService{}
	productService.Logger = logger
	productService.repository = productRepository
	productService.restockRepo = restockRepository
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

func (prs *ProductService) PopulateExportData() (data [][]string, grandTotalPrice int64, err error) {
	products, err := prs.repository.GetAllProduct()
	for _, product := range products {
		averagePrice := prs.restockRepo.GetAveragePriceByProductID(product.ID)
		totalPrice := product.Quantity * int64(averagePrice)

		tempData := []string{product.Sku, product.Name, strconv.FormatInt(product.Quantity, 10), common.FormatCurrency("id_ID", int64(averagePrice), true), common.FormatCurrency("id_ID", totalPrice, true)}

		data = append(data, tempData)

		grandTotalPrice += totalPrice
	}

	return
}

func (prs *ProductService) ProductStatistics() map[string]int {
	var statistics = make(map[string]int)

	statistics["sku"] = prs.repository.GetTotalSku()
	statistics["stock"] = prs.repository.GetTotalStock()

	return statistics
}
