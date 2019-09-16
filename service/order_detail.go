package service

import (
	"github.com/triardn/inventory/common"
	"github.com/triardn/inventory/model"
	"github.com/triardn/inventory/repository"
)

type OrderDetailService struct {
	repository repository.IOrderDetailRepository
	ServiceOption
}

func NewOrderDetailService(orderRepository repository.IOrderDetailRepository, logger *common.APILogger) *OrderDetailService {
	orderService := &OrderDetailService{}
	orderService.Logger = logger
	orderService.repository = orderRepository
	return orderService
}

func (ods *OrderDetailService) GetAllOrderDetail() (orderDetails []model.OrderDetail, err error) {
	return ods.repository.GetAllOrderDetail()
}

func (ods *OrderDetailService) GetDetailByOrderID(orderID uint64) (orderDetails []model.OrderDetail, err error) {
	return ods.repository.GetDetailByOrderID(orderID)
}

func (ods *OrderDetailService) CreateOrderDetail(orderDetails []model.OrderDetail) ([]model.OrderDetail, error) {
	var newOrderDetails []model.OrderDetail
	for _, detail := range orderDetails {
		newDetail, err := ods.repository.CreateOrderDetail(detail)
		if err != nil {
			return nil, err
		}

		newOrderDetails = append(newOrderDetails, newDetail)
	}

	return newOrderDetails, nil
}
