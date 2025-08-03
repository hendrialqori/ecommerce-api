package mapper

import (
	"internship-mini-project/internal/domain"
	"internship-mini-project/internal/model"
)

func ToProductPhotoResponse(photo *domain.ProductPhoto) *model.ProductPhotoResponse {
	return &model.ProductPhotoResponse{
		ID:        photo.ID,
		IDProduct: photo.IDProduk,
		Url:       photo.Url,
		CreatedAt: &photo.CreatedAt,
		UpdatedAt: &photo.UpdatedAt,
	}
}
