package repository

import (
	"go-file-processing-engine/internal/model/entity"

	"gorm.io/gorm"
)

type OrderDetailRepository struct {
	Repository[entity.Order]
}

func NewOrderDetailRepository(db *gorm.DB) *OrderDetailRepository {
	return &OrderDetailRepository{
		Repository: Repository[entity.Order]{
			DB: db,
		},
	}
}
