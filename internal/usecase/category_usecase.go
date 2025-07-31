package usecase

import (
	"context"
	"internship-mini-project/internal/domain"
	"internship-mini-project/internal/exception"
	"internship-mini-project/internal/model"
	"internship-mini-project/internal/model/mapper"
	"internship-mini-project/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type CategoryUseCase interface {
	FindAll(ctx context.Context) ([]*model.CategoryResponse, error)
	FindById(ctx context.Context, id int) (*model.CategoryResponse, error)
	Create(ctx context.Context, req *model.CreateCategoryRequest) (*model.CategoryResponse, error)
	Update(ctx context.Context, req *model.UpdateCategoryRequest) (*model.CategoryResponse, error)
	Delete(ctx context.Context, id int) error
}

type CategoryUseCaseImpl struct {
	CategoryRepository repository.CategoryRepository
	Logger             *logrus.Logger
	Validate           *validator.Validate
}

// Create implements CategoryUseCase.
func (c *CategoryUseCaseImpl) Create(ctx context.Context, req *model.CreateCategoryRequest) (*model.CategoryResponse, error) {
	if err := c.Validate.Struct(req); err != nil {
		c.Logger.WithError(err).Error("Validation failed for create request")
		return nil, err
	}

	category := &domain.Category{
		NamaKategori: req.NamaKategori,
	}

	if err := c.CategoryRepository.Create(ctx, category); err != nil {
		c.Logger.WithError(err).Error("Failed to create category")
		return nil, err
	}

	return mapper.ToCategoryResponse(category), nil
}

// Delete implements CategoryUseCase.
func (c *CategoryUseCaseImpl) Delete(ctx context.Context, id int) error {
	if err := c.CategoryRepository.Delete(ctx, id); err != nil {
		c.Logger.WithError(err).Error("Failed to delete category")
		return err
	}
	return nil
}

// FindAll implements CategoryUseCase.
func (c *CategoryUseCaseImpl) FindAll(ctx context.Context) ([]*model.CategoryResponse, error) {
	categories, err := c.CategoryRepository.FindAll(ctx)
	if err != nil {
		c.Logger.WithError(err).Error("Failed to find all categories")
		return nil, err
	}

	categoryResponses := make([]*model.CategoryResponse, len(categories))
	for i, category := range categories {
		categoryResponses[i] = mapper.ToCategoryResponse(category)
	}

	return categoryResponses, nil
}

// FindById implements CategoryUseCase.
func (c *CategoryUseCaseImpl) FindById(ctx context.Context, id int) (*model.CategoryResponse, error) {
	category, err := c.CategoryRepository.FindById(ctx, id)
	if err != nil {
		c.Logger.WithError(err).Error("Failed to find category by ID")
		return nil, exception.ErrDataNotFound
	}

	if category == nil {
		c.Logger.Warn("Category not found")
		return nil, exception.ErrDataNotFound
	}

	return mapper.ToCategoryResponse(category), nil
}

// Update implements CategoryUseCase.
func (c *CategoryUseCaseImpl) Update(ctx context.Context, req *model.UpdateCategoryRequest) (*model.CategoryResponse, error) {
	if err := c.Validate.Struct(req); err != nil {
		c.Logger.WithError(err).Error("Validation failed for update request")
		return nil, err
	}

	category := &domain.Category{
		ID:           req.ID,
		NamaKategori: req.NamaKategori,
	}

	if err := c.CategoryRepository.Update(ctx, category); err != nil {
		c.Logger.WithError(err).Error("Failed to update category")
		return nil, err
	}

	return mapper.ToCategoryResponse(category), nil
}

func NewCategoryUseCase(
	categoryRepository repository.CategoryRepository,
	logger *logrus.Logger,
	validate *validator.Validate,
) CategoryUseCase {
	return &CategoryUseCaseImpl{
		CategoryRepository: categoryRepository,
		Logger:             logger,
		Validate:           validate,
	}
}
