package util

import (
	"strconv"
	"strings"
)

func ConvertKelvintoCelsius(temp float64) float64 {
	return temp - 273.15
}
func CheckNumber(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

func Capitalize(s *string) {
	*s = strings.Title(*s)
}

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
