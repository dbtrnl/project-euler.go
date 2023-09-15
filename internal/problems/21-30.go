package problems

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"sync"

	"github.com/dbtrnl/project-euler.go/pkg/utils"
)

// Problem 21 - Amicable numbers
//
// Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
//
// If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.
//
// For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284.
//
// The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.
//
// Evaluate the sum of all the amicable numbers under 10000.
func Problem21(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	var answer int
	amicableNums := make(map[int]bool)
	for i := 1; i <= 10000; i++ {
		if amicableNums[i] {
			continue
		}
		amicability, _ := utils.IsNumberAmicable(i)
		if amicability.IsAmicable {
			amicableNums[i] = true
			amicableNums[amicability.Numbers[1]] = true
		}
	}
	for num, _ := range amicableNums {
		answer += num
	}
	ch <- answer
}

// Problem 22 - Names scores
//
// Using a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order.
//
// Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.
//
// For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list.
// So, COLIN would obtain a score of 938 × 53 = 49714.
//
// What is the total of all the name scores in the file?
func Problem22(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	var totalScore, strScore int
	filePath := "./internal/input_data/problem22_input.txt"
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("error reading file %s: %v\n", filePath, err)
		ch <- 0
	}
	names := strings.Split(strings.ReplaceAll(string(f), `"`, ""), ",")
	sort.Strings(names)

	for i, n := range names {
		pos := i + 1
		strScore = utils.CalculateStringScore(n)
		totalScore += strScore * pos
	}
	ch <- totalScore
}

// Problem 23 - Non-abundant sums
//
// A perfect number is a number for which the sum of its proper divisors is exactly equal to the number.
// For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.
//
// A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.
//
// As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24.
//
// By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers.
// However, this upper limit cannot be reduced any further by analysis even though it is known that
// the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.
//
// Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
func Problem23(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	var answer int
	limit := 28123
	abundantNums, _ := utils.FindAbundantNumbersUntil(limit)

	nums := utils.FindUniqueCombinatorialSums(abundantNums, limit)

	for i := 0; i <= limit; i++ {
		if nums[i] {
			continue
		}
		answer += i
	}
	ch <- answer
}

// Problem 24 - Lexicographic permutations
//
// A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4.
// If all of the permutations are listed numerically or alphabetically, we call it lexicographic order.
//
// The lexicographic permutations of 0, 1 and 2 are:
//
// 012   021   102   120   201   210
//
// What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
func Problem24(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	// Not implemented
	ch <- -1
}

// Problem 25 - 1000-digit Fibonacci number
//
// The Fibonacci sequence is defined by the recurrence relation:
//
// Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
//
// Hence the first 12 terms will be:
//
// F(1)=1, F(2)=1, F(3)=2, F(4)=3, F(5)=5, F(6)=8, F(7)=13, F(8)=21, F(9)=34, F(10)=55, F(11)=89, F(12)=144
//
// The 12th term, F12, is the first term to contain three digits.
//
// What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
func Problem25(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	var answer int
	limit := 5000
	for i := 0; i <= limit; i++ {
		seq, _ := utils.FibonacciSequenceBigInt(i)
		lastNum := seq[len(seq)-1]
		if len(lastNum.String()) >= 1000 {
			answer = i
			break
		}
	}
	ch <- answer
}

func Problem26(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	ch <- -1
}

func Problem27(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	ch <- -1
}

func Problem28(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	ch <- -1
}

func Problem29(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	ch <- -1
}

func Problem30(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	ch <- -1
}
