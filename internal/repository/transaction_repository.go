package repository

import (
	"context"
	"internship-mini-project/internal/domain"

	"gorm.io/gorm"
)

type TrxRepository interface {
	FindAll(ctx context.Context) (*[]domain.Transaction, error)
	FindByID(ctx context.Context, id int) (*domain.Transaction, error)
	Create(ctx context.Context, trx *domain.Transaction) error
}

type TrxRepositoryImpl struct {
	DB *gorm.DB
}

// FindByID implements TrxRepository.
func (t *TrxRepositoryImpl) FindByID(ctx context.Context, id int) (*domain.Transaction, error) {
	trx := &domain.Transaction{}
	if err := t.DB.WithContext(ctx).
		Preload("AlamatPengiriman").
		Preload("TransactionDetail").
		Preload("TransactionDetail.LogProduk").
		Preload("TransactionDetail.LogProduk.Category").
		Preload("TransactionDetail.LogProduk.Product.ProductPhoto").
		Preload("TransactionDetail.Toko").
		Where("id = ?", id).
		First(trx).Error; err != nil {
		return nil, err
	}
	return trx, nil
}

func (t *TrxRepositoryImpl) FindAll(ctx context.Context) (*[]domain.Transaction, error) {
	trx := &[]domain.Transaction{}

	if err := t.DB.WithContext(ctx).
		Preload("AlamatPengiriman").
		Preload("TransactionDetail").
		Preload("TransactionDetail.LogProduk").
		Preload("TransactionDetail.LogProduk.Category").
		Preload("TransactionDetail.LogProduk.Product.ProductPhoto").
		Preload("TransactionDetail.Toko").
		Find(trx).Error; err != nil {
		return nil, err
	}

	return trx, nil
}

func (t *TrxRepositoryImpl) Create(ctx context.Context, trx *domain.Transaction) error {
	return t.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(trx).Error; err != nil {
			return err
		}
		return nil
	})
}

func NewTrxRepository(DB *gorm.DB) TrxRepository {
	return &TrxRepositoryImpl{DB: DB}
}
