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

func (ods *OrderDetailService) GetDetailByOrderID(orderID uint64) (orderDetails []model.OrderDetail, err error) {
	return ods.repository.GetDetailByOrderID(orderID)
}
