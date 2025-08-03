package model

import (
	"time"
)

type ProductResponse struct {
	ID            uint                    `json:"id"`
	NamaProduk    string                  `json:"nama_produk"`
	Slug          string                  `json:"slug"`
	HargaReseller uint                    `json:"harga_reseller"`
	HargaKonsumen uint                    `json:"harga_konsumen"`
	Stok          uint                    `json:"stok"`
	Deskripsi     string                  `json:"deskripsi"`
	CreatedAt     time.Time               `json:"created_at"`
	UpdatedAt     time.Time               `json:"updated_at"`
	Category      *CategoryResponse       `json:"category,omitempty"`
	Toko          *TokoResponse           `json:"toko,omitempty"`
	ProductPhoto  *[]ProductPhotoResponse `json:"foto,omitempty"`
}

type ProductQueryParams struct {
	Page       int     `json:"page"`
	Limit      int     `json:"limit"`
	NamaProduk *string `json:"nama_produk,omitempty"`
	IDCategory *int    `json:"id_category,omitempty"`
	IDToko     *int    `json:"id_toko,omitempty"`
	MinHarga   *string `json:"min_harga,omitempty"`
	MaxHarga   *string `json:"max_harga,omitempty"`
}

type CreateProductRequest struct {
	IDToko        uint   `json:"id_toko" validate:"required"`
	IDCategory    uint   `json:"id_category" validate:"required"`
	NamaProduk    string `json:"nama_produk" validate:"required,min=3,max=255"`
	Slug          string `json:"slug" validate:"required,min=3,max=255"`
	HargaReseller uint   `json:"harga_reseller" validate:"required,numeric"`
	HargaKonsumen uint   `json:"harga_konsumen" validate:"required,numeric"`
	Stok          uint   `json:"stok" validate:"required,numeric"`
	Deskripsi     string `json:"deskripsi" validate:"required,min=3,max=1000"`
}

type UpdateProductRequest struct {
	ID            uint   `json:"id" validate:"required"`
	IDToko        uint   `json:"id_toko" validate:"required"`
	IDCategory    uint   `json:"id_category" validate:"required"`
	NamaProduk    string `json:"nama_produk" validate:"required,min=3,max=255"`
	Slug          string `json:"slug" validate:"required,min=3,max=255"`
	HargaReseller uint   `json:"harga_reseller" validate:"required,numeric"`
	HargaKonsumen uint   `json:"harga_konsumen" validate:"required,numeric"`
	Stok          uint   `json:"stok" validate:"required,numeric"`
	Deskripsi     string `json:"deskripsi" validate:"required,min=3,max=1000"`
}
