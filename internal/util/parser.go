package util

import (
	"encoding/json"
	"errors"
	"weather/internal/types"
)

func ParseLocationInfo(data []byte) (*types.Coord, error) {
	var locInfo []types.CityLocData

	err := json.Unmarshal(data, &locInfo)

	if err != nil {
		return nil, err
	}

	if len(locInfo) == 0 {
		return nil, errors.New("location info not found")
	}

	location := &types.Coord{
		Lat: locInfo[0].Latitude,
		Lon: locInfo[0].Longitude,
	}

	return location, nil
}

func ParseWeatherData(data []byte) (*types.WeatherData, error) {
	var weatherData types.WeatherData
	err := json.Unmarshal(data, &weatherData)

	if err != nil {
		return nil, err
	}

	return &weatherData, nil
}
