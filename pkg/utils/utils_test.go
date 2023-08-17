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
		{10, 5},
		{13195, 29},
		{600851475143, 6857},
	}

	for _, test := range tests {
		actual := FindLargestPrimeFactor(test.input)
		assert.Equal(t, test.expected, actual, "For input %d", test.input)
	}
}

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

func TestFindAllProperDivisorsOf(t *testing.T) {
	expected := []int{1}
	result, err := FindAllProperDivisorsOf(1)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1}
	result, err = FindAllProperDivisorsOf(2)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1, 2, 4}
	result, err = FindAllProperDivisorsOf(8)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1, 2, 4, 5, 10, 20, 25, 50}
	result, err = FindAllProperDivisorsOf(100)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1, 2, 4, 5, 8, 10, 20, 25, 40, 50, 100, 125, 200, 250, 500}
	result, err = FindAllProperDivisorsOf(1000)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1, 7, 11, 13, 77, 91, 143}
	result, err = FindAllProperDivisorsOf(1001)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	expected = []int{1, 17, 59}
	result, err = FindAllProperDivisorsOf(1003)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result)

	result, err = FindAllProperDivisorsOf(-1)
	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid value: -1")
}
