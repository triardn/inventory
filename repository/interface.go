package repository

import (
	"github.com/triardn/inventory/model"
)

type IPersonRepository interface {
	GetAllPerson() ([]model.Person, error)
	SavePerson(person model.Person) error
}

type IProductRepository interface {
	GetAllProduct() (products []model.Product, err error)
	GetProductDetail(id uint64) (product model.Product, err error)
	UpdateProduct(product *model.Product, payload map[string]interface{}) (err error)
}

type IOrderRepository interface {
	GetAllOrder() (orders []model.Order, err error)
	GetOrderByID(id uint64) (order model.Order, err error)
	GetOrderByInvoice(invoice string) (order model.Order, err error)
}

type IOrderDetailRepository interface {
	GetAllOrderDetail() (orderDetails []model.OrderDetail, err error)
	GetDetailByOrderID(orderID uint64) (orderDetails []model.OrderDetail, err error)
}
