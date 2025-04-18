package repository

import (
	"gorm.io/gorm"
)

type IRepository[E any, I any] interface {
	GetOneByID(id I) (E, error)
	Insert(entity E) error
}

type Repository[E any, I any] struct {
	DB *gorm.DB
}

func (r *Repository[E, I]) Insert(entity *E) error {
	return r.DB.Create(entity).Error
}

func (r *Repository[E, I]) GetOneByID(id I) (E, error) {
	var result E

	err := r.DB.First(result, id).Error

	return result, err
}
