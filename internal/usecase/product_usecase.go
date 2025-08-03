package usecase

import (
	"context"
	"internship-mini-project/internal/domain"
	"internship-mini-project/internal/exception"
	"internship-mini-project/internal/helper"
	"internship-mini-project/internal/model"
	"internship-mini-project/internal/model/mapper"
	"internship-mini-project/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProductUseCase interface {
	FindAll(ctx context.Context, query *model.ProductQueryParams) ([]*model.ProductResponse, error)
	FindByID(ctx context.Context, id uint) (*model.ProductResponse, error)
	Create(ctx context.Context, req *model.CreateProductRequest) (*model.ProductResponse, error)
	Update(ctx context.Context, req *model.UpdateProductRequest) (*model.ProductResponse, error)
	Delete(ctx context.Context, id uint) error
}

type ProductUseCaseImpl struct {
	productRepo repository.ProductRepository
	Logger      *logrus.Logger
	Validate    *validator.Validate
}

// Create implements ProductUseCase.
func (p *ProductUseCaseImpl) Create(ctx context.Context, req *model.CreateProductRequest) (*model.ProductResponse, error) {
	if err := p.Validate.Struct(req); err != nil {
		p.Logger.WithError(err).Error("Validation failed for create request")
		return nil, err
	}

	product := &domain.Product{
		IDToko:        req.IDToko,
		IDCategory:    req.IDCategory,
		NamaProduk:    req.NamaProduk,
		Slug:          req.Slug,
		HargaReseller: req.HargaReseller,
		HargaKonsumen: req.HargaKonsumen,
		Stok:          req.Stok,
		Deskripsi:     req.Deskripsi,
	}

	if err := p.productRepo.Create(ctx, product); err != nil {
		p.Logger.WithError(err).Error("Failed to create product")
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	res := &model.ProductResponse{
		ID:            product.ID,
		NamaProduk:    product.NamaProduk,
		Slug:          product.Slug,
		HargaReseller: product.HargaReseller,
		HargaKonsumen: product.HargaKonsumen,
		Stok:          product.Stok,
		Deskripsi:     product.Deskripsi,
		CreatedAt:     product.CreatedAt,
		UpdatedAt:     product.UpdatedAt,
	}

	return res, nil
}

// Delete implements ProductUseCase.
func (p *ProductUseCaseImpl) Delete(ctx context.Context, id uint) error {
	if err := p.productRepo.Delete(ctx, id); err != nil {
		p.Logger.WithError(err).Error("Failed to delete product")
		return err
	}
	return nil
}

// FindAll implements ProductUseCase.
func (p *ProductUseCaseImpl) FindAll(ctx context.Context, query *model.ProductQueryParams) ([]*model.ProductResponse, error) {
	products, err := p.productRepo.FindAll(ctx, query)
	if err != nil {
		p.Logger.WithError(err).Error("Failed to find all products")
		return nil, err
	}

	var productResponses []*model.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, mapper.ToProductResponse(product))
	}

	return productResponses, nil
}

// FindByID implements ProductUseCase.
func (p *ProductUseCaseImpl) FindByID(ctx context.Context, id uint) (*model.ProductResponse, error) {
	product, err := p.productRepo.FindById(ctx, id)
	if err != nil {
		p.Logger.WithError(err).Error("Failed to find product by ID")
		return nil, exception.ErrDataNotFound
	}

	if product == nil {
		p.Logger.Error("Product not found")
		return nil, exception.ErrDataNotFound
	}

	return mapper.ToProductResponse(product), nil
}

// Update implements ProductUseCase.
func (p *ProductUseCaseImpl) Update(ctx context.Context, req *model.UpdateProductRequest) (*model.ProductResponse, error) {
	if err := p.Validate.Struct(req); err != nil {
		p.Logger.WithError(err).Error("Validation failed for update request")
		return nil, err
	}

	product := &domain.Product{
		ID:            req.ID,
		IDToko:        req.IDToko,
		IDCategory:    req.IDCategory,
		NamaProduk:    req.NamaProduk,
		Slug:          helper.ParseToSlug(req.NamaProduk),
		HargaReseller: req.HargaReseller,
		HargaKonsumen: req.HargaKonsumen,
		Stok:          req.Stok,
		Deskripsi:     req.Deskripsi,
	}

	if err := p.productRepo.Update(ctx, product); err != nil {
		p.Logger.WithError(err).Error("Failed to update product")
		return nil, err
	}

	res := &model.ProductResponse{
		ID:            product.ID,
		NamaProduk:    product.NamaProduk,
		Slug:          product.Slug,
		HargaReseller: product.HargaReseller,
		HargaKonsumen: product.HargaKonsumen,
		Stok:          product.Stok,
		Deskripsi:     product.Deskripsi,
		CreatedAt:     product.CreatedAt,
		UpdatedAt:     product.UpdatedAt,
	}

	return res, nil
}

func NewProductUseCase(productRepo repository.ProductRepository, logger *logrus.Logger, validate *validator.Validate) ProductUseCase {
	return &ProductUseCaseImpl{
		productRepo: productRepo,
		Logger:      logger,
		Validate:    validate,
	}
}
