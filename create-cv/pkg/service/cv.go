package service

import (
	"CV_MANAGER/models"
	"errors"

	"gorm.io/gorm"
)

type CVService interface {
	CreateCV(title, name, lastName, email, phone, experience, skills, languages, education string, userID uint) (models.CV, error)
}

type cvService struct {
	db *gorm.DB
}

func NewCVService(db *gorm.DB) CVService {
	return &cvService{db: db}
}

func (s *cvService) CreateCV(title, name, lastName, email, phone, experience, skills, languages, education string, userID uint) (models.CV, error) {

	if name == "" || lastName == "" || email == "" {
		return models.CV{}, errors.New("nombre, apellido y correo electrónico son obligatorios")
	}

	cv := models.CV{
		Title:      title,
		Name:       name,
		LastName:   lastName,
		Email:      email,
		Phone:      phone,
		Experience: experience,
		Skills:     skills,
		Languages:  languages,
		Education:  education,
		UserID:     userID,
	}

	if err := s.db.Create(&cv).Error; err != nil {
		return models.CV{}, err
	}

	return cv, nil
}
