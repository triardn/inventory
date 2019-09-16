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
	GetProductIDBySKU(sku string) (uint64, error)
	CreateProduct(product model.Product) (model.Product, error)
	PopulateExportData() (data [][]string, grandTotalPrice int64, err error)
	ProductStatistics() map[string]int
}

type IOrderService interface {
	GetAllOrder() (orders []model.Order, err error)
	GetOrderByID(id uint64) (order model.Order, err error)
	GetOrderIDByInvoice(invoice string) (orderID uint64, err error)
	PopulateExportData(start string, end string) (data [][]string, grandTotalProfit int64, err error)
	ReportStatistics(start string, end string) (statistics map[string]int64)
	CreateOrder(payloadOrder model.Order) (order model.Order, err error)
	UpdateOrder(order *model.Order, payload map[string]interface{}) (err error)
}

type IOrderDetailService interface {
	GetAllOrderDetail() (orderDetails []model.OrderDetail, err error)
	GetDetailByOrderID(orderID uint64) (orderDetails []model.OrderDetail, err error)
	CreateOrderDetail(orderDetails []model.OrderDetail) ([]model.OrderDetail, error)
}

type IRestockService interface {
	GetAllRestockData() (restocks []model.Restock, err error)
	CreateRestockData(payloadRestock model.Restock) (restock model.Restock, err error)
}

type ISoldService interface {
	GetAllSoldProducts() ([]model.Sold, error)
	CreateSoldRecord(payloadSold model.Sold) (sold model.Sold, err error)
}
