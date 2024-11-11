// Package weather contains methods to compute forecast for Goblinocus.
package weather

// CurrentCondition holds the current weather at a given location being
// analysed.
var CurrentCondition string

// CurrentLocation holds the value to the location being analysed.
var CurrentLocation string

// Forecast computes the current condition at a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
