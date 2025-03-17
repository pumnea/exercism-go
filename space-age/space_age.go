// Package space calculates ages in planetary years from seconds.
package space

import "strings"

type Planet string

const (
	Mercury Planet = "mercury"
	Venus   Planet = "venus"
	Earth   Planet = "earth"
	Mars    Planet = "mars"
	Jupiter Planet = "jupiter"
	Saturn  Planet = "saturn"
	Uranus  Planet = "uranus"
	Neptune Planet = "neptune"
)

// Seconds in an Earth year (365.25 days).
const earthSeconds = 31557600.0

var planets = map[Planet]float64{
	Mercury: 0.2408467,
	Venus:   0.61519726,
	Earth:   1.0,
	Mars:    1.8808158,
	Jupiter: 11.862615,
	Saturn:  29.447498,
	Uranus:  84.016846,
	Neptune: 164.79132,
}

// Age returns age in planetary years from seconds, or an error for invalid planets.
func Age(seconds float64, planet Planet) float64 {
	inpNormalize := Planet(strings.ToLower(string(planet)))
	period, exists := planets[inpNormalize]
	if !exists {
		return -1
	}
	return seconds / (earthSeconds * period)
}
