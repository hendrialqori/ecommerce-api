package usecase

import (
	"context"
	"internship-mini-project/internal/domain"
	"internship-mini-project/internal/model"
	"internship-mini-project/internal/model/mapper"
	"internship-mini-project/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProductPhotoUseCase interface {
	Delete(ctx context.Context, id uint) error
	Create(ctx context.Context, req *model.CreateProductPhotoRequest) (*model.ProductPhotoResponse, error)
}

type ProductPhotoUseCaseImpl struct {
	PhotoRepo repository.ProductPhotoRepository
	Logger    *logrus.Logger
	Validate  *validator.Validate
}

// Create implements ProductPhotoUseCase.
func (p *ProductPhotoUseCaseImpl) Create(ctx context.Context, req *model.CreateProductPhotoRequest) (*model.ProductPhotoResponse, error) {
	if err := p.Validate.Struct(req); err != nil {
		p.Logger.WithError(err).Error("Validation failed for create request")
		return nil, err
	}

	photo := &domain.ProductPhoto{
		IDProduk: req.IDProduct,
		Url:      req.Url,
	}

	if err := p.PhotoRepo.Create(ctx, photo); err != nil {
		p.Logger.WithError(err).Error("Failed to create product photo")
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return mapper.ToProductPhotoResponse(photo), nil
}

// Delete implements ProductPhotoUseCase.
func (p *ProductPhotoUseCaseImpl) Delete(ctx context.Context, id uint) error {
	if err := p.PhotoRepo.Delete(ctx, id); err != nil {
		p.Logger.WithError(err).Error("Failed to delete product photo")
		return err
	}
	return nil
}

func NewProductPhotoUseCase(
	photoRepo repository.ProductPhotoRepository,
	logger *logrus.Logger,
	validate *validator.Validate,
) ProductPhotoUseCase {
	return &ProductPhotoUseCaseImpl{
		PhotoRepo: photoRepo,
		Logger:    logger,
		Validate:  validate,
	}
}
