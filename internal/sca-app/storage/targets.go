package storage

import (
	"errors"
	"fmt"

	"github.com/rvxt21/sca-agency/internal/sca-app/models"
	"gorm.io/gorm"
)

type TargetsStorage struct {
	db *gorm.DB
}

func NewTargetsStore(db *gorm.DB) *TargetsStorage {
	return &TargetsStorage{
		db: db,
	}
}

func (s *TargetsStorage) AddTargetToMission(id uint, target *models.Target) error {
	var mission models.Mission

	if err := s.db.First(&mission, id).Error; err != nil {
		return fmt.Errorf("mission with ID %d not found: %w", id, err)
	}

	if mission.IsCompleted {
		return fmt.Errorf("failed to add target, mission is already completed")
	}

	if len(mission.Targets) == 3 {
		return fmt.Errorf("failed to add target, mission can have only 3 targets")
	}

	if !mission.CheckTargetsUnique() {
		return fmt.Errorf("failed to add target, mission already have this target")
	}

	if err := s.db.Create(&target).Error; err != nil {
		return fmt.Errorf("failed to add target to mission: %w", err)
	}

	return nil
}

func (s *TargetsStorage) DeleteTarget(missionId, targetId uint) error {
	var mission models.Mission

	if err := s.db.First(&mission, missionId).Error; err != nil {
		return fmt.Errorf("mission with ID %d not found: %w", missionId, err)
	}

	if mission.IsCompleted {
		return fmt.Errorf("failed to delete target, mission is complited")
	}

	var target models.Target
	if err := s.db.Where("id = ? AND mission_id = ?", targetId, missionId).First(&target).Error; err != nil {
		return fmt.Errorf("target with ID %d not found in mission %d: %w", targetId, missionId, err)
	}

	if err := s.db.Delete(&target).Error; err != nil {
		return fmt.Errorf("failed to delete target: %w", err)
	}
	return nil
}

func (s *TargetsStorage) UpdateNotes(missionId, targetId uint, notes string) error {
	var target models.Target

	err := s.db.Where("id = ? AND mission_id = ?", targetId, missionId).First(&target).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("target with ID %d not found in mission %d", targetId, missionId)
		}
		return fmt.Errorf("failed to fetch target: %w", err)
	}

	if target.IsCompleted {
		return fmt.Errorf("cannot update notes for completed target")
	}

	target.Notes = notes

	if err := s.db.Save(&target).Error; err != nil {
		return fmt.Errorf("failed to update notes: %w", err)
	}

	return nil
}

func (s *TargetsStorage) UpdateTargerStatus(missionId, targetId uint, is_completed bool) error {
	var target models.Target

	err := s.db.Where("id = ? AND mission_id = ?", targetId, missionId).First(&target).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("target with ID %d not found in mission %d", targetId, missionId)
		}
		return fmt.Errorf("failed to fetch target: %w", err)
	}

	target.IsCompleted = is_completed

	if err := s.db.Save(&target).Error; err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}

	return nil

}
