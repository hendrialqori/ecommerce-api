package repository

import (
	"context"
	"internship-mini-project/internal/domain"

	"gorm.io/gorm"
)

type ProductPhotoRepository interface {
	Create(ctx context.Context, photo *domain.ProductPhoto) error
	Delete(ctx context.Context, id uint) error
}

type ProductPhotoRepositoryImpl struct {
	DB *gorm.DB
}

// Create implements ProductPhotoRepository.
func (p *ProductPhotoRepositoryImpl) Create(ctx context.Context, photo *domain.ProductPhoto) error {
	return p.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(photo).Error; err != nil {
			return err
		}
		return nil
	})
}

// Delete implements ProductPhotoRepository.
func (p *ProductPhotoRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return p.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&domain.ProductPhoto{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func NewProductPhotoRepository(db *gorm.DB) ProductPhotoRepository {
	return &ProductPhotoRepositoryImpl{
		DB: db,
	}
}
