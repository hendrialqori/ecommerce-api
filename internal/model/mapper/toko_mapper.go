package mapper

import (
	"internship-mini-project/internal/domain"
	"internship-mini-project/internal/model"
)

func ToTokoResponse(toko *domain.Toko) *model.TokoResponse {
	return &model.TokoResponse{
		ID:        toko.ID,
		IDUser:    toko.IDUser,
		NamaToko:  toko.NamaToko,
		UrlFoto:   toko.UrlFoto,
		CreatedAt: toko.CreatedAt,
		UpdatedAt: toko.UpdatedAt,
	}
}
