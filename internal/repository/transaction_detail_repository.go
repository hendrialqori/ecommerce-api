package repository

import (
	"context"
	"internship-mini-project/internal/domain"

	"gorm.io/gorm"
)

type TrxDetailRepository interface {
	Create(ctx context.Context, trx *domain.TransactionDetail) error
}

type TrxDetailRepositoryImpl struct {
	DB *gorm.DB
}

func (t *TrxDetailRepositoryImpl) Create(ctx context.Context, trx *domain.TransactionDetail) error {
	return t.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(trx).Error; err != nil {
			return err
		}
		return nil
	})
}

func NewTrxDetailRepository(DB *gorm.DB) TrxDetailRepository {
	return &TrxDetailRepositoryImpl{DB: DB}
}
