package repository

import (
	"context"
	"internship-mini-project/internal/domain"

	"gorm.io/gorm"
)

type TokoRepository interface {
	Create(ctx context.Context, toko *domain.Toko) error
	Update(ctx context.Context, toko *domain.Toko) error
}

type TokoRepositoryImpl struct {
	DB *gorm.DB
}

// Create implements TokoRepository.
func (t *TokoRepositoryImpl) Create(ctx context.Context, toko *domain.Toko) error {
	return t.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(toko).Error; err != nil {
			return err
		}
		return nil
	})
}

// Update implements TokoRepository.
func (t *TokoRepositoryImpl) Update(ctx context.Context, toko *domain.Toko) error {
	panic("unimplemented")
}

func NewTokoRepository(db *gorm.DB) TokoRepository {
	return &TokoRepositoryImpl{DB: db}
}
