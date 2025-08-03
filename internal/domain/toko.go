package domain

import "time"

type Toko struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IDUser    uint      `gorm:"column:id_user;not null;unique" json:"id_user,omitempty"`
	NamaToko  string    `gorm:"column:nama_toko;type:varchar(255);not null" json:"nama_toko"`
	UrlFoto   string    `gorm:"column:url_foto;type:varchar(255)" json:"url_foto"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	User      *User     `gorm:"foreignKey:IDUser;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
}
