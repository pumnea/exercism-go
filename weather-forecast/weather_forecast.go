// Package weather provides tools
// to get weather for a certain location.
package weather

// CurrentCondition presents weather conditions.
var CurrentCondition string

// CurrentLocation presents the location for which the forecast is provided.
var CurrentLocation string

// Forecast returns a string holding data with weather condition for a given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
