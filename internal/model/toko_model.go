package model

import "time"

type TokoResponse struct {
	ID        uint          `json:"id"`
	IDUser    uint          `json:"id_user"`
	NamaToko  string        `json:"nama_toko"`
	UrlFoto   string        `json:"url_foto"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	User      *UserResponse `json:"user,omitempty"`
}

type CreateTokoRequest struct {
	IDUser   uint   `json:"id_user" validate:"required"`
	NamaToko string `json:"nama_toko" validate:"required,min=3,max=255"` // nama_toko varchar(255), wajib
	UrlFoto  string `json:"url_foto" validate:"omitempty,url,max=255"`   // url_foto varchar(255), opsional, harus format URL
}

type UpdateTokoRequest struct {
	ID       uint   `json:"id" validate:"required"`
	IDUser   uint   `json:"id_user" validate:"required"`
	NamaToko string `json:"nama_toko" validate:"required,min=3,max=255"`
	UrlFoto  string `json:"url_foto" validate:"omitempty,url,max=255"` // url_foto varchar(255), opsional, harus format URL
}
