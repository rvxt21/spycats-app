package storage

import (

	// "sync"

	_ "github.com/lib/pq"
	"github.com/rvxt21/sca-agency/internal/sca-app/models"
)

type Storage interface {
	CreateSpyCat(spyCat *models.SpyCat) error
	DeleteSpyCat(id uint) error
	UpdateSalary(id uint, newSalary float64) error
	GetAllSpyCats() ([]models.SpyCat, error)
	GetSpyCat(id uint) (*models.SpyCat, error)
}
