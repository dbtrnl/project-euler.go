package utils

import (
	"math"
	"strconv"
)

func ReturnFibonacciNumbersUntil(limit int) []int {
	fibArray := []int{1, 2}

	iterator := 2
	currentNumber := fibArray[iterator-1] + fibArray[iterator-2]

	for currentNumber <= limit {
		fibArray = append(fibArray, currentNumber)
		iterator++
		currentNumber = fibArray[iterator-1] + fibArray[iterator-2]
	}
	return fibArray
}

func FindLargestPrimeFactor(inputNum int) int {
	num := inputNum

	for num%2 == 0 { num /= 2 }

	for i := 3; i <= int(math.Sqrt(float64(num))); i += 2 {
		for num%i == 0 { num /= i }
	}
	return num
}

func IsNumberPalindrome(inputNumber int) bool {
	number := strconv.Itoa(inputNumber)
	inverseNum := ""

	for i := len(number) - 1; i >= 0; i-- {
		inverseNum += string(number[i])
	}
	return number == inverseNum
}

func IsNumberEvenlyDivisibleBy(inputNum, divisor int) bool {
	if inputNum <= 0 || divisor <= 0 {
		// Refactor to return an error
		panic("Both arguments must be greater than zero!")
	}
	if inputNum % divisor == 0 { return true }
	return false
}

func IsEvenlyDivisibleByEveryNumberInInterval(num, interval_start, interval_end int, order string) bool {
	result := false
	var isDivisible bool
	// Add error checking for invalid intervals

	if (order == "asc") {
		for i := interval_start; i <= interval_end; i++ {
			isDivisible = IsNumberEvenlyDivisibleBy(num, i)
			if (!isDivisible) { return false }
		}
		result = true
	}
	if (order == "desc") {
		for i := interval_end; i >= interval_start; i-- {
			isDivisible = IsNumberEvenlyDivisibleBy(num, i)
			if (!isDivisible) { return false }
		}
		result = true
	}
	return result
}

/*
  if (order === "descending") {
    for (let i = intervalEnd; i >= intervalStart; i--) {
      isDivisible = isNumberEvenlyDivisibleBy(number, i)
      if (!isDivisible) return false
    }
    result = true
  }

  return result
} */