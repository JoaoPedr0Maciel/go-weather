package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/JoaoPedr0Maciel/go-weather/internal"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("⏱️  Tempo total: %s\n", time.Since(start))
	}()

	if len(os.Args) < 2 {
		fmt.Println("Uso: weather <cidade>")
		os.Exit(1)
	}

	city := strings.Join(os.Args[1:], " ")

	weather, err := internal.GetWeatherData(city)
	if err != nil {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}

	fmt.Printf("\n📍 %s\n\n", weather.City)
	fmt.Printf("🌡️  Temperatura: %.1f°C\n", weather.Current.Temperature)
	fmt.Printf("💧 Umidade: %d%%\n", weather.Current.Humidity)
	fmt.Printf("💨 Vento: %.1f km/h\n", weather.Current.WindSpeed)
	fmt.Printf("🕒 Atualizado: %s\n", weather.Current.Time)
}
