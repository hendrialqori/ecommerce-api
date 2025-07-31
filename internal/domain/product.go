package domain

import "time"

type Produk struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaProduk    string    `gorm:"column:nama_produk;type:varchar(255);not null" json:"nama_produk"`
	Slug          string    `gorm:"column:slug;type:varchar(255);unique;not null" json:"slug"`
	HargaReseller string    `gorm:"column:harga_reseller;type:varchar(255);not null" json:"harga_reseller"`
	HargaKonsumen string    `gorm:"column:harga_konsumen;type:varchar(255);not null" json:"harga_konsumen"`
	Stok          int       `gorm:"column:stok;not null" json:"stok"`
	Deskripsi     string    `gorm:"column:deskripsi;type:text" json:"deskripsi"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	// Foreign keys
	IDToko     uint `gorm:"column:id_toko;not null" json:"id_toko"`
	IDCategory uint `gorm:"column:id_category;not null;unique" json:"id_category"`

	// Relasi
	Toko     Toko     `gorm:"foreignKey:IDToko;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"toko"`
	Category Category `gorm:"foreignKey:IDCategory;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"category"`
}
