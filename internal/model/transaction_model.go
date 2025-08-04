package model

import (
	"time"
)

type TrxResponse struct {
	ID                uint                 `json:"id"`
	IDUser            uint                 `json:"id_user"`
	HargaTotal        int                  `json:"harga_total"`
	KodeInvoice       string               `json:"kode_invoice"`
	MethodBayar       string               `json:"method_bayar"`
	AlamatPengiriman  *AddressResponse     `json:"alamat_kirim"`
	CreatedAt         *time.Time           `json:"created_at,omitempty"`
	UpdatedAt         *time.Time           `json:"updated_at,omitempty"`
	TransactionDetail *[]TrxDetailResponse `json:"detail_trx"`
}

type TrxDetailResponse struct {
	ID         uint                `json:"id"`
	Kuantitas  uint                `json:"kuantitas"`
	HargaTotal uint                `json:"harga_total"`
	Product    *TrxProductResponse `json:"product"`
	Toko       *TrxTokoResonse     `json:"toko"`
}

type TrxProductResponse struct {
	ID            uint       `json:"id"`
	NamaProduk    string     `json:"nama_produk"`
	Slug          string     `json:"slug"`
	HargaReseller uint       `json:"harga_reseller"`
	HargaKonsumen uint       `json:"harga_konsumen"`
	Deskripsi     string     `json:"deskripsi"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`

	// Relasi
	// Product  *domain.Product  `json:"produk,omitempty"`
	// Toko     *domain.Toko     `json:"toko,omitempty"`
	Category *TrxCategoryResponse `json:"category,omitempty"`
}

type TrxCategoryResponse struct {
	ID           uint   `json:"id"`
	NamaKategori string `json:"nama_category"`
}

type TrxTokoResonse struct {
	ID       uint   `json:"id"`
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto,omitempty"`
}

type CreateTrxRequest struct {
	IDUser        uint                   `json:"id_user" validate:"required"`
	MethodBayar   string                 `json:"method_bayar"`
	AlamatKirimID uint                   `json:"alamat_kirim"`
	DetailTrx     []DetailTrxItemRequest `json:"detail_trx"`
}

type DetailTrxItemRequest struct {
	ProductID uint `json:"product_id"`
	Kuantitas int  `json:"kuantitas"`
}
