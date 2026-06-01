package internal

import (
	"encoding/json"
	"fmt"
)

var weatherBaseUrl = "https://api.open-meteo.com/v1/forecast?"

type WeatherResponse struct {
	Current CurrentReponse `json:"current"`
	City    string
}
type CurrentReponse struct {
	Time        string  `json:"time"`
	Temperature float64 `json:"temperature_2m"`
	WindSpeed   float64 `json:"wind_speed_10m"`
	Humidity    int     `json:"relative_humidity_2m"`
}

func GetWeatherData(city string) (WeatherResponse, error) {
	geoCoding, err := GetCityData(city)
	if err != nil {
		return WeatherResponse{}, err
	}

	url := fmt.Sprintf(
		"%slatitude=%f&longitude=%f&current=temperature_2m,relative_humidity_2m,wind_speed_10m",
		weatherBaseUrl,
		geoCoding.Results[0].Latitude,
		geoCoding.Results[0].Longitude,
	)

	resp, err := HttpClient.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}

	defer resp.Body.Close()

	var weatherResponse WeatherResponse

	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return WeatherResponse{}, err
	}

	return WeatherResponse{
		Current: weatherResponse.Current,
		City:    geoCoding.Results[0].Name,
	}, nil

}
