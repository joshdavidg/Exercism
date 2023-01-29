package booking

import (
	"fmt"
	"log"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// expected input format to parse: M/D/YYYY 24hr:MM:SS
	timeString := "1/2/2006 15:04:05"
	time, err := time.Parse(timeString, date)
	if err != nil {
		log.Fatal(err)
	}
	return time
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// expected input format to parse: Month D, YYYY 24hr:MM:SS
	formatString := "January 2, 2006 15:04:05"
	formatted, err := time.Parse(formatString, date)
	if err != nil {
		log.Fatal(err)
	}
	return formatted.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// expected input format to parse: DayOfWeek, Month D, YYYY 24hr:MM:SS
	formatString := "Monday, January 2, 2006 15:04:05"
	formatted, err := time.Parse(formatString, date)
	if err != nil {
		log.Fatal(err)
	}
	appointmentHr := formatted.Hour()
	return appointmentHr >= 12 && appointmentHr < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	scheduledAppt := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s, at %s.", scheduledAppt.Format("Monday, January 2, 2006"), scheduledAppt.Format("15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
