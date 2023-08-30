package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTriangleValidData(t *testing.T) {
	inputData := [][]int{
		{1},
		{23, 84},
		{12, 3, 99},
	}
	expected := &Triangle{
		{1},
		{23, 84},
		{12, 3, 99},
	}

	triangle, err := NewTriangle(inputData)
	assert.NoError(t, err)
	assert.NotNil(t, triangle)
	assert.Equal(t, expected, triangle)

	inputData = [][]int{
		{99},
		{0, 12},
		{0, 0, -1},
	}
	expected = &Triangle{
		{99},
		{0, 12},
		{0, 0, -1},
	}

	triangle, err = NewTriangle(inputData)
	assert.NoError(t, err)
	assert.NotNil(t, triangle)
	assert.Equal(t, expected, triangle)
}

func TestNewTriangleInvalidData(t *testing.T) {
	inputData := [][]int{
		{},
		{23, 84},
		{12, 3, 94, 99},
	}
	triangle, err := NewTriangle(inputData)
	assert.EqualError(t, err, "invalid data at row 0: [], triangle not created")
	assert.Nil(t, triangle)

	inputData = [][]int{
		{9},
		{23, 84},
		{12, 3, 94, 99},
	}
	triangle, err = NewTriangle(inputData)
	assert.EqualError(t, err, "invalid data at row 2: [12 3 94 99], triangle not created")
	assert.Nil(t, triangle)

	inputData = [][]int{
		{1},
		{23, 84},
		{12, 3},
	}
	triangle, err = NewTriangle(inputData)
	assert.EqualError(t, err, "invalid data at row 2: [12 3], triangle not created")
	assert.Nil(t, triangle)
}
