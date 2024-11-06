package models

import "gorm.io/gorm"

type CV struct {
	gorm.Model
	ID         uint   `gorm:"primary_key;autoIncrement" json:"id"`
	Title      string `json:"title"`
	Name       string `gorm:"not null" json:"name"`
	LastName   string `gorm:"not null" json:"last_name"`
	Email      string `gorm:"not null" json:"email"`
	Phone      string `json:"phone"`
	Experience string `json:"experience"`
	Skills     string `json:"skills"`
	Languages  string `json:"languages"`
	Education  string `json:"education"`
	UserID     uint   `gorm:"not null" json:"user_id"` // Relación con el usuario
}
