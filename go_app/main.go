package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	getWeather("https://api.openweathermap.org/data/2.5/weather?lat=45.4211435&lon=-75.6900574&appid=fe1d80e1e103cff8c6afd190cad23fa5")
}

func getWeather(URLAddress string) {
	response, err := http.Get(URLAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var data WeatherData

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	weather := int(data.Main.Temp - 273.15)
	fmt.Printf("Temperature: %d°C\n", weather)

	fmt.Println("\nDone.")
}