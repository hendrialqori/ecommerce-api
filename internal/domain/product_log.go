package domain

import "time"

type ProductLog struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaProduk    string    `gorm:"column:nama_produk;type:varchar(255);not null" json:"nama_produk"`
	Slug          string    `gorm:"column:slug;type:varchar(255);unique;not null" json:"slug"`
	HargaReseller uint      `gorm:"column:harga_reseller;not null" json:"harga_reseller"`
	HargaKonsumen uint      `gorm:"column:harga_konsumen;not null" json:"harga_konsumen"`
	Deskripsi     string    `gorm:"column:deskripsi;type:text" json:"deskripsi"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	// Foreign keys
	IDProduk   uint `gorm:"column:id_produk;" json:"id_user"`
	IDToko     uint `gorm:"column:id_toko;" json:"id_toko"`
	IDCategory uint `gorm:"column:id_category;" json:"id_category"`

	// Relasi
	Product  *Product  `gorm:"foreignKey:IDProduk;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"produk,omitempty"`
	Toko     *Toko     `gorm:"foreignKey:IDToko;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"toko,omitempty"`
	Category *Category `gorm:"foreignKey:IDCategory;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"category,omitempty"`
}
