package storage

import (
	"fmt"

	"github.com/rvxt21/sca-agency/internal/sca-app/models"
	"gorm.io/gorm"
)

type MissionStorage struct {
	db *gorm.DB
}

func NewMissionStorage(db *gorm.DB) *MissionStorage {
	return &MissionStorage{db: db}
}

func (s *MissionStorage) CreateMission(mission *models.Mission) error {

	if err := s.db.Create(mission).Error; err != nil {
		return fmt.Errorf("failed to create mission: %w", err)
	}

	return nil
}

func (s *MissionStorage) DeleteMission(id uint) error {
	var mission models.Mission

	if err := s.db.First(&mission, id).Error; err != nil {
		return fmt.Errorf("mission with ID %d not found: %w", id, err)
	}

	if mission.CatID != 0 {
		return fmt.Errorf("failed to delete mission, mission already assigned to a cat")
	}

	if err := s.db.Delete(&models.Mission{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete mission: %w", err)
	}
	return nil
}

func (s *MissionStorage) UpdateMissionStatus(id uint, isComplited bool) error {
	if err := s.db.Model(&models.Mission{}).Where("id = ?", id).Update("is_complited", isComplited).Error; err != nil {
		return fmt.Errorf("failed to update mission status: %v", err)
	}
	return nil
}

func (s *MissionStorage) GetAllMissions() ([]models.Mission, error) {
	var missions []models.Mission
	if err := s.db.Preload("Targets").Find(&missions).Error; err != nil {
		return nil, fmt.Errorf("failed to get all missions: %v", err)
	}
	return missions, nil
}

func (s *MissionStorage) GetMission(id uint) (*models.Mission, error) {
	var mission models.Mission
	if err := s.db.Preload("Targets").First(&mission, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get mission: %v", err)
	}
	return &mission, nil
}
