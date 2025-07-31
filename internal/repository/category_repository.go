package repository

import (
	"context"
	"internship-mini-project/internal/domain"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll(ctx context.Context) ([]*domain.Category, error)
	FindById(ctx context.Context, id int) (*domain.Category, error)
	Create(ctx context.Context, category *domain.Category) error
	Update(ctx context.Context, category *domain.Category) error
	Delete(ctx context.Context, id int) error
}

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

// Create implements CategoryRepository.
func (c *CategoryRepositoryImpl) Create(ctx context.Context, category *domain.Category) error {
	return c.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(category).Error; err != nil {
			return err
		}
		return nil
	})
}

// Delete implements CategoryRepository.
func (c *CategoryRepositoryImpl) Delete(ctx context.Context, id int) error {
	return c.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&domain.Category{}).Error; err != nil {
			return err
		}
		return nil
	})
}

// FindAll implements CategoryRepository.
func (c *CategoryRepositoryImpl) FindAll(ctx context.Context) ([]*domain.Category, error) {
	categories := make([]*domain.Category, 0)
	if err := c.DB.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// FindById implements CategoryRepository.
func (c *CategoryRepositoryImpl) FindById(ctx context.Context, id int) (*domain.Category, error) {
	category := &domain.Category{}
	if err := c.DB.WithContext(ctx).Where("id = ?", id).First(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

// Update implements CategoryRepository.
func (c *CategoryRepositoryImpl) Update(ctx context.Context, category *domain.Category) error {
	return c.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existing domain.Category

		if err := tx.Where("id = ?", category.ID).First(&existing).Error; err != nil {
			return err
		}

		if err := tx.Model(&existing).Updates(category).Error; err != nil {
			return err
		}

		return nil
	})
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{DB: db}
}
