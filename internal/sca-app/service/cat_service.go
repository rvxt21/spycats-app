package service

import (
	"fmt"

	"github.com/rvxt21/sca-agency/internal/sca-app/models"
)

func (s *CatService) CreateSpyCat(spyCat *models.SpyCat) error {
	err := s.storage.CreateSpyCat(spyCat)
	if err != nil {

		return fmt.Errorf("failed to create spy cat: %w", err)
	}
	return nil
}

func (s *CatService) DeleteSpyCat(id uint) error {
	err := s.storage.DeleteSpyCat(id)
	if err != nil {
		return fmt.Errorf("failed to delete spy cat: %w", err)
	}
	return nil
}

func (s *CatService) UpdateSalary(id uint, newSalary float64) error {
	err := s.storage.UpdateSalary(id, newSalary)
	if err != nil {
		return fmt.Errorf("failed to update salary for spy cat with ID %d: %w", id, err)
	}
	return nil
}

func (s *CatService) GetAllSpyCats() ([]models.SpyCat, error) {
	spyCats, err := s.storage.GetAllSpyCats()
	if err != nil {
		return nil, fmt.Errorf("failed to get all spy cats: %w", err)
	}
	return spyCats, nil
}

func (s *CatService) GetSpyCat(id uint) (*models.SpyCat, error) {

	spyCat, err := s.storage.GetSpyCat(id)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get spy cat with ID %d: %w", id, err)
	}
	return spyCat, nil
}
