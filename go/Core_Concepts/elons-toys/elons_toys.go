package elon

import (
	"fmt"
)

//Drive increases distance by the speed of the car and reduces batter power by the batteryDrain
func (c *Car) Drive() {
	if c.battery < c.batteryDrain {
		return
	}
	c.battery = c.battery - c.batteryDrain
	c.distance = c.distance + c.speed
}

//DisplayDistance displays the current distance the car has driven
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

//DisplayBattery displays the current battery percentage
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

//CanFinish determines whether a car can finish race given the track distance
func (c *Car) CanFinish(trackDistance int) bool {
	return (c.battery/c.batteryDrain)*c.speed >= trackDistance
}
