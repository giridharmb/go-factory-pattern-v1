package factory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	endpoint                = "https://api.openweathermap.org/data/2.5"
	pathFormatWeatherByCity = "/weather?q=%s&appid=%s&units=metric"
)

// -----------------------------------------------------

type WeatherQuery struct {
	ApiKey string
	City   string
}

type Weather struct {
	Temp     float32
	Pressure float32
	MinTemp  float32
	MaxTemp  float32
}

type WeatherDataHTTP struct {
	APIKey string `json:"apikey"`
	City   string `json:"city"`
}

type WeatherResponse struct {
	Message string
	Main    struct {
		Temp     float32 `json:"temp"`
		Pressure float32 `json:"pressure"`
		TempMin  float32 `json:"temp_min"`
		TempMax  float32 `json:"temp_max"`
	}
}

// -----------------------------------------------------

func (r WeatherResponse) ToWeather() Weather {
	return Weather{
		Temp:     r.Main.Temp,
		Pressure: r.Main.Pressure,
		MinTemp:  r.Main.TempMin,
		MaxTemp:  r.Main.TempMax,
	}
}

// -----------------------------------------------------

type IWeather interface {
	GetWeather() (Weather, error)
}

// -----------------------------------------------------

func (wq WeatherQuery) GetWeather() (Weather, error) {
	path := fmt.Sprintf(pathFormatWeatherByCity, wq.City, wq.ApiKey)
	log.Printf("@ provider : path : %v", path)

	completeURL := endpoint + path
	log.Printf("@ provider : completeURL : %v", completeURL)

	res, err := http.Get(completeURL)
	if err != nil {
		return Weather{}, fmt.Errorf("GetWeather() : failed http GET: %s", err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	// read the response body and encode it into the respose struct
	bodyRaw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Weather{}, fmt.Errorf("GetWeather() : failed reading body: %s", err)
	}

	var weatherResponse WeatherResponse
	if err = json.Unmarshal(bodyRaw, &weatherResponse); err != nil {
		return Weather{}, fmt.Errorf("GetWeather() : failed encoding body: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		return Weather{}, fmt.Errorf("GetWeather() : got error from OpenWeather: %s", weatherResponse.Message)
	}

	// return the external response converted into an entity
	return weatherResponse.ToWeather(), nil
}

// -----------------------------------------------------

func NewWeatherFactory(apiKey string, city string) interface{} {
	weather := WeatherQuery{
		ApiKey: apiKey,
		City:   city,
	}
	return weather
}

// -----------------------------------------------------

// Generic Factory

type GFactory func(string, string) interface{}

func GAbstractFactory(fact string) GFactory {
	switch fact {
	case "weather":
		return NewWeatherFactory
	default:
		return nil
	}
}

// -----------------------------------------------------
