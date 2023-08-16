package utils

import (
	"math"
	"strconv"

	"github.com/dbtrnl/project-euler/golang/internal/entities"
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

func FindSumOfNumberIntervalSquares(intervalStart, intervalEnd int) int {
  result := 0
  for i := intervalStart; i <= intervalEnd; i++{
		result += int(math.Pow(float64(i), 2.0))
	}
  return result
}

func FindSquareOfNumberIntervalSum(intervalStart, intervalEnd int) int {
  result := 0
	for i := intervalStart; i <= intervalEnd; i++{
		result += i
	}
  return int(math.Pow(float64(result), 2.0))
}

func IsPrime(num int) bool {
	if (num <= 1) {
		return false
	}
  // 2 and 3 are both primes
  if (num <= 3) { return true }
  // Research this property of primes... Every prime is odd, but what about %3?
  if (num % 2 == 0 || num % 3 == 0) { return false }
  for i := 5; i * i <= num; i += 6 {
    if (num % i == 0 || num % (i + 2) == 0) { return false }
  }
  return true
}

func FindNthPrime(nthPrime int) int {
  var primeArray []int
  currNum := 0

	for len(primeArray) < nthPrime {
		isCurrNumPrime := IsPrime(currNum)
		if !isCurrNumPrime { currNum++; continue }
		primeArray = append(primeArray, currNum); currNum++
	}
  return primeArray[nthPrime - 1]
}

func FindProductOfDigitsInNumberSeries(digitSeries string) int {
	product := 1
	for _, digitStr := range digitSeries {
		digit, _ := strconv.Atoi(string(digitStr))
		product *= digit
	}
	return product
}

func IsSetPythagoreanTriplet(inputSet entities.TripletSet) bool {
  aSquared := math.Pow(float64(inputSet.A), 2.0)
	bSquared := math.Pow(float64(inputSet.B), 2.0)
	cSquared := math.Pow(float64(inputSet.C), 2.0)

  if (aSquared + bSquared == cSquared) {
		return true
	}
  return false
}

func FindAllPrimesSmallerThan(limit int) []int {
  var primeArr []int
	iterator := 0
  if (limit == 0) { return primeArr }

	for iterator < limit {
		if IsPrime(iterator) { primeArr = append(primeArr, iterator) }
		iterator++
	}
  return primeArr
}