package repository

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"

	"github.com/triardn/inventory/common"
	"github.com/triardn/inventory/config"
)

type RepositoryOption struct {
	DB        *gorm.DB
	Cache     *redis.Pool
	Logger    *common.APILogger
	AppConfig *config.Config
}

type Repository struct {
	Person      IPersonRepository
	Product     IProductRepository
	Order       IOrderRepository
	OrderDetail IOrderDetailRepository
	Restock     IRestockRepository
	Sold        ISoldRepository
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) SetPersonRepository(personRepository IPersonRepository) {
	r.Person = personRepository
}

func (r *Repository) SetProductRepository(productRepository IProductRepository) {
	r.Product = productRepository
}

func (r *Repository) SetOrderRepository(orderRepository IOrderRepository) {
	r.Order = orderRepository
}

func (r *Repository) SetOrderDetailRepository(orderDetailRepository IOrderDetailRepository) {
	r.OrderDetail = orderDetailRepository
}

func (r *Repository) SetRestockRepository(restockRepository IRestockRepository) {
	r.Restock = restockRepository
}

func (r *Repository) SetSoldRepository(soldRepository ISoldRepository) {
	r.Sold = soldRepository
}
