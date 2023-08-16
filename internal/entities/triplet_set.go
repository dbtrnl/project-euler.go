package entities

import (
	"math"
)

// A triplet set of numbers A, B and C.
type TripletSet struct {
	A, B, C int
}

// A Pythagorean triplet is a set of three natural numbers, A < B < C, for which A² + B² = C².
func (t *TripletSet) IsPythagoreanTriplet() bool {
	if t.A > t.B || t.A > t.C {
		return false
	}
	if t.B > t.C {
		return false
	}
	if math.Pow(float64(t.A), 2.0)+math.Pow(float64(t.B), 2.0) != math.Pow(float64(t.C), 2.0) {
		return false
	}
	return true
}
