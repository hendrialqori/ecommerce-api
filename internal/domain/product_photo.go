package domain

import "time"

type ProductPhoto struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IDProduk  uint      `gorm:"column:id_produk;not null" json:"id_produk"`
	Url       string    `gorm:"column:url;type:varchar(255);not null" json:"url"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	// Relasi ke produk
	Produk Produk `gorm:"foreignKey:IDProduk;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"produk"`
}
