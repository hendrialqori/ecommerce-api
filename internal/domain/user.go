package domain

import (
	"time"
)

type User struct {
	ID           uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Nama         string     `gorm:"column:nama;type:varchar(255);not null" json:"nama"`
	KataSandi    string     `gorm:"column:kata_sandi;type:varchar(255);not null" json:"-"`
	NoTelp       string     `gorm:"column:notelp;type:varchar(255);unique;not null" json:"notelp"`
	TanggalLahir *time.Time `gorm:"column:tanggal_lahir;type:date" json:"tanggal_lahir"`
	JenisKelamin string     `gorm:"column:jenis_kelamin;type:varchar(255)" json:"jenis_kelamin"`
	Tentang      string     `gorm:"column:tentang;type:text" json:"tentang"`
	Pekerjaan    string     `gorm:"column:pekerjaan;type:varchar(255)" json:"pekerjaan"`
	Email        string     `gorm:"column:email;type:varchar(255);unique;not null" json:"email"`
	IdProvinsi   string     `gorm:"column:id_provinsi;type:varchar(255)" json:"id_provinsi"`
	IdKota       string     `gorm:"column:id_kota;type:varchar(255)" json:"id_kota"`
	IsAdmin      bool       `gorm:"column:isAdmin;type:boolean;default:false" json:"isAdmin"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	CreatedAt    time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}
