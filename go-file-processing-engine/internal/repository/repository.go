package repository

import (
	"gorm.io/gorm"
)

type Repository[E any, I any] struct {
	db *gorm.DB
}

func (r *Repository[E, I]) GetByID(id I) (E, error) {

	var result E

	err := r.db.First(result, id).Error

	return result, err
}

func (r *Repository[E, I]) Insert(entity E) error {
	return r.db.Create(entity).Error
}

func (r *Repository[E, I]) Update(entity E) error {
	return r.db.Updates(entity).Error
}

func (r *Repository[E, I]) Delete(entity E) (*E, error) {

	if err := r.db.Delete(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}
