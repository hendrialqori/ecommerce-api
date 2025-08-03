package model

import "time"

type CategoryResponse struct {
	ID           uint       `json:"id"`
	NamaKategori string     `json:"nama_kategori"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

type CreateCategoryRequest struct {
	NamaKategori string `json:"nama_kategori" validate:"required,min=3,max=255"` // nama_kategori varchar(255), wajib

}

type UpdateCategoryRequest struct {
	ID           uint   `json:"id" validate:"required"`                          // ID kategori yang akan diupdate
	NamaKategori string `json:"nama_kategori" validate:"required,min=3,max=255"` // nama_kategori varchar(255), wajib
}
