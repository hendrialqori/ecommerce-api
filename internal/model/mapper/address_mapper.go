package mapper

import (
	"internship-mini-project/internal/domain"
	"internship-mini-project/internal/model"
)

func ToAddressResponse(address *domain.Address) *model.AddressResponse {
	if address == nil {
		return nil
	}

	return &model.AddressResponse{
		ID:           address.ID,
		JudulAlamat:  address.JudulAlamat,
		NamaPenerima: address.NamaPenerima,
		NoTelp:       address.NoTelp,
		DetailAlamat: address.DetailAlamat,
		CreatedAt:    &address.CreatedAt,
		UpdatedAt:    &address.UpdatedAt,
	}
}
