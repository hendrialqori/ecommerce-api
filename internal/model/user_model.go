package model

import (
	"internship-mini-project/internal/domain"
	"time"
)

type RegisterUserRequest struct {
	Nama         string `json:"nama" validate:"required,min=3,max=255"`
	Email        string `json:"email" validate:"required,email"`
	KataSandi    string `json:"kata_sandi" validate:"required,min=8"`
	NoTelp       string `json:"no_telp" validate:"required,numeric,min=10,max=15"`
	TanggalLahir string `json:"tanggal_lahir" validate:"omitempty,datetime=2006-01-02"`
	JenisKelamin string `json:"jenis_kelamin" validate:"omitempty,oneof=Laki-Laki Perempuan"`
	Tentang      string `json:"tentang" validate:"omitempty,max=500"`
	Pekerjaan    string `json:"pekerjaan" validate:"omitempty,max=255"`
	IdProvinsi   string `json:"id_provinsi" validate:"omitempty,max=255"`
	IdKota       string `json:"id_kota" validate:"omitempty,max=255"`
}

type UpdateUserRequest struct {
	Nama         string `json:"nama" validate:"omitempty,min=3,max=255"`
	Email        string `json:"email" validate:"required,email"`
	NoTelp       string `json:"no_telp" validate:"required,numeric,min=10,max=15"`
	TanggalLahir string `json:"tanggal_lahir" validate:"omitempty,datetime=2006-01-02"`
	JenisKelamin string `json:"jenis_kelamin" validate:"omitempty,oneof=Laki-Laki Perempuan"`
	Tentang      string `json:"tentang" validate:"omitempty,max=500"`
	Pekerjaan    string `json:"pekerjaan" validate:"omitempty,max=255"`
	IdProvinsi   string `json:"id_provinsi" validate:"omitempty,max=255"`
	IdKota       string `json:"id_kota" validate:"omitempty,max=255"`
}

type LoginUserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	KataSandi string `json:"kata_sandi" validate:"required,max=100"`
}

type UserResponse struct {
	ID           uint         `json:"id"`
	Nama         string       `json:"nama"`
	NoTelp       string       `json:"no_telp"`
	TanggalLahir string       `json:"tanggal_lahir"`
	JenisKelamin string       `json:"jenis_kelamin"`
	Tentang      string       `json:"tentang"`
	Pekerjaan    string       `json:"pekerjaan"`
	Email        string       `json:"email"`
	IdProvinsi   string       `json:"id_provinsi"`
	IdKota       string       `json:"id_kota"`
	IsAdmin      bool         `json:"is_admin"`
	Toko         *domain.Toko `json:"toko"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

type VerifyUserRequest struct {
	Token string `json:"token,omitempty"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
