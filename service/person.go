package service

import (
	"github.com/triardn/inventory/common"
	"github.com/triardn/inventory/model"
	"github.com/triardn/inventory/repository"
)

type PersonService struct {
	repository repository.IPersonRepository
	ServiceOption
}

func NewPersonService(personRepository repository.IPersonRepository, logger *common.APILogger) *PersonService {
	personService := &PersonService{}
	personService.Logger = logger
	personService.repository = personRepository
	return personService
}

func (ps *PersonService) GetAllPerson() ([]model.Person, error) {
	ps.Logger.Stdout.Println("Ini output log")
	ps.Logger.Stderr.Println("Ini error log")
	return ps.repository.GetAllPerson()
}

func (ps *PersonService) CreatePerson(person model.Person) error {
	return ps.repository.SavePerson(person)
}
