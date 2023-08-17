package entities

import (
	"errors"
	"fmt"
)

type rowOrCol string

const (
	row rowOrCol = "row"
	col rowOrCol = "col"
)

type mainOrAnti string

const (
	anti  mainOrAnti = "anti"
	main mainOrAnti = "main"
)

type Matrix [][]int

// Returns a "rows x cols" matrix filled with zeroes.
func (m *Matrix) ResizeMatrix(rows, cols int) {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}
	*m = matrix
}

// Assigns values to the zero-filled matrix.
func (m *Matrix) AssignValues(data []int) {
	currNum, dataLen := 0, len(data)
	rows, cols := len(*m), len((*m)[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if currNum == dataLen {
				break // 10x10 matrix requires 100 numbers, if there are less, loop stops to avoid error
			}
			(*m)[r][c] = data[currNum]
			currNum++
		}
	}
}

// Shows if given row, col coordinates exist in Matrix m
func (m Matrix) IsIndexValid(row, col int) bool {
	if (len(m)) == 0 {
		return false
	}

	numRows, numCols := len(m), len((m)[0])

	isRowValid := row >= 0 && row < numRows
	isColValid := col >= 0 && col < numCols

	if !isRowValid || !isColValid {
		return false
	}
	return true
}

func (m Matrix) Print() {
	for _, row := range m {
		for _, num := range row {
			fmt.Printf("%02d ", num)
		}
		fmt.Println()
	}
}

// Multiplies 'adjNums' numbers in the same row or column, starting from index [rowStart][colStart]
// For more info on the theory see https://en.wikipedia.org/wiki/Main_diagonal
func (m *Matrix) ReturnProductOfAdjacentNumbers(rowOrCol rowOrCol, adjNums, rowStart, colStart int) (int, error) {
	result := 1
	isIterationFinished := false
	currRow, currCol := rowStart, colStart
	isCurrentIndexValid := m.IsIndexValid(rowStart, colStart)
	if !isCurrentIndexValid {
		return 0, errors.New(fmt.Sprintf("invalid starting index: [%v][%v]", rowStart, colStart))
	}
	if adjNums < 0 {
		return 0, errors.New(fmt.Sprintf("invalid adjNums value: %v", adjNums))
	}

	switch rowOrCol {
	case row:
		for {
			isIterationFinished = currCol > colStart+(adjNums-1)
			isCurrentIndexValid = m.IsIndexValid(rowStart, currCol)
			if !isCurrentIndexValid || isIterationFinished {
				break
			}
			result *= (*m)[rowStart][currCol]
			currCol++
		}
	case col:
		for {
			isIterationFinished = currRow > rowStart+(adjNums-1)
			isCurrentIndexValid = m.IsIndexValid(currRow, colStart)
			if !isCurrentIndexValid || isIterationFinished {
				break
			}
			result *= (*m)[currRow][colStart]
			currRow++
		}
	default:
		return 0, errors.New(fmt.Sprintf("invalid 'rowOrCol' value: %v", rowOrCol))
	}
	return result, nil
}

// Multiplies 'adjNums' numbers in the diagonal (main or secondary), starting from index [rowStart][colStart]
// For more info on the theory see https://en.wikipedia.org/wiki/Main_diagonal
func (m *Matrix) ReturnProductOfAdjacentNumbersDiag(mainOrAnti mainOrAnti, adjNums, rowStart, colStart int) (int, error) {
	result := 1
	isIterationFinished := false
	currRow, currCol := rowStart, colStart
	isCurrentIndexValid := m.IsIndexValid(rowStart, colStart)
	if !isCurrentIndexValid {
		return 0, errors.New(fmt.Sprintf("invalid starting index: [%v][%v]", rowStart, colStart))
	}
	if adjNums < 0 {
		return 0, errors.New(fmt.Sprintf("invalid adjNums value: %v", adjNums))
	}
	switch mainOrAnti {
	case anti:
		for {
			isIterationFinished = currCol == colStart-adjNums && currRow == rowStart+adjNums
			isCurrentIndexValid = m.IsIndexValid(currRow, currCol)
			if !isCurrentIndexValid || isIterationFinished {
				break
			}
			result *= (*m)[currRow][currCol]
			currCol--
			currRow++
		}
	case main:
		for {
			isIterationFinished = currCol == colStart+adjNums && currRow == rowStart+adjNums
			isCurrentIndexValid = m.IsIndexValid(currRow, currCol)
			if !isCurrentIndexValid || isIterationFinished {
				break
			}
			result *= (*m)[currRow][currCol]
			currRow++
			currCol++
		}
	default:
		return 0, errors.New(fmt.Sprintf("invalid 'mainOrAnti' value: %v", mainOrAnti))
	}
	return result, nil
}
