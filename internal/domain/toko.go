package domain

import "time"

type Toko struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`     // id AUTO Increment
	UserID    uint      `gorm:"column:id_user;not null" json:"id_user"` // Kolom foreign key
	NamaToko  string    `gorm:"column:nama_toko;type:varchar(255);not null" json:"nama_toko"`
	UrlFoto   string    `gorm:"column:url_foto;type:varchar(255)" json:"url_foto"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"` // created_at date
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // updated_at date
	User      User      `gorm:"foreignKey:UserID;references:ID" json:"user,omitzero"`
}
