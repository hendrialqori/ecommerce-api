package domain

import (
	"time"
)

type Transaction struct {
	ID     uint  `gorm:"primaryKey;autoIncrement" json:"id"`
	IDUser uint  `gorm:"column:id_user;" json:"id_user"`
	User   *User `gorm:"foreignKey:IDUser;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user,omitempty"`

	AlamatPengirimanID uint     `gorm:"column:alamat_pengiriman;,unique" json:"alamat_pengiriman"`
	AlamatPengiriman   *Address `gorm:"foreignKey:AlamatPengirimanID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"alamat_pengiriman_detail,omitempty"`

	HargaTotal  int       `gorm:"column:harga_total;" json:"harga_total"`
	KodeInvoice string    `gorm:"colum:kode_invoice;type:varchar(255);" json:"kode_invoice"`
	MethodBayar string    `gorm:"colum:kode_invoice;type:varchar(100);" json:"method_bayar"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	TransactionDetail *[]TransactionDetail `gorm:"foreignKey:IDTrx" json:"detail_trx,omitempty"`
}
