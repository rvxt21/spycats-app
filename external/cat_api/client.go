package catapi

import (
	"sync"
	"time"
)

type CatAPIClient struct {
	cache      map[string]struct{}
	lastUpdate time.Time
	mu         sync.RWMutex
	apiURL     string
}

