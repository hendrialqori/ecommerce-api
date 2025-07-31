package domain

import "time"

type Category struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaKategori string    `gorm:"column:nama_category;type:varchar(255);not null" json:"nama_category"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}
