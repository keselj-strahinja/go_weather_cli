package config

type ConfigLoader interface {
	LoadConfig() (interface{}, error)
}

type GeoConfigLoader struct {
	APIKey
}

type WeatherConfigLoader struct {
	APIKey
}

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

func (g *GeoConfigLoader) LoadConfig() (interface{}, error) {

	config := &GeoConfig{
		BaseURL: BaseURL{
			URL: geoBaseURL,
		},
		APIKey: APIKey{
			Value: g.APIKey.Value,
		},
	}

	return config, nil
}

func (w *WeatherConfigLoader) LoadConfig() (interface{}, error) {

	config := &WeatherConfig{
		BaseURL: BaseURL{
			URL: weatherBaseURL,
		},
		APIKey: APIKey{
			Value: w.APIKey.Value,
		},
	}

	return config, nil
}
