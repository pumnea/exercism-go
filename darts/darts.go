// Package darts provides utilities to calculate points in a darts game.
package darts

// Score calculates the dart score based on the x, y coordinates of where the dart lands.
// The dartboard is centered at (0, 0) with concentric circles:
// - Radius ≤ 1: 10 points (bullseye)
// - Radius ≤ 5: 5 points (middle ring)
// - Radius ≤ 10: 1 point (outer ring)
// - Beyond radius 10: 0 points
func Score(x, y float64) int {
	// Squared distance from origin
	distanceSquared := x*x + y*y

	// Precomputed squared radii
	const (
		oneSquared  = 1.0
		fiveSquared = 25.0
		tenSquared  = 100.0
	)

	switch {
	case distanceSquared <= oneSquared:
		return 10
	case distanceSquared <= fiveSquared:
		return 5
	case distanceSquared <= tenSquared:
		return 1
	default:
		return 0
	}
}
