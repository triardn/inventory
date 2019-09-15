package repository

import (
	"github.com/triardn/inventory/model"
)

type OrderRepository struct {
	RepositoryOption
}

func NewOrderRepository(option RepositoryOption) *OrderRepository {
	orderRepo := &OrderRepository{}
	orderRepo.RepositoryOption = option

	return orderRepo
}

func (or *OrderRepository) GetAllOrder() (orders []model.Order, err error) {
	err = or.DB.
		Find(&orders).
		Error
	if err != nil {
		return nil, err
	}

	return
}

func (or *OrderRepository) GetOrderByID(id uint64) (order model.Order, err error) {
	err = or.DB.
		Where("id = ?", id).
		Find(&order).
		Error
	if err != nil {
		return
	}

	return
}

func (or *OrderRepository) GetOrderIDByInvoice(invoice string) (orderID uint64, err error) {
	order := model.Order{}

	err = or.DB.
		Select("id").
		Where("invoice = ?", invoice).
		Find(&order).
		Error
	if err != nil {
		return
	}

	return order.ID, nil
}
