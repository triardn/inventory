package service

import (
	"math/rand"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/triardn/inventory/common"
)

var randomSource rand.Source = rand.NewSource(time.Now().UnixNano())
var random *rand.Rand = rand.New(randomSource)

type ServiceOption struct {
	Cache  *redis.Pool
	Logger *common.APILogger
}

type Service struct {
	Person      IPersonService
	Product     IProductService
	Order       IOrderService
	OrderDetail IOrderDetailService
	Restock     IRestockService
	Sold        ISoldService
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SetPersonService(ps IPersonService) {
	s.Person = ps
}

func (s *Service) SetProductService(prs IProductService) {
	s.Product = prs
}

func (s *Service) SetOrderService(os IOrderService) {
	s.Order = os
}

func (s *Service) SetOrderDetailService(ods IOrderDetailService) {
	s.OrderDetail = ods
}

func (s *Service) SetRestockService(rs IRestockService) {
	s.Restock = rs
}

func (s *Service) SetSoldService(ss ISoldService) {
	s.Sold = ss
}
