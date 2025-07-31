package domain

import "time"

type Address struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IDUser       uint      `gorm:"column:id_user;not null" json:"id_user"`
	JudulAlamat  string    `gorm:"column:judul_alamat;type:varchar(255);not null" json:"judul_alamat"`
	NamaPenerima string    `gorm:"column:nama_penerima;type:varchar(255);not null" json:"nama_penerima"`
	NoTelp       string    `gorm:"column:no_telp;type:varchar(255);not null" json:"no_telp"`
	DetailAlamat string    `gorm:"column:detail_alamat;type:varchar(255);not null" json:"detail_alamat"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	User User `gorm:"foreignKey:IDUser;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
}
