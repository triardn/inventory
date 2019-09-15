package service

import (
	"github.com/triardn/inventory/model"
)

type IPersonService interface {
	GetAllPerson() ([]model.Person, error)
	CreatePerson(person model.Person) error
}

type IProductService interface {
	GetAllProduct() (products []model.Product, err error)
	GetProductDetail(id uint64) (oroduct model.Product, err error)
	UpdateProduct(product *model.Product, payload map[string]interface{}) (err error)
}

type IOrderService interface {
	GetAllOrder() (orders []model.Order, err error)
	GetOrderByID(id uint64) (order model.Order, err error)
	GetOrderByInvoice(invoice string) (order model.Order, err error)
}

type IOrderDetailService interface {
	GetDetailByOrderID(orderID uint64) (orderDetails []model.OrderDetail, err error)
}
