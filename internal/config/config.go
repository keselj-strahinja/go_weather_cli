package config

import (
	"errors"
	"os"
)

type ConfigLoader interface {
	LoadConfig() (interface{}, error)
}

type GeoConfigLoader struct{}

type WeatherConfigLoader struct{}

type APIKey struct {
	Value string
}

type BaseURL struct {
	URL string
}

type GeoConfig struct {
	BaseURL
	APIKey
}

type WeatherConfig struct {
	BaseURL
	APIKey
}

const (
	geoBaseURL     = "http://api.openweathermap.org/geo/1.0/direct"
	weatherBaseURL = "https://api.openweathermap.org/data/2.5/weather"
)

func getApiKey() (string, error) {
	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		return "", errors.New("API key is empty")
	}

	return apiKey, nil

}

func (g *GeoConfigLoader) LoadConfig() (interface{}, error) {
	apiKey, err := getApiKey()

	if err != nil {
		panic(err)
	}

	config := &GeoConfig{
		BaseURL: BaseURL{
			URL: geoBaseURL,
		},
		APIKey: APIKey{
			Value: apiKey,
		},
	}

	return config, nil
}

func (w *WeatherConfigLoader) LoadConfig() (interface{}, error) {
	apiKey, err := getApiKey()

	if err != nil {
		panic(err)
	}

	config := &WeatherConfig{
		BaseURL: BaseURL{
			URL: weatherBaseURL,
		},
		APIKey: APIKey{
			Value: apiKey,
		},
	}

	return config, nil
}
