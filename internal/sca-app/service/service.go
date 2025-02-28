package service

import (
	"github.com/rvxt21/sca-agency/internal/sca-app/models"
)

type Service interface {
	CreateSpyCat(spyCat *models.SpyCat) error
	DeleteSpyCat(id uint) error
	UpdateSalary(id uint, newSalary float64) error
	GetAllSpyCats() ([]models.SpyCat, error)
	GetSpyCat(id uint) (*models.SpyCat, error)
}
