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
	aSq := math.Pow(float64(t.A), 2.0)
	bSq := math.Pow(float64(t.B), 2.0)
	cSq := math.Pow(float64(t.C), 2.0)

	if aSq+bSq == cSq {
		return true
	}
	return false
}
