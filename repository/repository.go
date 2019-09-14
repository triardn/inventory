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
	Person         IPersonRepository
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) SetPersonRepository(personRepository IPersonRepository) {
	r.Person = personRepository
}
