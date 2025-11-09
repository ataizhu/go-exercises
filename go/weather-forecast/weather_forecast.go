// Package weather provides tools to store and present simple current weather forecasts for cities.
package weather

// CurrentCondition holds a short description of the observed weather, such as "sunny" or "rainy".
var CurrentCondition string

// CurrentLocation holds the name of the city for which the weather information applies.
var CurrentLocation string

// Forecast records the provided city and condition and returns a formatted forecast message.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
