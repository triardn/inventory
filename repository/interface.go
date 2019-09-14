package repository

import (
	"github.com/triardn/inventory/model"
)

type IPersonRepository interface {
	GetAllPerson() ([]model.Person, error)
	SavePerson(person model.Person) error
}
