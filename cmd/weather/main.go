package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"weather/internal/config"
	"weather/internal/util"
	"weather/pkg/api"

	"github.com/urfave/cli/v2"
)

func main() {

	geoLoader := &config.GeoConfigLoader{}
	weatherLoader := &config.WeatherConfigLoader{}

	err := runCli(geoLoader, weatherLoader)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func runCli(geoLoader *config.GeoConfigLoader, weatherLoader *config.WeatherConfigLoader) error {
	app := &cli.App{
		Name:  "Weather CLI",
		Usage: "Get weather information for a city",
		Action: func(c *cli.Context) error {
			for {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Please enter the city name (or type 'exit' to quit): ")
				city, _ := reader.ReadString('\n')
				city = strings.TrimSpace(city)

				if city == "" || util.IsNumeric(city) {
					fmt.Print("Please enter a valid city name.\n")
					continue
				}

				if city == "exit" {
					break
				}
				util.Capitalize(&city)
				coords, err := api.Geocode(city, geoLoader)

				if err != nil {
					fmt.Printf("Error fetching weather data: %s\n", err.Error())
					continue
				}

				weatherData, err := api.FetchWeather(coords, weatherLoader)

				fmt.Printf("Weather information for %s:\n", city)
				fmt.Printf("Temperature: %.2f°C\n", util.ConvertKelvintoCelsius(weatherData.Main.Temp))
				fmt.Printf("Humidity: %v\n", weatherData.Main.Humidity)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	return nil
}
