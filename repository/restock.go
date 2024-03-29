package repository

import (
	"github.com/triardn/inventory/model"
)

type RestockRepository struct {
	RepositoryOption
}

func NewRestockRepository(option RepositoryOption) *RestockRepository {
	restockRepo := &RestockRepository{}
	restockRepo.RepositoryOption = option

	return restockRepo
}

func (rr *RestockRepository) GetAllRestockData() (restock []model.Restock, err error) {
	err = rr.DB.
		Preload("Product").
		Find(&restock).
		Error
	if err != nil {
		return nil, err
	}

	return
}

func (rr *RestockRepository) CreateRestockData(restock model.Restock) (model.Restock, error) {
	err := rr.DB.
		Create(&restock).
		Error
	if err != nil {
		return model.Restock{}, err
	}

	return restock, nil
}

func (rr *RestockRepository) GetAveragePriceByProductID(productID uint64) int {
	var result ResultFloat
	err := rr.DB.
		Table("Restock").
		Select("SUM(total) as total, SUM(received_quantity) as qty").
		Where("product_id = ?", productID).
		Scan(&result).
		Error
	if err != nil {
		return 0
	}

	if result.Total != 0 && result.Qty != 0 {
		return int(result.Total) / result.Qty
	}

	return 0
}
