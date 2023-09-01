package entities

import (
	"fmt"
)

type Triangle [][]int

func NewTriangle(data [][]int) (*Triangle, error) {
	var triangle Triangle
	rowLen := len(data[0])
	if rowLen == 0 {
		return nil, fmt.Errorf("invalid data at row 0: %v, triangle not created", data[0])
	}
	for i, d := range data {
		isCurrentRowCorrectSize := len(d) != rowLen
		if isCurrentRowCorrectSize && i != 0 {
			return nil, fmt.Errorf("invalid data at row %d: %v, triangle not created", i, d)
		}
		triangle = append(triangle, d)
		rowLen = rowLen + 1
	}
	return &triangle, nil
}

func (t Triangle) Display() {
	for _, row := range t {
		fmt.Println(row)
	}
	fmt.Println()
}

/*
// This solution gives the correct sum - 10, still have no idea why
// 75 95 47 87 82 75 73 28 83 47 43 73 91 67 98
func (t Triangle) FindMaximumSumPath() int {
	sum := t[0][0]
	prevIndex := 0
	for i, slice := range t {
		if i == 0 {
			continue
		}
		// fmt.Printf("Row %v: %v\n", i, slice)
		biggest, j := compareNumbers(slice[prevIndex], slice[prevIndex+1])
		prevIndex += j
		fmt.Printf("Curr sum: %v, summing num %v from line %v: %v\n", sum, biggest, i, slice)
		sum += biggest
	}
	return sum
}

func compareNumbers(a, b int) (int, int) {
	if a < b {
		return b, 1
	}
	if a > b {
		return a, 0
	}
	if a == b {
		return 0, 0
	}
	return 0, 0
}
*/

func (t Triangle) FindMaximumSumPath() int {
	for row := len(t) - 2; row >= 0; row-- {
		for column := 0; column <= row; column++ {
			t[row][column] += max(t[row+1][column], t[row+1][column+1])
		}
		// t.Display()
		t = t[:len(t)-1]
	}
	return t[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
