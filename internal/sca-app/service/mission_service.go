package service

import (
	"errors"
	"fmt"

	"github.com/rvxt21/sca-agency/internal/sca-app/models"
	"github.com/rvxt21/sca-agency/internal/sca-app/storage"
)

type MissionService struct {
	st storage.MissionStorage
}

func NewMissionService(st storage.MissionStorage) *MissionService {
	return &MissionService{
		st: st,
	}
}

func (s *MissionService) CreateMission(m *models.Mission) error {

	if uniqueTargets := m.CheckTargetsUnique(); !uniqueTargets {
		return fmt.Errorf("targets are not unique")
	}
	if len(m.Targets) > 3 {
		return errors.New("mission cannot have more than 3 targets")
	}

	if err := s.st.CreateMission(m); err != nil {
		return errors.New("error creating mission")
	}
	return nil
}
