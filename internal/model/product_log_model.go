package model

import "time"

type ProductLogResponse struct {
	ID            uint      `json:"id"`
	NamaProduk    string    `json:"nama_produk"`
	Slug          string    `json:"slug"`
	HargaReseller uint      `json:"harga_reseller"`
	HargaKonsumen uint      `json:"harga_konsumen"`
	Deskripsi     string    `json:"deskripsi"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	IDProduk   uint `json:"produk"`
	IDToko     uint `json:"id_toko"`
	IDCategory uint `json:"id_category"`

	// Category      *CategoryResponse       `json:"category,omitempty"`
	// Toko          *TokoResponse           `json:"toko,omitempty"`
	// ProductPhoto  *[]ProductPhotoResponse `json:"foto,omitempty"`
}

type CreateProductLogRequest struct {
	IDProduk      uint   `json:"id_produk" validate:"required"`
	IDToko        uint   `json:"id_toko" validate:"required"`
	IDCategory    uint   `json:"id_category" validate:"required"`
	NamaProduk    string `json:"nama_produk" validate:"required,min=3,max=255"`
	Slug          string `json:"slug" validate:"required,min=3,max=255"`
	HargaReseller uint   `json:"harga_reseller" validate:"required,numeric"`
	HargaKonsumen uint   `json:"harga_konsumen" validate:"required,numeric"`
	Deskripsi     string `json:"deskripsi" validate:"required,min=3,max=1000"`
}
