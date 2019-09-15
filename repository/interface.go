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
}
