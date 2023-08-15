package utils

import (
	"testing"

	"github.com/dbtrnl/project-euler/golang/internal/entities.go"
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

func TestFindSumOfNumberIntervalSquares(t *testing.T) {
	assert.Equal(t, 385, FindSumOfNumberIntervalSquares(1, 10))
}

func TestFindSquareOfNumberIntervalSum(t *testing.T) {
	assert.Equal(t, 3025, FindSquareOfNumberIntervalSum(1, 10))
}

func TestFindNthPrime(t *testing.T) {
	assert.Equal(t, 2, FindNthPrime(1))
	assert.Equal(t, 11, FindNthPrime(5))
	assert.Equal(t, 229, FindNthPrime(50))
	assert.Equal(t, 86719, FindNthPrime(8428))
	assert.Equal(t, 104743, FindNthPrime(10001))
}

func TestFindProductOfDigitsInNumberSeries(t *testing.T) {
	assert.Equal(t, 6, FindProductOfDigitsInNumberSeries("123"))
	assert.Equal(t, 120, FindProductOfDigitsInNumberSeries("12345"))
	assert.Equal(t, 4354560, FindProductOfDigitsInNumberSeries("8652393278"))
	assert.Equal(t, 0, FindProductOfDigitsInNumberSeries("86523932780"))
}

func TestIsSetPythagoreanTriplet(t *testing.T) {
	assert.Equal(t, false, IsSetPythagoreanTriplet(entities.TripletSetObject{A:1, B:2, C:3}))
	assert.Equal(t, true, IsSetPythagoreanTriplet(entities.TripletSetObject{A:3, B:4, C:5}))
}

func TestFindAllPrimesSmallerThan(t *testing.T) {
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79,
		83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
	}
	assert.Equal(t, expected, FindAllPrimesSmallerThan(174))
	assert.Equal(t, expected, FindAllPrimesSmallerThan(179))

	expected = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79,
		83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179,
	}
	assert.Equal(t, expected, FindAllPrimesSmallerThan(180))
}