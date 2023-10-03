package booking

import (
	"time"
	"fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)

	return t

	// Pro solution. Schedule can be used by others methods if implemented like this
	// for _,layout := range [...]string {
	// 	"1/2/2006 15:04:05",
	// 	"January 2, 2006 15:04:05",
	// 	"Monday, January 2, 2006 15:04:05",
	// } {
	// 	t,err := time.Parse(layout, date)

	// 	if (err == nil) {
	// 				return t
	// 	}
	// }

	// panic("Unknown date format: " + date)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	targetDate, _ := time.Parse("January 2, 2006 15:04:05", date)
	
	return targetDate.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	targetDate, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	targetHour := targetDate.Hour()

	return targetHour >= 12 && targetHour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	
	day := t.Day()
	month := t.Month()
	year := t.Year()
	dayOfWeek := t.Weekday()
	timeStr := t.Format("15:04")
	
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %s.", dayOfWeek, month, day, year, timeStr)

	// Better solution
	// d, _ := time.Parse("1/2/2006 15:04:05", date)
	// return fmt.Sprintf("You have an appointment on %s", d.Format("Monday, January 2, 2006, at 15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
