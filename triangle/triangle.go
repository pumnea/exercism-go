// Package triangle provides utilities for classifying triangles.
package triangle

// Kind represents the type of a triangle.
type Kind int

// Triangle classification constants.
const (
	NaT = iota // Not a triangle
	Equ        // Equilateral (all sides equal)
	Iso        // Isosceles (at least two sides equal, not all)
	Sca        // Scalene (no sides equal)
)

// KindFromSides determines the type of triangle formed by sides a, b, and c.
// Returns:
// - NaT if the sides cannot form a valid triangle (e.g., negative, zero, or violate triangle inequality).
// - Equ if all three sides are equal (equilateral).
// - Iso if exactly two sides are equal (isosceles, excluding equilateral).
// - Sca if no sides are equal (scalene).
func KindFromSides(a, b, c float64) Kind {
	// Check for invalid triangle (non-positive sides or triangle inequality failure)
	if a <= 0 || b <= 0 || c <= 0 || a+b <= c || b+c <= a || a+c <= b {
		return NaT
	}

	// Check for equilateral (all sides equal)
	if a == b && b == c {
		return Equ
	}

	// Check for isosceles (exactly two sides equal)
	if a == b || b == c || a == c {
		return Iso
	}
	// If not invalid, equilateral, or isosceles, it must be scalene
	return Sca
}
