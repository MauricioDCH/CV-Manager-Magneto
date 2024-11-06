package service

import (
	"CV_MANAGER/models"
	"errors"

	"gorm.io/gorm"
)

type CVService interface {
    UpdateCV(cvID uint, updatedCV models.CV) (models.CV, error)
    ListCVsByUser(userID uint) ([]models.CV, error)
    DeleteCV(cvID uint) error
}

type cvService struct {
	db *gorm.DB
}

func NewCVService(db *gorm.DB) CVService {
	return &cvService{db: db}
}

func (s *cvService) UpdateCV(cvID uint, updatedCV models.CV) (models.CV, error) {
	var cv models.CV
	if err := s.db.First(&cv, cvID).Error; err != nil {
		return models.CV{}, errors.New("hoja de vida no encontrada")
	}

	// Actualiza los campos de la hoja de vida
	cv.Name = updatedCV.Name
	cv.LastName = updatedCV.LastName
	cv.Email = updatedCV.Email
	cv.Phone = updatedCV.Phone
	cv.Experience = updatedCV.Experience
	cv.Skills = updatedCV.Skills
	cv.Languages = updatedCV.Languages
	cv.Education = updatedCV.Education

	if err := s.db.Save(&cv).Error; err != nil {
		return models.CV{}, errors.New("error actualizando la hoja de vida")
	}

	return cv, nil
}

func (s *cvService) ListCVsByUser(userID uint) ([]models.CV, error) {
	var cvs []models.CV
	err := s.db.Where("user_id = ?", userID).Find(&cvs).Error
	return cvs, err
}

func (s *cvService) DeleteCV(cvID uint) error {
	return s.db.Delete(&models.CV{}, cvID).Error
}
