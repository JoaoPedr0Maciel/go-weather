# Go Weather

A simple command-line weather application built with Go.

The application receives a city name, retrieves its geographic coordinates using the Open-Meteo Geocoding API, and then fetches the current weather conditions for that location.

## Features

- Current temperature
- Humidity
- Wind speed
- City geolocation lookup
- Local file caching to reduce API requests
- Simple and fast CLI experience

## Technologies

- Go
- Open-Meteo Geocoding API
- Open-Meteo Weather API

## Installation

```bash
git clone https://github.com/JoaoPedr0Maciel/go-weather.git

cd go-weather

go run cmd/main.go "New York"
```

## Example

```bash
go run cmd/main.go "London"
```

Output:

```text
📍 London

🌡️ Temperature: 18.2°C
💧 Humidity: 71%
💨 Wind Speed: 12.5 km/h
🕒 Updated: 2026-05-31T18:00
```

## Project Structure

```text
go-weather/
├── cmd/
│   └── main.go
├── internal/
│   ├── cache.go
│   ├── geocoding.go
│   ├── weather.go
│   ├── http.go
│   └── geocoding_test.go
├── cache.json
└── go.mod
```

## Cache

The application stores geolocation data in a local `cache.json` file.

Cached entries remain valid for 10 minutes, reducing unnecessary requests to the geocoding service.

## License

This project is available for learning, experimentation, and personal use.
