package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConverter(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}, сonverter([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))

	assert.Equal(t, []int{1, 2, 3, 1, 4, 7, 7, 9, 8, 7, 7, 4, 5, 6, 9, 8}, сonverter([][]int{
		{1, 2, 3, 1},
		{4, 5, 6, 4},
		{7, 8, 9, 7},
		{7, 8, 9, 7},
	}))
}
