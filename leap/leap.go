// Package leap provides utilities to check if a given year is leap.
package leap

// IsLeapYear returns true if input is a leap year.
// Divisible by 4 and either not by 100 or by 400.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
