package repository

import (
	"context"
	"internship-mini-project/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	FindByNoTelp(ctx context.Context, noTelp string) (*domain.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

// FindByNoTelp implements UserRepository.
func (u *UserRepositoryImpl) FindByNoTelp(ctx context.Context, noTelp string) (*domain.User, error) {
	user := &domain.User{}
	if err := u.DB.WithContext(ctx).Where("notelp = ?", noTelp).Take(user).Error; err != nil {
		return nil, err
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
