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
	GetProductIDBySKU(sku string) (productID uint64, err error)
	CreateProduct(product model.Product) (model.Product, error)
	GetTotalSku() (totalSku int)
	GetTotalStock() int
}

type IOrderRepository interface {
	GetAllOrder() (orders []model.Order, err error)
	GetAllOrderWithTimeFrame(start int64, end int64) (orders []model.Order, err error)
	GetOrderByID(id uint64) (order model.Order, err error)
	GetOrderIDByInvoice(invoice string) (orderID uint64, err error)
	GetTotalTurnover(start int64, end int64) int64
	GetCountOrder(start int64, end int64) int64
	CreateOrder(orderPayload model.Order) (order model.Order, err error)
	UpdateOrder(order *model.Order, payload map[string]interface{}) (err error)
}

type IOrderDetailRepository interface {
	GetAllOrderDetail() (orderDetails []model.OrderDetail, err error)
	GetDetailByOrderID(orderID uint64) (orderDetails []model.OrderDetail, err error)
	GetTotalOrderedProduct(start int64, end int64) int64
	CreateOrderDetail(orderDetailPayload model.OrderDetail) (model.OrderDetail, error)
}

type IRestockRepository interface {
	GetAllRestockData() (restock []model.Restock, err error)
	CreateRestockData(restock model.Restock) (model.Restock, error)
	GetAveragePriceByProductID(productID uint64) int
}

type ISoldRepository interface {
	GetAllSoldProduct() (soldProducts []model.Sold, err error)
	CreateSoldData(sold model.Sold) (model.Sold, error)
}
