package service

import "github.com/rvxt21/sca-agency/internal/sca-app/storage"

type Service interface {
}

type CatService struct {
	storage storage.Storage
}