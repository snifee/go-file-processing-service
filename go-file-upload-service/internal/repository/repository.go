package repository

import "gorm.io/gorm"

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Create(entity *T) error {
	return r.DB.Create(entity).Error
}
