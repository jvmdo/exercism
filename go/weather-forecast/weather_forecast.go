// Package weather forecasts current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition stores the current weather condition of a city.
var CurrentCondition string
// CurrentLocation stores the city name.
var CurrentLocation string

// Forecast receives the city name and condition and returns a string containing the weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
