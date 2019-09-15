package repository

import (
	"github.com/triardn/inventory/model"
)

type SoldRepository struct {
	RepositoryOption
}

func NewSoldRepository(option RepositoryOption) *SoldRepository {
	soldRepo := &SoldRepository{}
	soldRepo.RepositoryOption = option

	return soldRepo
}

func (sr *SoldRepository) GetAllSoldProduct() (soldProducts []model.Sold, err error) {
	err = sr.DB.
		Find(&soldProducts).
		Error
	if err != nil {
		return nil, err
	}

	return
}

func (sr *SoldRepository) CreateSoldData(sold model.Sold) (model.Sold, error) {
	err := sr.DB.
		Create(&sold).
		Error
	if err != nil {
		return model.Sold{}, err
	}

	return sold, nil
}
