package purchase

import (
	"fmt"
	"sort"
)

// reductionUpToThreeYears represents the percentage price reduction for a car from 0 to 3 years old.
const reductionUpToThreeYears = 0.8

// reductionAfterTenYears represents the percentage price reduction for a car 10 years or older.
const reductionAfterTenYears = 0.5

// reductionBetweenThreeAndTenYears represents the percentage price reduction for a car between 3 and 10 years old.
const reductionBetweenThreeAndTenYears = 0.7

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	vehicles := []string{option1, option2}
	sort.Strings(vehicles)
	return fmt.Sprintf("%s is clearly the better choice.", vehicles[0])
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * reductionUpToThreeYears
	}
	if age >= 10 {
		return originalPrice * reductionAfterTenYears
	}
	// Price reduction if between 3 and < 10 years
	return originalPrice * reductionBetweenThreeAndTenYears
}
