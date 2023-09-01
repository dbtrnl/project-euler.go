package problems

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/dbtrnl/project-euler.go/internal/entities"
	"github.com/dbtrnl/project-euler.go/pkg/utils"
)

func Problem67() int {
	var trData [][]int
	var trRow []int
	var answer int

	filePath := "./internal/input_data/problem67_input.txt"
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}
	// Assembling the triangle
	input := string(fileContents)
	input = strings.TrimSuffix(input, "\n")
	buffer := utils.NewDigitBuffer(2)

	for i, currRune := range input {
		if buffer.IsFull {
			trRow = append(trRow, buffer.Data)
			buffer.Empty()
		}
		if currRune == 10 {
			trData = append(trData, trRow)
			trRow = []int{}
			buffer.Empty()
			continue
		}
		if i == len(input)-1 {
			buffer.Push(currRune)
			trRow = append(trRow, buffer.Data)
			break
		}
		buffer.Push(currRune)
	}
	trData = append(trData, trRow)

	// Create the triangle and solve problem
	triangle, err := entities.NewTriangle(trData)
	if err != nil {
		panic(err)
	}
	answer = triangle.FindMaximumSumPath()
	return answer
}