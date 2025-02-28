package storage

import (
	"fmt"

	"github.com/rvxt21/sca-agency/internal/sca-app/models"
)

func (s *PostgreStorage) CreateSpyCat(spyCat *models.SpyCat) error {
	if err := s.DB.Create(spyCat).Error; err != nil {
		return fmt.Errorf("failed to create spyCat: %v", err)
	}
	return nil
}

func (s *PostgreStorage) DeleteSpyCat(id uint) error {
	if err := s.DB.Delete(&models.SpyCat{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete spyCat: %v", err)
	}
	return nil
}

func (s *PostgreStorage) UpdateSalary(id uint, newSalary float64) error {
	if err := s.DB.Model(&models.SpyCat{}).Where("id = ?", id).Update("salary", newSalary).Error; err != nil {
		return fmt.Errorf("failed to update salary: %v", err)
	}
	return nil
}

func (s *PostgreStorage) GetAllSpyCats() ([]models.SpyCat, error) {
	var spyCats []models.SpyCat
	if err := s.DB.Find(&spyCats).Error; err != nil {
		return nil, fmt.Errorf("failed to get all spyCats: %v", err)
	}
	return spyCats, nil
}

func (s *PostgreStorage) GetSpyCat(id uint) (*models.SpyCat, error) {
	var spyCat models.SpyCat
	if err := s.DB.First(&spyCat, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get spyCat: %v", err)
	}
	return &spyCat, nil
}
