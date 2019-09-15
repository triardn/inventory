package service

import (
	"github.com/triardn/inventory/common"
	"github.com/triardn/inventory/model"
	"github.com/triardn/inventory/repository"
)

type RestockService struct {
	repository  repository.IRestockRepository
	productRepo repository.IProductRepository
	ServiceOption
}

func NewRestockService(restockRepository repository.IRestockRepository, productRepository repository.IProductRepository, logger *common.APILogger) *RestockService {
	restockService := &RestockService{}
	restockService.Logger = logger
	restockService.repository = restockRepository
	restockService.productRepo = productRepository

	return restockService
}

func (rs *RestockService) GetAllRestockData() (restocks []model.Restock, err error) {
	return rs.repository.GetAllRestockData()
}

func (rs *RestockService) CreateRestockData(payloadRestock model.Restock) (restock model.Restock, err error) {
	restock, err = rs.repository.CreateRestockData(payloadRestock)
	if err != nil {
		return
	}

	// update product quantity
	updateProductPayload := make(map[string]interface{})
	updateProductPayload["quantity"] = restock.Product.Quantity + restock.ReceivedQuantity

	err = rs.productRepo.UpdateProduct(&restock.Product, updateProductPayload)
	if err != nil {
		return
	}

	return
}
