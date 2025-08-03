package mapper

import (
	"internship-mini-project/internal/domain"
	"internship-mini-project/internal/model"
)

func ToProductResponse(product *domain.Product) *model.ProductResponse {
	photos := make([]model.ProductPhotoResponse, 0)
	for _, photo := range *product.ProductPhoto {
		photos = append(photos, model.ProductPhotoResponse{
			Url: photo.Url,
		})
	}

	return &model.ProductResponse{
		ID:            product.ID,
		NamaProduk:    product.NamaProduk,
		Slug:          product.Slug,
		HargaReseller: product.HargaReseller,
		HargaKonsumen: product.HargaKonsumen,
		Stok:          product.Stok,
		Deskripsi:     product.Deskripsi,
		CreatedAt:     product.CreatedAt,
		UpdatedAt:     product.UpdatedAt,
		Category: &model.CategoryResponse{
			ID:           product.Category.ID,
			NamaKategori: product.Category.NamaKategori,
		},
		Toko: &model.TokoResponse{
			ID:       product.Toko.ID,
			NamaToko: product.Toko.NamaToko,
		},
		ProductPhoto: &photos,
	}
}
