package mapper

import (
	"internship-mini-project/internal/domain"
	"internship-mini-project/internal/model"
)

func ToUserResponse(user *domain.User) *model.UserResponse {
	return &model.UserResponse{
		ID:           user.ID,
		Nama:         user.Nama,
		NoTelp:       user.NoTelp,
		TanggalLahir: user.TanggalLahir,
		JenisKelamin: user.JenisKelamin,
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		IdProvinsi:   user.IdProvinsi,
		IdKota:       user.IdKota,
		IsAdmin:      user.IsAdmin,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Toko: &domain.Toko{
			ID:        user.Toko.ID,
			NamaToko:  user.Toko.NamaToko,
			UrlFoto:   user.Toko.UrlFoto,
			CreatedAt: user.Toko.CreatedAt,
			UpdatedAt: user.Toko.UpdatedAt,
		},
	}
}
