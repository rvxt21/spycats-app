package storage

import (
	"errors"
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

	if mission.CatID != nil {
		return fmt.Errorf("failed to delete mission, mission already assigned to a cat")
	}

	if err := s.db.Delete(&models.Mission{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete mission: %w", err)
	}
	return nil
}

func (s *MissionStorage) UpdateMissionStatus(id uint, isCompleted bool) error {
	var mission models.Mission
	err := s.db.Where("id = ?", id).First(&mission).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("mission with ID %d not found", id)
		}
		return fmt.Errorf("failed to fetch mission: %w", err)
	}
	if mission.IsCompleted {
		return fmt.Errorf("cannot update status for completed mission")
	}

	mission.IsCompleted = isCompleted

	if err := s.db.Save(&mission).Error; err != nil {
		return fmt.Errorf("failed to update mission status: %w", err)
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

func (s *MissionStorage) SetCat(missionId, catId uint) error {
	var mission models.Mission
	if err := s.db.First(&mission, missionId).Error; err != nil {
		return fmt.Errorf("mission with ID %d not found: %v", missionId, err)
	}

	var existingMission models.Mission
	if err := s.db.Where("cat_id = ?", catId).First(&existingMission).Error; err == nil {
		return fmt.Errorf("cat with ID %d is already assigned to a mission", catId)
	}

	var cat models.SpyCat
	if err := s.db.First(&cat, catId).Error; err != nil {
		return fmt.Errorf("cat with ID %d not found: %v", catId, err)
	}

	mission.CatID = &catId

	if err := s.db.Save(&mission).Error; err != nil {
		return fmt.Errorf("failed to update mission with ID %d: %v", missionId, err)
	}

	return nil
}
