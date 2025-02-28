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
