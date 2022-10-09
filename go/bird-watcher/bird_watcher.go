package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, num := range birdsPerDay {
		total += num
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	startIndex := 7 * (week - 1)
	total := 0
	for i := startIndex; i < startIndex+7; i++ {
		total += birdsPerDay[i]
	}
	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, num := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] = num + 1
		}
	}
	return birdsPerDay
}
