package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResizeMatrix(t *testing.T) {
	var m Matrix
	var expected, actual []int
	m.ResizeMatrix(2, 2)
	expected, actual = []int{2, 2}, []int{len(m), len(m[0])}
	assert.Equal(t, expected, actual)

	m.ResizeMatrix(2, 5)
	expected, actual = []int{2, 5}, []int{len(m), len(m[0])}
	assert.Equal(t, expected, actual)

	m.ResizeMatrix(5, 2)
	expected, actual = []int{5, 2}, []int{len(m), len(m[0])}
	assert.Equal(t, expected, actual)
}

func TestAssignValues(t *testing.T) {
	var expected, m Matrix
	m.ResizeMatrix(1, 4)
	m.AssignValues([]int{4, 8, 15, 16, 23, 42})
	expected = Matrix{{4, 8, 15, 16}}
	assert.Equal(t, expected, m)

	m.ResizeMatrix(1, 7)
	m.AssignValues([]int{4, 8, 15, 16, 23, 42})
	expected = Matrix{{4, 8, 15, 16, 23, 42, 0}}
	assert.Equal(t, expected, m)

	m.ResizeMatrix(4, 1)
	m.AssignValues([]int{4, 8, 15, 16, 23, 42})
	expected = make([][]int, 4);
	expected[0] = []int{4}; expected[1] = []int{8}; expected[2] = []int{15}; expected[3] = []int{16}
	assert.Equal(t, expected, m)
}

func TestIsIndexValid(t *testing.T) {
	var emptyMatrix Matrix
	assert.Equal(t, emptyMatrix.IsIndexValid(-1, -1), false)
	assert.Equal(t, emptyMatrix.IsIndexValid(0, 0), false)
	assert.Equal(t, emptyMatrix.IsIndexValid(1, 1), false)

	matrix := Matrix{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	assert.Equal(t, matrix.IsIndexValid(0, 0), true)
	assert.Equal(t, matrix.IsIndexValid(1, 2), true)
	assert.Equal(t, matrix.IsIndexValid(2, 1), true)
	
	assert.Equal(t, matrix.IsIndexValid(-1, -1), false)
	assert.Equal(t, matrix.IsIndexValid(-1, 2), false)
	assert.Equal(t, matrix.IsIndexValid(2, -1), false)
	assert.Equal(t, matrix.IsIndexValid(2, 3), false)
	assert.Equal(t, matrix.IsIndexValid(3, 2), false)
	assert.Equal(t, matrix.IsIndexValid(8, 8), false)
}

func TestIsReturnProductOfAdjacentNumbers(t *testing.T) {
	var m Matrix
	m.ResizeMatrix(3, 3)
	m.AssignValues([]int{4, 8, 15, 16, 23, 42, 2, 0, 1})

	// First row, all 3 numbers
	result, err := m.ReturnProductOfAdjacentNumbers("row", 3, 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, 480, result)
	// First row, last 2 numbers
	result, err = m.ReturnProductOfAdjacentNumbers("row", 3, 0, 1)
	assert.Nil(t, err)
	assert.Equal(t, 120, result)
	// First row, last number
	result, err = m.ReturnProductOfAdjacentNumbers("row", 3, 0, 2)
	assert.Nil(t, err)
	assert.Equal(t, 15, result)
	// Last two rows
	result, err = m.ReturnProductOfAdjacentNumbers("row", 3, 1, 0)
	assert.Nil(t, err)
	assert.Equal(t, 15456, result)
	result, err = m.ReturnProductOfAdjacentNumbers("row", 3, 2, 0)
	assert.Nil(t, err)
	assert.Equal(t, 0, result)

	// First column, all 3 numbers
	result, err = m.ReturnProductOfAdjacentNumbers("col", 3, 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, 128, result)
	// First column, last 2 numbers
	result, err = m.ReturnProductOfAdjacentNumbers("col", 3, 1, 0)
	assert.Nil(t, err)
	assert.Equal(t, 32, result)
	// First column, last 2 numbers
	result, err = m.ReturnProductOfAdjacentNumbers("col", 3, 2, 0)
	assert.Nil(t, err)
	assert.Equal(t, 2, result)
	// Last two columns
	result, err = m.ReturnProductOfAdjacentNumbers("col", 3, 0, 1)
	assert.Nil(t, err)
	assert.Equal(t, 0, result)
	result, err = m.ReturnProductOfAdjacentNumbers("col", 3, 0, 2)
	assert.Nil(t, err)
	assert.Equal(t, 630, result)

	// First row, invalid column
	result, err = m.ReturnProductOfAdjacentNumbers("row", 3, 0, 3)
	assert.Equal(t, 0, result)
	assert.EqualError(t, err, "invalid starting index: [0][3]")
	// First column, invalid row
	result, err = m.ReturnProductOfAdjacentNumbers("col", 3, 3, 0)
	assert.Equal(t, 0, result)
	assert.EqualError(t, err, "invalid starting index: [3][0]")
	// Invalid adjNums value
	result, err = m.ReturnProductOfAdjacentNumbers("col", -1, 0, 2)
	assert.Equal(t, 0, result)
	assert.EqualError(t, err, "invalid adjNums value: -1")
	// Invalid rowOrCol values
	result, err = m.ReturnProductOfAdjacentNumbers("invalid_value", 3, 0, 2)
	assert.Equal(t, 0, result)
	assert.EqualError(t, err, "invalid 'rowOrCol' value: invalid_value")
}

func TestIsReturnProductOfAdjacentNumbersDiag(t *testing.T) {
	var m Matrix
	m.ResizeMatrix(4, 4)
	m.AssignValues([]int{4, 8, 15, 16, 23, 42, 2, 6, 1, 0, 7, 32, 57, 99, 25, 5})

	// Left diagonals
	result, err := m.ReturnProductOfAdjacentNumbersDiag("anti", 2, 0, 3)
	assert.Nil(t, err)
	assert.Equal(t, 32, result)
	result, err = m.ReturnProductOfAdjacentNumbersDiag("anti", 3, 0, 3)
	assert.Nil(t, err)
	assert.Equal(t, 0, result)
	result, err = m.ReturnProductOfAdjacentNumbersDiag("anti", 3, 0, 2)
	assert.Nil(t, err)
	assert.Equal(t, 630, result)
	result, err = m.ReturnProductOfAdjacentNumbersDiag("anti", 2, 2, 3)
	assert.Nil(t, err)
	assert.Equal(t, 800, result)
	result, err = m.ReturnProductOfAdjacentNumbersDiag("anti", 1, 1, 3)
	assert.Nil(t, err)
	assert.Equal(t, 6, result)
	result, err = m.ReturnProductOfAdjacentNumbersDiag("anti", 3, 1, 3)
	assert.Nil(t, err)
	assert.Equal(t, 4158, result)
	result, err = m.ReturnProductOfAdjacentNumbersDiag("anti", 4, 0, 3)
	assert.Nil(t, err)
	assert.Equal(t, 0, result)

	// Right diagonals
	result, err = m.ReturnProductOfAdjacentNumbersDiag("main", 4, 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, 5880, result)
	result, err = m.ReturnProductOfAdjacentNumbersDiag("main", 4, 1, 1)
	assert.Nil(t, err)
	assert.Equal(t, 1470, result)
	result, err = m.ReturnProductOfAdjacentNumbersDiag("main", 1, 1, 0)
	assert.Nil(t, err)
	assert.Equal(t, 23, result)
	result, err = m.ReturnProductOfAdjacentNumbersDiag("main", 3, 1, 0)
	assert.Nil(t, err)
	assert.Equal(t, 0, result)

	// Edge cases
	result, err = m.ReturnProductOfAdjacentNumbersDiag("anti", 99, 2, 3)
	assert.Nil(t, err)
	assert.Equal(t, 800, result)
	result, err = m.ReturnProductOfAdjacentNumbersDiag("anti", 2, -2, -3)
	assert.Equal(t, 0, result)
	assert.EqualError(t, err, "invalid starting index: [-2][-3]")
	// Invalid rowOrCol values
	result, err = m.ReturnProductOfAdjacentNumbersDiag("invalid_value", 3, 0, 2)
	assert.Equal(t, 0, result)
	assert.EqualError(t, err, "invalid 'mainOrAnti' value: invalid_value")
	// Invalid adjNums value
	result, err = m.ReturnProductOfAdjacentNumbersDiag("anti", -1, 0, 2)
	assert.Equal(t, 0, result)
	assert.EqualError(t, err, "invalid adjNums value: -1")
}