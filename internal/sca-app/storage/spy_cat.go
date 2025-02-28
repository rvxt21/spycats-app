package storage

import (
	"fmt"

	"github.com/rvxt21/sca-agency/internal/sca-app/models"
	"gorm.io/gorm"
)

type CatStorage struct {
	db *gorm.DB
	// m     sync.Mutex
}

func New(db *gorm.DB) (Storage, error) {

	if db == nil {
		return nil, fmt.Errorf("provided gorm.db is nil")
	}
	sqldb, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("getting sql.db from GORM: %w", err)
	}

	if err := sqldb.Ping(); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	if err := db.AutoMigrate(&models.SpyCat{}, &models.Mission{}, &models.Target{}); err != nil {
		return nil, fmt.Errorf("error migrating database: %w", err)
	}

	return &CatStorage{
		db: db,
	}, nil
}

func (s *CatStorage) CreateSpyCat(spyCat *models.SpyCat) error {
	if err := s.db.Create(spyCat).Error; err != nil {
		return fmt.Errorf("failed to create spyCat: %v", err)
	}
	return nil
}

func (s *CatStorage) DeleteSpyCat(id uint) error {
	if err := s.db.Delete(&models.SpyCat{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete spyCat: %v", err)
	}
	return nil
}

func (s *CatStorage) UpdateSalary(id uint, newSalary float64) error {
	if err := s.db.Model(&models.SpyCat{}).Where("id = ?", id).Update("salary", newSalary).Error; err != nil {
		return fmt.Errorf("failed to update salary: %v", err)
	}
	return nil
}

func (s *CatStorage) GetAllSpyCats() ([]models.SpyCat, error) {
	var spyCats []models.SpyCat
	if err := s.db.Find(&spyCats).Error; err != nil {
		return nil, fmt.Errorf("failed to get all spyCats: %v", err)
	}
	return spyCats, nil
}

func (s *CatStorage) GetSpyCat(id uint) (*models.SpyCat, error) {
	var spyCat models.SpyCat
	if err := s.db.First(&spyCat, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get spyCat: %v", err)
	}
	return &spyCat, nil
}
