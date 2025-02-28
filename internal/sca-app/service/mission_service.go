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

func (s *MissionService) DeleteMission(id uint) error {
	if err := s.st.DeleteMission(id); err != nil {
		return err
	}
	return nil
}

func (s *MissionService) UpdateMissionStatus(id uint, isCompleted bool) error {
	if err := s.st.UpdateMissionStatus(id, isCompleted); err != nil {
		return err
	}
	return nil
}

func (s *MissionService) GetAllMissions() ([]models.Mission, error) {
	missions, err := s.st.GetAllMissions()
	if err != nil {
		return nil, err
	}
	return missions, nil

}

func (s *MissionService) GetMission(id uint) (*models.Mission, error) {
	mission, err := s.st.GetMission(id)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get mission with ID %d: %w", id, err)
	}
	return mission, nil
}

func (s *MissionService) AssignCatToMission(missionId, catId uint) error {
	if err := s.st.SetCat(missionId, catId); err != nil {
		return err
	}

	return nil
}


