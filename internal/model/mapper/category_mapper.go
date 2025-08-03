package mapper

import (
	"internship-mini-project/internal/domain"
	"internship-mini-project/internal/model"
)

func ToCategoryResponse(category *domain.Category) *model.CategoryResponse {
	return &model.CategoryResponse{
		ID:           category.ID,
		NamaKategori: category.NamaKategori,
		CreatedAt:    &category.CreatedAt,
		UpdatedAt:    &category.UpdatedAt,
	}
}
