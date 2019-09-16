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

func (or *OrderRepository) GetAllOrderWithTimeFrame(start int64, end int64) (orders []model.Order, err error) {
	err = or.DB.
		Where("created between ? and ?", start, end).
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

func (or *OrderRepository) GetTotalTurnover(start int64, end int64) int64 {
	var result Result
	err := or.DB.
		Table("Order").
		Select("SUM(total) as total").
		Where("created between ? and ?", start, end).
		Scan(&result).
		Error
	if err != nil {
		return 0
	}

	return int64(result.Total)
}

func (or *OrderRepository) GetCountOrder(start int64, end int64) int64 {
	var result int64
	err := or.DB.
		Table("Order").
		Select("invoice").
		Where("created between ? and ?", start, end).
		Count(&result).
		Error
	if err != nil {
		return 0
	}

	return result
}
