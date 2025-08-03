package repository

import (
	"context"
	"fmt"
	"internship-mini-project/internal/domain"
	"internship-mini-project/internal/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(ctx context.Context, QueryParams *model.ProductQueryParams) ([]*domain.Product, error)
	FindById(ctx context.Context, id uint) (*domain.Product, error)
	Create(ctx context.Context, product *domain.Product) error
	Update(ctx context.Context, product *domain.Product) error
	Delete(ctx context.Context, id uint) error
}

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

// Create implements ProductRepository.
func (p *ProductRepositoryImpl) Create(ctx context.Context, product *domain.Product) error {
	return p.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(product).Error; err != nil {
			return err
		}
		return nil
	})
}

// Delete implements ProductRepository.
func (p *ProductRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return p.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&domain.Product{}).Error; err != nil {
			return err
		}
		return nil
	})
}

// FindAll implements ProductRepository.
func (p *ProductRepositoryImpl) FindAll(ctx context.Context, QueryParams *model.ProductQueryParams) ([]*domain.Product, error) {
	products := []*domain.Product{}

	query := p.DB.WithContext(ctx).Model(&domain.Product{}).Preload("Category").Preload("Toko").Preload("ProductPhoto")

	if QueryParams.NamaProduk != nil {
		query = query.Where("nama_produk LIKE ?", "%"+*QueryParams.NamaProduk+"%")
	}

	if QueryParams.IDCategory != nil {
		query = query.Where("id_category = ?", *QueryParams.IDCategory)
	}

	if QueryParams.IDToko != nil {
		query = query.Where("id_toko = ?", *QueryParams.IDToko)
	}

	if QueryParams.MinHarga != nil {
		query = query.Where("harga_konsumen >= ?", *QueryParams.MinHarga)
	}

	if QueryParams.MaxHarga != nil {
		query = query.Where("harga_konsumen <= ?", *QueryParams.MaxHarga)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

// FindById implements ProductRepository.
func (p *ProductRepositoryImpl) FindById(ctx context.Context, id uint) (*domain.Product, error) {
	product := &domain.Product{}
	if err := p.DB.WithContext(ctx).Preload("Category").Preload("Toko").Preload("ProductPhoto").Where("id = ?", id).First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// Update implements ProductRepository.
func (p *ProductRepositoryImpl) Update(ctx context.Context, product *domain.Product) error {
	return p.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existing domain.Product

		if err := tx.Where("id = ?", product.ID).First(&existing).Error; err != nil {
			fmt.Println("Error finding existing product:", err)
			return err
		}

		if err := tx.Model(&existing).Updates(product).Error; err != nil {
			fmt.Println("Error finding existing product:", err)
			return err
		}

		return nil
	})
}

func NewProductRepository(DB *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{DB: DB}
}
