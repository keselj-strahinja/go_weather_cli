package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"weather/internal/config"
	"weather/internal/types"
	"weather/internal/util"
)

func fetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("unexpected status code: %d; and failed to read body: %v", resp.StatusCode, err)
		}
		return nil, fmt.Errorf("unexpected status code: %d; body: %s", resp.StatusCode, string(bodyBytes))
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func FetchWeather(loc *types.Coord, loader config.ConfigLoader) (data *types.WeatherData, err error) {

	cfg, err := loader.LoadConfig()

	if err != nil {
		return nil, err
	}

	weatherConfig, ok := cfg.(*config.WeatherConfig)

	if !ok {
		return nil, errors.New("unexpected config type")
	}

	baseURL := weatherConfig.BaseURL.URL
	apiKey := weatherConfig.APIKey.Value

	url := fmt.Sprintf("%s?lat=%v&lon=%v&appid=%s", baseURL, loc.Lat, loc.Lon, apiKey)

	body, err := fetchData(url)

	if err != nil {
		return nil, err
	}

	return util.ParseWeatherData(body)

}

func Geocode(city string, loader config.ConfigLoader) (location *types.Coord, err error) {

	cfg, err := loader.LoadConfig()

	if err != nil {
		return nil, err
	}

	geoConfig, ok := cfg.(*config.GeoConfig)

	if !ok {
		return nil, errors.New("unexpected config type")
	}

	baseURL := geoConfig.BaseURL.URL
	apiKey := geoConfig.APIKey.Value

	url := fmt.Sprintf("%s?q=%s&limit=5&appid=%s", baseURL, city, apiKey)

	body, err := fetchData(url)

	if err != nil {
		return nil, err
	}

	return util.ParseLocationInfo(body)

}
