package mapper

import (
	"internship-mini-project/internal/domain"
	"internship-mini-project/internal/model"
)

func ToTrxResponse(trx *domain.Transaction) *model.TrxResponse {
	details := make([]model.TrxDetailResponse, 0)
	for _, detail := range *trx.TransactionDetail {
		details = append(details, model.TrxDetailResponse{
			ID:         detail.ID,
			Kuantitas:  detail.Kuantitas,
			HargaTotal: detail.HargaTotal,
			Product: &model.TrxProductResponse{
				ID:            detail.LogProduk.ID,
				NamaProduk:    detail.LogProduk.NamaProduk,
				Slug:          detail.LogProduk.Slug,
				HargaReseller: detail.LogProduk.HargaReseller,
				HargaKonsumen: detail.LogProduk.HargaKonsumen,
				Deskripsi:     detail.LogProduk.Deskripsi,
				Category: &model.TrxCategoryResponse{
					ID:           detail.LogProduk.Category.ID,
					NamaKategori: detail.LogProduk.Category.NamaKategori,
				},
			},
			Toko: &model.TrxTokoResonse{
				ID:       detail.Toko.ID,
				NamaToko: detail.Toko.NamaToko,
				UrlFoto:  detail.Toko.UrlFoto,
			},
		})
	}

	return &model.TrxResponse{
		ID:          trx.ID,
		IDUser:      trx.IDUser,
		MethodBayar: trx.MethodBayar,
		HargaTotal:  trx.HargaTotal,
		KodeInvoice: trx.KodeInvoice,
		CreatedAt:   &trx.CreatedAt,
		UpdatedAt:   &trx.UpdatedAt,
		AlamatPengiriman: &model.AddressResponse{
			ID:           trx.AlamatPengiriman.ID,
			JudulAlamat:  trx.AlamatPengiriman.JudulAlamat,
			NamaPenerima: trx.AlamatPengiriman.NamaPenerima,
			DetailAlamat: trx.AlamatPengiriman.DetailAlamat,
			NoTelp:       trx.AlamatPengiriman.NoTelp,
		},
		TransactionDetail: &details,
	}
}
