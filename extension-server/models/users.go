package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint   `gorm:"primary_key;" json:"id"`
	Nombre     string `gorm:"not null" json:"name"`
	Correo     string `gorm:"unique;not null" json:"email"`
	Contraseña string `gorm:"not null" json:"password"`
}
