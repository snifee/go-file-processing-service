package repository

import (
	"go-file-processing-engine/internal/model/entity"

	"gorm.io/gorm"
)

type ProductRepository struct {
	Repository[entity.Product, int]
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		Repository: Repository[entity.Product, int]{
			db: db,
		},
	}
}

func (r *ProductRepository) TransactionUpdateProduct(productName string, id int) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		return tx.Update("product_name", productName).Error
	})
	return err
}
