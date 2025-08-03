package model

import "time"

type ProductPhotoResponse struct {
	ID        uint       `json:"id,omitempty"`
	IDProduct uint       `json:"id_product,omitempty"`
	Url       string     `json:"url"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type CreateProductPhotoRequest struct {
	IDProduct uint   `json:"id_product" validate:"required"`
	Url       string `json:"url" validate:"required,url"`
}
