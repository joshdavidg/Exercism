package cars

// costPerCar represents the amount of money required to produce one car.
const costPerCar = 10000

//costPer10Cars is the amount of money require to produce a group of 10 cars.
const costPer10Cars = 95000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	percentage := successRate / 100
	return float64(productionRate) * percentage
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupsOf10 := carsCount / 10
	remainingCars := carsCount % 10
	return uint(groupsOf10)*costPer10Cars + uint(remainingCars)*costPerCar
}
