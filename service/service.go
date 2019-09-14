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
	Person IPersonService
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SetPersonService(ps IPersonService) {
	s.Person = ps
}
