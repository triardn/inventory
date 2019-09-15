package service

import (
	"github.com/triardn/inventory/common"
	"github.com/triardn/inventory/model"
	"github.com/triardn/inventory/repository"
)

type OrderService struct {
	repository repository.IOrderRepository
	ServiceOption
}

func NewOrderService(orderRepository repository.IOrderRepository, logger *common.APILogger) *OrderService {
	orderService := &OrderService{}
	orderService.Logger = logger
	orderService.repository = orderRepository
	return orderService
}

func (os *OrderService) GetAllOrder() (orders []model.Order, err error) {
	return os.repository.GetAllOrder()
}

func (os *OrderService) GetOrderByID(id uint64) (order model.Order, err error) {
	return os.repository.GetOrderByID(id)
}

func (os *OrderService) GetOrderByInvoice(invoice string) (order model.Order, err error) {
	return os.repository.GetOrderByInvoice(invoice)
}
