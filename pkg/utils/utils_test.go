package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLargestPrimeFactor(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{10, 5},    // Largest prime factor of 10 is 5
		{13195, 29}, // Largest prime factor of 13195 is 29
		{600851475143, 6857}, // Largest prime factor of 600851475143 is 6857
	}

	for _, test := range tests {
		actual := FindLargestPrimeFactor(test.input)
		assert.Equal(t, test.expected, actual, "For input %d", test.input)
	}
}

// func TestFindLargestPrimeFactor(t *testing.T) {
// 	assert.Equal(t, 5, FindLargestPrimeFactor(10), "For input 10")
// 	assert.Equal(t, 29, FindLargestPrimeFactor(13195), "For input 13195")
// 	assert.Equal(t, 6857, FindLargestPrimeFactor(600851475143), "For input 600851475143")
// }
