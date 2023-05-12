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
			reader := bufio.NewReader(os.Stdin)
			for {

				fmt.Print("Please enter your OpenWeather API key: ")
				apiKey, _ := reader.ReadString('\n')
				apiKey = strings.TrimSpace(apiKey)

				if apiKey == "" {
					fmt.Print("API key cannot be empty, enter a valid api key...\n")
					continue
				}

				geoLoader.APIKey.Value = apiKey
				weatherLoader.APIKey.Value = apiKey

				for {
					fmt.Print("Please enter the city name (or type 'exit' to quit): ")
					city, _ := reader.ReadString('\n')
					city = strings.TrimSpace(city)

					if city == "" || util.IsNumeric(city) {
						fmt.Print("Please enter a valid city name.\n")
						continue
					}

					if city == "exit" {
						os.Exit(0)
					}
					util.Capitalize(&city)

					coords, err := api.Geocode(city, geoLoader)

					if err != nil {
						if strings.Contains(err.Error(), "401") {
							fmt.Println("Invalid API key. Please check your key and try again.")
							break
						} else {
							fmt.Printf("An error occurred: %s\n", err.Error())
						}
						continue
					}

					weatherData, err := api.FetchWeather(coords, weatherLoader)

					fmt.Printf("Weather information for %s:\n", city)
					fmt.Printf("Temperature: %.2f°C\n", util.ConvertKelvintoCelsius(weatherData.Main.Temp))
					fmt.Printf("Condition outside is: %s\n", weatherData.Weather[0].Main)
					fmt.Printf("Humidity: %v\n", weatherData.Main.Humidity)
				}
			}
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	return nil
}
