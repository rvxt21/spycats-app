package storage

import (
	"fmt"

	// "sync"

	_ "github.com/lib/pq"
	"github.com/rvxt21/sca-agency/internal/sca-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage interface {
}

type PostgreStorage struct {
	DB *gorm.DB
	// m     sync.Mutex
}

func New(connStr string) (*PostgreStorage, error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("getting sql.DB from GORM: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	if err := db.AutoMigrate(&models.SpyCat{}, &models.Mission{}, &models.Target{}); err != nil {
		return nil, fmt.Errorf("error migrating database: %w", err)
	}

	return &PostgreStorage{
		DB: db,
	}, nil
}

func (s *PostgreStorage) CreateSpyCat() {

}

func (s *PostgreStorage) DeleteSpyCat() {

}

func (s *PostgreStorage) UpdateSalary() {

}

func (s *PostgreStorage) GetAllSpyCats() {

}

func (s *PostgreStorage) GetSpyCat() {

}
