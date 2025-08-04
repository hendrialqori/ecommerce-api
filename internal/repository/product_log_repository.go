package repository

import (
	"context"
	"fmt"
	"internship-mini-project/internal/domain"

	"gorm.io/gorm"
)

type ProductLogRepository interface {
	Create(ctx context.Context, productLog *domain.ProductLog) error
}

type ProductLogImpl struct {
	DB *gorm.DB
}

// Create implements ProductLog.
func (p *ProductLogImpl) Create(ctx context.Context, productLog *domain.ProductLog) error {
	return p.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(productLog).Error; err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println(productLog)
		return nil
	})
}

func NewProductLogRepository(DB *gorm.DB) ProductLogRepository {
	return &ProductLogImpl{
		DB: DB,
	}
}
