package utils

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"unicode"
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

	for num%2 == 0 {
		num /= 2
	}

	for i := 3; i <= int(math.Sqrt(float64(num))); i += 2 {
		for num%i == 0 {
			num /= i
		}
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
	if inputNum%divisor == 0 {
		return true
	}
	return false
}

func IsEvenlyDivisibleByEveryNumberInInterval(num, interval_start, interval_end int, order string) bool {
	result := false
	var isDivisible bool
	// Add error checking for invalid intervals

	if order == "asc" {
		for i := interval_start; i <= interval_end; i++ {
			isDivisible = IsNumberEvenlyDivisibleBy(num, i)
			if !isDivisible {
				return false
			}
		}
		result = true
	}
	if order == "desc" {
		for i := interval_end; i >= interval_start; i-- {
			isDivisible = IsNumberEvenlyDivisibleBy(num, i)
			if !isDivisible {
				return false
			}
		}
		result = true
	}
	return result
}

func FindSumOfNumberIntervalSquares(intervalStart, intervalEnd int) int {
	result := 0
	for i := intervalStart; i <= intervalEnd; i++ {
		result += int(math.Pow(float64(i), 2.0))
	}
	return result
}

func FindSquareOfNumberIntervalSum(intervalStart, intervalEnd int) int {
	result := 0
	for i := intervalStart; i <= intervalEnd; i++ {
		result += i
	}
	return int(math.Pow(float64(result), 2.0))
}

// Returns if a given number N is prime.
func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}
	// 2 and 3 are both primes
	if num <= 3 {
		return true
	}
	// Research this property of primes... Every prime is odd, but what about %3?
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	for i := 5; i*i <= num; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Finds the "nth" prime number
func FindNthPrime(nthPrime int) int {
	var primeArray []int
	currNum := 0

	for len(primeArray) < nthPrime {
		isCurrNumPrime := IsPrime(currNum)
		if !isCurrNumPrime {
			currNum++
			continue
		}
		primeArray = append(primeArray, currNum)
		currNum++
	}
	return primeArray[nthPrime-1]
}

// Multiplies every number in a given continuous digit series
func FindProductOfDigitsInNumberSeries(digitSeries string) int {
	product := 1
	for _, digitStr := range digitSeries {
		digit, _ := strconv.Atoi(string(digitStr))
		product *= digit
	}
	return product
}

// Finds all prime numbers smaller than N.
func FindAllPrimesSmallerThan(n int) []int {
	var primeArr []int
	iterator := 0
	if n == 0 {
		return primeArr
	}

	for iterator < n {
		if IsPrime(iterator) {
			primeArr = append(primeArr, iterator)
		}
		iterator++
	}
	return primeArr
}

// Formula to generate the "nth" triangular number.
func GenerateTriangularNumber(nth int) int {
	return nth * (nth + 1) / 2
}

// Finds all the proper divisors of a number including the number itself.
func FindAllDivisorsOf(num int) ([]int, error) {
	if num < 0 {
		return nil, errors.New(fmt.Sprintf("invalid value: %v", num))
	}
	divisors := []int{}

	if num == 0 {
		return divisors, nil
	}
	divisors = append(divisors, 1) // 1 is a divisor of every number > 0

	for i := 2; i <= num; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors, nil
}

// Finds all the divisors of a number except the number itself.
func FindAllProperDivisorsOf(num int) ([]int, error) {
	if num < 0 {
		return nil, errors.New(fmt.Sprintf("invalid value: %v", num))
	}
	divisors := []int{}
	numSqrt := int(math.Sqrt(float64(num)))

	if num == 0 {
		return divisors, nil
	}
	divisors = append(divisors, 1) // 1 is a divisor of every number > 0

	for i := 2; i <= numSqrt; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
			if i != num/i {
				divisors = append(divisors, num/i)
			}
		}
	}
	return divisors, nil
}

// Returns a sequence of numbers calculated according to the Collatz Conjecture
// Source: en.wikipedia.org/wiki/Collatz_conjecture
func FindCollatzSequenceOf(inputNumber int) []int {
	sequence := []int{inputNumber}
	num := inputNumber

	for num > 1 {
		if num%2 == 0 {
			num /= 2
		} else {
			num = 3*num + 1
		}
		sequence = append(sequence, num)
	}
	return sequence
}

// Returns the longest Collatz Sequence of all numbers under limit N
func FindLongestCollatzSequenceUnder(limit int) []int {
	var longestSequence []int

	for i := 1; i < limit; i++ {
		currSequence := FindCollatzSequenceOf(i)
		if len(currSequence) > len(longestSequence) {
			longestSequence = currSequence
		}
	}
	return longestSequence
}

func Factorial(num int) (int, error) {
	var result int
	// Errors and edge cases
	if num <= 0 {
		return 0, errors.New(fmt.Sprintf("invalid number. should be greater than zero: %v", num))
	}
	if num > 20 {
		return 0, errors.New(fmt.Sprintf("result of %v! is too big for int/int32. use 'FactorialBig()' instead.", num))
	}
	if num == 1 {
		return 1, nil
	}
	if num == 2 {
		return 2, nil
	}

	for i := num; i > 1; i-- {
		if i == num {
			result += i * (i - 1)
		} else {
			result *= i - 1
		}
	}
	return result, nil
}

func FactorialBig(n int) (*big.Int, error) {
	num, result := new(big.Int), new(big.Int)
	one := big.NewInt(1)

	num.SetString(strconv.Itoa(n), 10) // setting num to value of n

	isNSmallerOrEqualToZero := num.Sign() == -1 && num.Sign() == 0
	isNOneOrTwo := num.Cmp(big.NewInt(1)) == 0 || num.Cmp(big.NewInt(2)) == 0

	// Errors and edge cases
	if isNSmallerOrEqualToZero {
		return big.NewInt(0), errors.New(fmt.Sprintf("invalid number. should be greater than zero: %v", num))
	}
	if isNOneOrTwo {
		return num, nil
	}

	for i := new(big.Int).Set(num); i.Cmp(one) > 0; i.Sub(i, one) {
		isIEqualToNum := i.Cmp(num) == 0
		if isIEqualToNum {
			// Unlike in a normal for loop with int, here i can't be changed, so new variables are needed
			iMinusOne := new(big.Int).Set(i)
			iMinusOne.Sub(iMinusOne, one) // i-1
			currI := new(big.Int).Set(i)
			currI.Mul(i, iMinusOne) // i * (i-1)
			result.Add(result, currI)
		} else {
			iMinusOne := new(big.Int).Set(i)
			iMinusOne.Sub(iMinusOne, one) // i-1
			result.Mul(result, iMinusOne)
		}
	}
	return result, nil
}

type numberToWrittenNumber map[string]string

var numberMap = numberToWrittenNumber{
	"0":   "zero",
	"1":   "one",
	"2":   "two",
	"3":   "three",
	"4":   "four",
	"5":   "five",
	"6":   "six",
	"7":   "seven",
	"8":   "eight",
	"9":   "nine",
	"10":  "ten",
	"11":  "eleven",
	"12":  "twelve",
	"13":  "thirteen",
	"14":  "fourteen",
	"15":  "fifteen",
	"16":  "sixteen",
	"17":  "seventeen",
	"18":  "eighteen",
	"19":  "nineteen",
	"20":  "twenty",
	"30":  "thirty",
	"40":  "forty",
	"50":  "fifty",
	"60":  "sixty",
	"70":  "seventy",
	"80":  "eighty",
	"90":  "ninety",
	"00":  "hundred",
	"000": "thousand",
}

func CountLettersInString(s string) int {
	count := 0
	for _, char := range s {
		if unicode.IsLetter(char) {
			count++
		}
	}
	return count
}

func ReturnNumberString(num int) string {
	var currStr string
	strNum := strconv.Itoa(num)
	numLen := len(strNum)

	if numLen == 0 {
		return numberMap["0"]
	}

	if numLen == 1 {
		return numberMap[strNum]
	}

	if numLen == 2 {
		firstDigit := strNum[0]
		secondDigit := strNum[1]
		isSecondDigitZero := secondDigit == 48
		// fmt.Println("secondDigit", secondDigit)
		if isSecondDigitZero {
			return numberMap[strNum]
		}
		if numberMap[strNum] != "" {
			return numberMap[strNum]
		}

		switch firstDigit {
		case '2':
			currStr = fmt.Sprintf("%s %s", numberMap["20"], numberMap[string(secondDigit)])
			break
		case '3':
			currStr = fmt.Sprintf("%s %s", numberMap["30"], numberMap[string(secondDigit)])
			break
		case '4':
			currStr = fmt.Sprintf("%s %s", numberMap["40"], numberMap[string(secondDigit)])
			break
		case '5':
			currStr = fmt.Sprintf("%s %s", numberMap["50"], numberMap[string(secondDigit)])
			break
		case '6':
			currStr = fmt.Sprintf("%s %s", numberMap["60"], numberMap[string(secondDigit)])
			break
		case '7':
			currStr = fmt.Sprintf("%s %s", numberMap["70"], numberMap[string(secondDigit)])
			break
		case '8':
			currStr = fmt.Sprintf("%s %s", numberMap["80"], numberMap[string(secondDigit)])
			break
		case '9':
			currStr = fmt.Sprintf("%s %s", numberMap["90"], numberMap[string(secondDigit)])
			break
		}
	}

	if numLen == 3 {
		firstDigit := strNum[0]
		twoLastDigitsStr := strNum[1:]
		twoLastDigits, _ := strconv.Atoi(twoLastDigitsStr)
		currStr = fmt.Sprintf("%s %s", numberMap[string(firstDigit)], numberMap["00"])

		switch twoLastDigitsStr {
		case "00":
			break
		default:
			if numberMap[twoLastDigitsStr] != "" {
				currStr = fmt.Sprintf("%s and %s", currStr, numberMap[twoLastDigitsStr])
			} else {
				currStr = fmt.Sprintf("%s and %s", currStr, ReturnNumberString(twoLastDigits))
			}
		}
	}

	if numLen == 4 {
		firstDigit := strNum[0]
		isSecondDigitZero := strNum[1] == 48
		threeLastDigitsStr := strNum[1:]
		threeLastDigits, _ := strconv.Atoi(threeLastDigitsStr)
		currStr = fmt.Sprintf("%s %s", numberMap[string(firstDigit)], numberMap["000"])

		switch threeLastDigitsStr {
		case "000":
			break
		default:
			if isSecondDigitZero {
				currStr = fmt.Sprintf("%s and %s", currStr, ReturnNumberString(threeLastDigits))
			} else {
				currStr = fmt.Sprintf("%s %s", currStr, ReturnNumberString(threeLastDigits))
			}
		}
	}
	/* TODO
	if numLen == 5 {
		firstTwoDigits := strNum[0:2]
		currStr = fmt.Sprintf("%s %s", numberMap[string(firstTwoDigits)], numberMap["000"])
	}
	*/
	return currStr
}

func FindAndSumAllProperDivisorsOf(num int) (int, error) {
	if num == 0 {
		return 0, nil
	}
	if num == 1 {
		return 0, nil
	}
	if num == 2 {
		return 1, nil
	}

	sum := 0
	properDivisors, err := FindAllProperDivisorsOf(num)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("error on FindAllProperDivisorsOf(): %v", err.Error()))
	}
	for _, n := range properDivisors {
		sum += n
	}
	return sum, nil
}

type AmicableNumberObject struct {
	IsAmicable bool
	Numbers    [2]int
}

func IsNumberAmicable(num int) (*AmicableNumberObject, error) {
	result := AmicableNumberObject{
		IsAmicable: false,
		Numbers:    [2]int{},
	}
	if num < 0 {
		return &result, errors.New(fmt.Sprintf("invalid number %v. should be greater than zero.", num))
	}

	sumDiv, err := FindAndSumAllProperDivisorsOf(num)
	if err != nil {
		return &result, errors.New(fmt.Sprintf("error on FindAndSumAllProperDivisorsOf(): %v", err.Error()))
	}
	sumDivSum, err := FindAndSumAllProperDivisorsOf(sumDiv)
	if err != nil {
		return &result, errors.New(fmt.Sprintf("error on FindAndSumAllProperDivisorsOf(): %v", err.Error()))
	}
	if sumDivSum == num && sumDiv != num {
		result.IsAmicable = true
		result.Numbers = [2]int{num, sumDiv}
		return &result, nil
	}

	// if sumDivSum, err := FindAndSumAllProperDivisorsOf(sumDiv); (err != nil && sumDivSum == num) {
	// 	result.IsAmicable = true
	// 	result.Numbers = [2]int{num, sumDiv}
	// 	return &result, nil
	// }

	return &result, nil
}
