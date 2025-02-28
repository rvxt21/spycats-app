package service

import (
	"github.com/rvxt21/sca-agency/internal/sca-app/models"
	"github.com/rvxt21/sca-agency/internal/sca-app/storage"
)

type TargetsService struct {
	st storage.TargetsStorage
}

func NewTargetService(st storage.TargetsStorage) *TargetsService {
	return &TargetsService{
		st: st,
	}
}

func (s *TargetsService) AddTargetToMission(id uint, target *models.Target) error {
	if err := s.st.AddTargetToMission(id, target); err != nil {
		return err
	}

	return nil
}

func (s *TargetsService) DeleteTarget(missionId, targetId uint) error {
	if err := s.st.DeleteTarget(missionId, targetId); err != nil {
		return err
	}
	return nil
}

func (s *TargetsService) UpdateNotes(missionId, targetId uint, notes string) error {
	if err := s.st.UpdateNotes(missionId, targetId, notes); err != nil {
		return err
	}
	return nil

}
