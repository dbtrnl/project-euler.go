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
}
