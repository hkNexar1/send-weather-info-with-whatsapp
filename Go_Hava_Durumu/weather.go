package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	
)

type Location struct {
	Name   string `json:"name"`
	Region string `json:"region"`
	TZID   string `json:"tz_id"`
}

type Condition struct {
	Text string `json:"text"`
}

type CurrentWeather struct {
	TempC   float64   `json:"temp_c"`
	WindMph float64   `json:"wind_mph"`
	WindDir string    `json:"wind_dir"`
	Humidity int       `json:"humidity"`
	UV      float64   `json:"uv"`
}

type WeatherData struct {
	Location Location       `json:"location"`
	Current  CurrentWeather `json:"current"`
	
}

func main() {
	url := "https://api.weatherapi.com/v1/current.json?key=APİ-KEY&q=London"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Hata: ", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Hata: ", err)
		return
	}

	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Println("Hata: ", err)
		return
	}

	file, err := os.Create("weather_info.txt")
	if err != nil {
		fmt.Println("Hata: ", err)
		return
	}
	defer file.Close()


	fmt.Fprintf(file, "Yer: %s\n", weatherData.Location.Name)
	fmt.Fprintf(file, "Region: %s\n", weatherData.Location.Region)
	fmt.Fprintf(file, "Bolge: %s\n", weatherData.Location.TZID)
	fmt.Fprintf(file, "Sicaklik (C): %.2f\n", weatherData.Current.TempC)
	fmt.Fprintf(file, "Ruzgar Hizi (mph): %.2f\n", weatherData.Current.WindMph)
	fmt.Fprintf(file, "Ruzgar Yonu: %s\n", weatherData.Current.WindDir)
	fmt.Fprintf(file, "Nem: %d\n", weatherData.Current.Humidity)
	fmt.Fprintf(file, "UV: %.2f\n", weatherData.Current.UV)

	fmt.Println("Txt cikti.")
}

