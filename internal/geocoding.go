package internal

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

var geoCodingBaseUrl = "https://geocoding-api.open-meteo.com/v1/search?language=pt&format=json&count=1"

type Params struct {
	Name string
}

type GeoCodingResponse struct {
	Results []Results `json:"results"`
}

type Results struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func GetCityData(city string) (GeoCodingResponse, error) {
	cityName := url.QueryEscape(city)
	url := fmt.Sprintf(
		"%s&name=%s",
		geoCodingBaseUrl,
		cityName,
	)

	cache, err := LoadCache()
	if err != nil {
		return GeoCodingResponse{}, err
	}

	geo, found := cache[strings.ToLower(city)]

	hasValidCache := found && time.Since(geo.RequestedAt) <= 10*time.Minute
	if hasValidCache {
		fmt.Println("Cache hit!")
		return GeoCodingResponse{
			Results: []Results{
				{
					Name:      city,
					Latitude:  geo.Latitude,
					Longitude: geo.Longitude,
				},
			},
		}, nil
	}
	resp, err := HttpClient.Get(url)
	if err != nil {
		return GeoCodingResponse{}, err
	}

	defer resp.Body.Close()

	var geoCodingResponse GeoCodingResponse
	if err = json.NewDecoder(resp.Body).Decode(&geoCodingResponse); err != nil {
		return GeoCodingResponse{}, err
	}

	if err = validateCityData(geoCodingResponse, city); err != nil {
		return GeoCodingResponse{}, err
	}

	cache[strings.ToLower(city)] = GeoCoding{
		Latitude:    geoCodingResponse.Results[0].Latitude,
		Longitude:   geoCodingResponse.Results[0].Longitude,
		RequestedAt: time.Now(),
	}
	if err = SaveCache(cache); err != nil {
		return GeoCodingResponse{}, err
	}

	return geoCodingResponse, nil
}

func validateCityData(geoCodingResponse GeoCodingResponse, city string) error {
	if len(geoCodingResponse.Results) == 0 {
		return fmt.Errorf("nenhuma cidade encontrada com nome %q", city)
	}

	return nil
}
