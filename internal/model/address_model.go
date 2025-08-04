package model

import "time"

type AddressResponse struct {
	ID           uint       `json:"id"`
	JudulAlamat  string     `json:"judul_alamat"`
	NamaPenerima string     `json:"nama_penerima"`
	NoTelp       string     `json:"no_telp"`
	DetailAlamat string     `json:"detail_alamat"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

type CreateAddressRequest struct {
	IDUser       uint   `json:"id_user" validate:"required"`
	JudulAlamat  string `json:"judul_alamat" validate:"required,min=3,max=255"`
	NamaPenerima string `json:"nama_penerima" validate:"required,min=3,max=255"`
	NoTelp       string `json:"no_telp" validate:"required,numeric,min=10,max=15"`
	DetailAlamat string `json:"detail_alamat" validate:"required,min=10,max=500"`
}

type UpdateAddressRequest struct {
	ID           uint   `json:"id" validate:"required"`
	IDUser       uint   `json:"id_user" validate:"required"`
	JudulAlamat  string `json:"judul_alamat" validate:"required,min=3,max=255"`
	NamaPenerima string `json:"nama_penerima" validate:"required,min=3,max=255"`
	NoTelp       string `json:"no_telp" validate:"required,numeric,min=10,max=15"`
	DetailAlamat string `json:"detail_alamat" validate:"required,min=10,max=500"`
}
