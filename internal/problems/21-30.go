package problems

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

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
func Problem21() int {
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
	return answer
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
func Problem22() int {
	var totalScore, strScore int
	filePath := "./internal/input_data/problem22_input.txt"
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("error reading file %s: %v\n", filePath, err)
		return 0
	}
	names := strings.Split(strings.ReplaceAll(string(f), `"`, ""), ",")
	sort.Strings(names)

	for i, n := range names {
		pos := i + 1
		strScore = utils.CalculateStringScore(n)
		totalScore += strScore * pos
	}
	return totalScore
}
