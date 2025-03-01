package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layouts := []string{
		"January 2, 2006 15:04:05",
		"1/2/2006 15:04:05",
		"2006-01-02 15:04:05",
		"02 Jan 2006 15:04:05",
		"Monday, January 2, 2006 15:04:05",
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, date); err == nil {
			return t
		}
	}
	return time.Time{}
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	return time.Now().After(Schedule(date))
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	appTime := Schedule(date).Hour()
	return appTime >= 12 && appTime <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	app := Schedule(date)
	weekday := app.Weekday().String()
	month := app.Month().String()
	day := app.Day()
	year := app.Year()
	hour, minute, _ := app.Clock()
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %02d:%02d.", weekday, month, day, year, hour, minute)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	day := 15
	month := time.September
	year := time.Now().Year()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
