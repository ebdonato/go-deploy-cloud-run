package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type openWeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

const weather_url = "https://api.openweathermap.org/data/2.5/weather?q=%s,br&units=metric&appid=%s"

type OpenWeather struct {
	apiKey string
}

func InstanceOpenWeather(apiKey string) *OpenWeather {
	return &OpenWeather{
		apiKey: apiKey,
	}
}

func (o *OpenWeather) GetTemperature(location string) (Temperature, error) {
	url := fmt.Sprintf(weather_url, location, o.apiKey)

	req, err := http.Get(url)
	if err != nil {
		return Temperature{}, fmt.Errorf("FAILED TO REQUEST WEATHER API: %v", err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return Temperature{}, fmt.Errorf("FAILED TO READ WEATHER API RESPONSE: %v", err)
	}

	var data openWeatherResponse
	err = json.Unmarshal(res, &data)
	if err != nil {
		return Temperature{}, fmt.Errorf("FAILED TO PARSE WEATHER API RESPONSE: %v", err)
	}

	if req.StatusCode != 200 {
		return Temperature{}, fmt.Errorf("FAILED TO REQUEST WEATHER API. StatusCode: %v", req.StatusCode)
	}

	return Temperature{
		Celsius:    data.Main.Temp,
		Fahrenheit: data.Main.Temp*1.8 + 32,
		Kelvin:     data.Main.Temp + +273,
	}, nil
}
