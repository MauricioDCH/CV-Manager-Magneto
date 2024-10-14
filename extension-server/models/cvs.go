package models

import (
	"gorm.io/gorm"
)

type CVS struct {
	gorm.Model
	ID         uint   `gorm:"primary_key;" json:"id"`
	CreatedAt  string `gorm:"not null" json:"created_at"`
	UpdatedAt  string `gorm:"not null" json:"updated_at"`
	DeletedAt  string `gorm:"not null" json:"deleted_at"`
	Name       string `gorm:"not null" json:"name"`
	LastName   string `gorm:"not null" json:"last_name"`
	Email      string `gorm:"not null" json:"email"`
	Phone      string `gorm:"not null" json:"phone"`
	Experience string `gorm:"not null" json:"experience"`
	Skills     string `gorm:"not null" json:"skills"`
	Lenguages  string `gorm:"not null" json:"lenguages"`
	Education  string `gorm:"not null" json:"education"`
	User_id    uint   `gorm:"not null" json:"user_id"`
}
