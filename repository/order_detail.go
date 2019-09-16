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

func (odr *OrderDetailRepository) CreateOrderDetail(orderDetailPayload model.OrderDetail) (model.OrderDetail, error) {
	err := odr.DB.
		Create(&orderDetailPayload).
		Error
	if err != nil {
		return model.OrderDetail{}, err
	}

	return orderDetailPayload, nil
}

func (odr *OrderDetailRepository) GetTotalOrderedProduct(start int64, end int64) int64 {
	var result Result
	err := odr.DB.
		Table("order_detail").
		Select("SUM(quantity) as total").
		Where("created between ? and ?", start, end).
		Scan(&result).
		Error
	if err != nil {
		return 0
	}

	return int64(result.Total)
}
