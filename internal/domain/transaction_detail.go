package domain

import "time"

type TransactionDetail struct {
	ID    uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	IDTrx uint         `gorm:"column:id_trx;" json:"id_trx,omitempty"`
	Trx   *Transaction `gorm:"foreignKey:IDTrx;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"trx,omitempty"`

	IDLogProduk uint        `gorm:"column:id_log_produk;" json:"id_log_produk,omitempty"`
	LogProduk   *ProductLog `gorm:"foreignKey:IDLogProduk;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"product,omitempty"`

	IDToko uint  `gorm:"column:id_toko" json:"id_toko"`
	Toko   *Toko `gorm:"foreignKey:IDToko;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"toko,omitempty"`

	Kuantitas  uint      `gorm:"column:kuantitas;not null;" json:"kuantitas"`
	HargaTotal uint      `gorm:"column:harga_total;" json:"harga_total"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}
