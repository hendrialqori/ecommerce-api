package repository

import (
	"context"
	"internship-mini-project/internal/domain"
	"internship-mini-project/internal/model"

	"gorm.io/gorm"
)

type TokoRepository interface {
	FindAll(ctx context.Context, query *model.QueryParams) ([]*domain.Toko, error)
	FindById(ctx context.Context, id int) (*domain.Toko, error)
	Current(ctx context.Context, userId int) (*domain.Toko, error)
	Create(ctx context.Context, toko *domain.Toko) error
	Update(ctx context.Context, toko *domain.Toko) error
}

type TokoRepositoryImpl struct {
	DB *gorm.DB
}

func (t *TokoRepositoryImpl) Current(ctx context.Context, userId int) (*domain.Toko, error) {
	toko := &domain.Toko{}
	if err := t.DB.WithContext(ctx).Where("id_user = ?", userId).First(toko).Error; err != nil {
		return nil, err
	}
	return toko, nil
}

// FindById implements TokoRepository.
func (t *TokoRepositoryImpl) FindById(ctx context.Context, id int) (*domain.Toko, error) {
	toko := &domain.Toko{}
	if err := t.DB.WithContext(ctx).Where("id = ?", id).First(toko).Error; err != nil {
		return nil, err
	}
	return toko, nil
}

func (t *TokoRepositoryImpl) FindAll(ctx context.Context, query *model.QueryParams) ([]*domain.Toko, error) {
	tokos := make([]*domain.Toko, 0)
	db := t.DB.WithContext(ctx).Model(&domain.Toko{})

	if query.Nama != nil {
		db = db.Where("nama_toko LIKE ?", "%"+*query.Nama+"%")
	}

	if err := db.Offset((query.Page - 1) * query.Limit).Limit(query.Limit).Find(&tokos).Error; err != nil {
		return nil, err
	}

	return tokos, nil
}

// Update implements TokoRepository.
func (t *TokoRepositoryImpl) Update(ctx context.Context, toko *domain.Toko) error {
	return t.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existing domain.Toko

		if err := tx.Where("id = ? AND id_user = ?", toko.ID, toko.IDUser).
			First(&existing).Error; err != nil {
			return err
		}

		if err := tx.Model(&existing).Updates(toko).Error; err != nil {
			return err
		}
		return nil
	})
}

func (t *TokoRepositoryImpl) Create(ctx context.Context, toko *domain.Toko) error {
	return t.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(toko).Error; err != nil {
			return err
		}
		return nil
	})
}

func NewTokoRepository(db *gorm.DB) TokoRepository {
	return &TokoRepositoryImpl{DB: db}
}
