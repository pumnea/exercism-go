// Package clock provides a simple 24-hour clock without dates.
package clock

import "fmt"

// Clock represents a time of day in minutes since midnight (0 to 1439).
type Clock struct {
	minutes int
}

// New creates a new Clock from hours and minutes, normalizing to a 24-hour period.
func New(h, m int) Clock {
	totalMinutes := h*60 + m
	totalMinutes = totalMinutes % 1440
	if totalMinutes < 0 {
		totalMinutes += 1440
	}
	return Clock{minutes: totalMinutes}
}

// Add returns a new Clock with the specified minutes added.
func (c Clock) Add(m int) Clock {
	return New(0, c.minutes+m)
}

// Subtract returns a new Clock with the specified minutes subtracted.
func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

// String returns the time in "HH:MM" format (24-hour, zero-padded).
func (c Clock) String() string {
	hours := c.minutes / 60
	minutes := c.minutes % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
