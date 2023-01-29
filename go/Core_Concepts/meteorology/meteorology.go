package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// TemperatureUnit Stringer implementation
func (tu TemperatureUnit) String() string {
	if tu == Celsius {
		return fmt.Sprint("°C")
	}

	return fmt.Sprint("°F")
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Temperature struct Stringer implementation
func (t Temperature) String() string {
	return fmt.Sprintf("%d %s", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// SpeedUnit Stringer implementation
func (su SpeedUnit) String() string {
	if su == KmPerHour {
		return fmt.Sprint("km/h")
	}

	return fmt.Sprint("mph")
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Speed Stringer implementation
func (s Speed) String() string {
	return fmt.Sprintf("%d %s", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Meteorology Stringer implementation
func (m MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", m.location,
		m.temperature, m.windDirection, m.windSpeed, m.humidity)
}
