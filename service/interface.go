package service

import (
	"github.com/triardn/inventory/model"
)

type IPersonService interface {
	GetAllPerson() ([]model.Person, error)
	CreatePerson(person model.Person) error
}