package entities

type Matrix [][]int

// Returns a "rows x cols" matrix filled with zeroes.
func (m *Matrix) NewMatrix(rows, cols int) {
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
				break // 10x10 matrix requires 100 numbers, if there are less, loops stops to avoid "index out of range" err
			}
			(*m)[r][c] = data[currNum]
			currNum++
		}
	}
}

/* This method is not necessary, since NewMatrix assures it will be valid
func (m *Matrix) IsValid() bool {
	if len(*m) <= 0 || len((*m)[0]) <= 0 {
		return false
	}
	numCols := len((*m)[0])
	for _, row := range *m {
		if len(row) != numCols {
			return false
		}
	}
	return true
}
*/