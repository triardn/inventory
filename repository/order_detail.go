package repository

import (
	"github.com/triardn/inventory/model"
)

type OrderDetailRepository struct {
	RepositoryOption
}

func NewOrderDetailRepository(option RepositoryOption) *OrderDetailRepository {
	orderDetailRepo := &OrderDetailRepository{}
	orderDetailRepo.RepositoryOption = option

	return orderDetailRepo
}

func (odr *OrderDetailRepository) GetAllOrderDetail() (orderDetails []model.OrderDetail, err error) {
	err = odr.DB.
		Preload("Order").
		Preload("Product").
		Find(&orderDetails).
		Error
	if err != nil {
		return nil, err
	}

	return
}

func (odr *OrderDetailRepository) GetDetailByOrderID(orderID uint64) (orderDetails []model.OrderDetail, err error) {
	err = odr.DB.
		Preload("Product").
		Where("order_id = ?", orderID).
		Find(&orderDetails).
		Error
	if err != nil {
		return nil, err
	}

	return
}
