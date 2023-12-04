package models

import "time"

type Photo struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Caption   string    `gorm:"type:varchar(255)" json:"caption"`
	PhotoURL  string    `gorm:"type:varchar(255);not null" json:"photo_url"`
	UserID    uint      `gorm:"not null;AUTO_INCREMENT" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}