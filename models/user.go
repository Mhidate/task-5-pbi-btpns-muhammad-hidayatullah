package models 

import "time"

type User struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Username  string    `gorm:"type:varchar(100);not null" json:"username"`
	Email     string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null;min:6" json:"-"`
	Photos    []Photo   `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}