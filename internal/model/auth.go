package model

import "internship-mini-project/internal/domain"

type Auth struct {
	ID      uint
	Email   string
	Nama    string
	NoTelp  string
	IsAdmin bool
	Toko    *domain.Toko
}
