package internal

import (
	"encoding/json"
	"os"
	"time"
)

type Cache map[string]GeoCoding

type GeoCoding struct {
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	RequestedAt time.Time `json:"requestedAt"`
}

func LoadCache() (Cache, error) {
	data, err := os.ReadFile("cache.json")

	if os.IsNotExist(err) {
		return make(Cache), nil
	}

	if err != nil {
		return nil, err
	}

	var cache Cache

	if err := json.Unmarshal(data, &cache); err != nil {
		return nil, err
	}

	return cache, nil
}

func SaveCache(cache Cache) error {
	data, err := json.MarshalIndent(cache, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("cache.json", data, 0644)
}
