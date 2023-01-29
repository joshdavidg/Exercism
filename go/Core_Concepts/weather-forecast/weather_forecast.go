//Package weather provides current weather conditions for a city.
package weather

//CurrentCondition is used to indicate the current weather condition for the city.
var CurrentCondition string

//CurrentLocation is used to indicate the current location for the weather forecast.
var CurrentLocation string

//Forecast takes in a city and condition and returns a forecast of the current weather for the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
