package services

import (
	"goproject/internal/models"

	"github.com/jinzhu/gorm"
)

type KonsumenService struct {
	db  *gorm.DB
	
}

func NewKonsumenService(db *gorm.DB) *KonsumenService {
    return &KonsumenService{db: db}
}

func (s *KonsumenService) CreateKonsumen(konsumen *models.Konsumen) error {
	if err := s.db.Create(konsumen).Error; err != nil {
        return err
    }
    return nil
}

func (s *KonsumenService) GetKonsumenByID(id string) (*models.Konsumen, error) {
    var konsumen models.Konsumen
    if err := s.db.First(&konsumen, id).Error; err != nil {
        return nil, err
    }
    return &konsumen, nil
}

func (s *KonsumenService) GetAllKonsumens() ([]models.Konsumen, error) {
    var konsumens []models.Konsumen
    if err := s.db.Find(&konsumens).Error; err != nil {
        return nil, err
    }
    return konsumens, nil
}