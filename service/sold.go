package service

import (
	"github.com/triardn/inventory/common"
	"github.com/triardn/inventory/model"
	"github.com/triardn/inventory/repository"
)

type SoldService struct {
	repository  repository.ISoldRepository
	productRepo repository.IProductRepository
	ServiceOption
}

func NewSoldService(soldRepository repository.ISoldRepository, productRepository repository.IProductRepository, logger *common.APILogger) *SoldService {
	soldService := &SoldService{}
	soldService.repository = soldRepository
	soldService.productRepo = productRepository
	soldService.Logger = logger

	return soldService
}

func (ss *SoldService) GetAllSoldProducts() ([]model.Sold, error) {
	return ss.repository.GetAllSoldProduct()
}

func (ss *SoldService) CreateSoldRecord(payloadSold model.Sold) (sold model.Sold, err error) {
	sold, err = ss.repository.CreateSoldData(payloadSold)
	if err != nil {
		return
	}

	// update product quantity
	updateProductPayload := make(map[string]interface{})
	updateProductPayload["quantity"] = sold.Product.Quantity - sold.Quantity

	err = ss.productRepo.UpdateProduct(&sold.Product, updateProductPayload)
	if err != nil {
		return
	}

	return
}
