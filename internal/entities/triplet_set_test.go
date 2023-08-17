package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSetPythagoreanTriplet(t *testing.T) {
	var set TripletSet
	set = TripletSet{A: 1, B: 2, C: 3}
	result := set.IsPythagoreanTriplet()
	assert.Equal(t, false, result)

	set = TripletSet{A: 3, B: 4, C: 5}
	result = set.IsPythagoreanTriplet()
	assert.Equal(t, true, result)
}
