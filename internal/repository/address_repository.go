package repository

import (
	"context"
	"internship-mini-project/internal/domain"

	"gorm.io/gorm"
)

type AddressRepository interface {
	FindAll(ctx context.Context, userId int) ([]*domain.Address, error)
	FindById(ctx context.Context, id int) (*domain.Address, error)
	Create(ctx context.Context, address *domain.Address) error
	Update(ctx context.Context, address *domain.Address) error
	Delete(ctx context.Context, id int) error
}

type AddressRepositoryImpl struct {
	DB *gorm.DB
}

// Create implements AddressRepository.
func (a *AddressRepositoryImpl) Create(ctx context.Context, address *domain.Address) error {
	return a.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(address).Error; err != nil {
			return err
		}
		return nil
	})
}

// Delete implements AddressRepository.
func (a *AddressRepositoryImpl) Delete(ctx context.Context, id int) error {
	return a.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&domain.Address{}).Error; err != nil {
			return err
		}
		return nil
	})
}

// FindAll implements AddressRepository.
func (a *AddressRepositoryImpl) FindAll(ctx context.Context, userId int) ([]*domain.Address, error) {
	addresses := []*domain.Address{}
	if err := a.DB.WithContext(ctx).Where("id_user = ?", userId).Find(&addresses).Error; err != nil {
		return nil, err
	}
	return addresses, nil
}

// FindById implements AddressRepository.
func (a *AddressRepositoryImpl) FindById(ctx context.Context, id int) (*domain.Address, error) {
	address := &domain.Address{}
	if err := a.DB.WithContext(ctx).Where("id = ?", id).First(address).Error; err != nil {
		return nil, err
	}
	return address, nil
}

// Update implements AddressRepository.
func (a *AddressRepositoryImpl) Update(ctx context.Context, address *domain.Address) error {
	return a.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existing domain.Address

		if err := tx.Where("id = ? AND id_user = ?", address.ID, address.IDUser).
			First(&existing).Error; err != nil {
			return err
		}

		if err := tx.Model(&existing).Updates(address).Error; err != nil {
			return err
		}
		return nil
	})
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &AddressRepositoryImpl{
		DB: db,
	}
}
