// Package weather provides information about
// the weather condition for the given location
// Its a tool.
package weather

var (
	// CurrentCondition provide info about condition.
	CurrentCondition string
	// CurrentLocation provide info about location.
	CurrentLocation string
)

// Forecast returns the currentLocation and the condition of the temperature there!.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
