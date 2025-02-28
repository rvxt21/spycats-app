package catapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type BreedAPIChecker struct {
	cache  map[string]struct{}
	apiURL string
}

func New(apiURL string) *BreedAPIChecker {
	return &BreedAPIChecker{
		cache:  make(map[string]struct{}),
		apiURL: apiURL,
	}
}

type CatBreed struct {
	Name string `json:"name"`
}

func (c *BreedAPIChecker) GetBreeds() error {
	resp, err := http.Get(c.apiURL)
	if err != nil {
		return fmt.Errorf("failed to fetch breeds: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var breeds []CatBreed
	if err := json.Unmarshal(body, &breeds); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	for _, breed := range breeds {
		c.cache[strings.ToLower(breed.Name)] = struct{}{}
	}

	return nil
}

func (c *BreedAPIChecker) CheckIfBreedExists(breedName string) bool {
	_, exists := c.cache[strings.ToLower(breedName)]
	return exists
}
