package repository

import (
	"context"
	"internship-mini-project/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	CountByEmail(ctx context.Context, email string) (int64, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

// CountByEmail implements UserRepository.
func (u *UserRepositoryImpl) CountByEmail(ctx context.Context, email string) (int64, error) {
	var countUser int64
	if err := u.DB.WithContext(ctx).Model(&domain.User{}).Where("email = ?", email).Count(&countUser).Error; err != nil {
		return 0, err
	}
	return countUser, nil
}

// FindByEmail implements UserRepository.
func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	if err := u.DB.WithContext(ctx).Preload("Toko").Where("email = ?", email).First(user).Error; err != nil {
		return nil, err // Other error
	}

	return user, nil
}

// Create implements UserRepository.
func (u *UserRepositoryImpl) Create(ctx context.Context, user *domain.User) error {
	return u.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(ctx context.Context, user *domain.User) error {
	return u.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(user).Error; err != nil {
			return err
		}
		return nil
	})
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}
