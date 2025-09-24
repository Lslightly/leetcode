package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	assert.Equal(t, 4, orangesRotting(grid))
}

func TestCase2(t *testing.T) {
	grid := [][]int{{0, 2}}
	assert.Equal(t, 0, orangesRotting(grid))
}

func TestCase3(t *testing.T) {
	grid := [][]int{{1}}
	assert.Equal(t, -1, orangesRotting(grid))
}

func TestCase4(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 1},
		{0, 1, 1, 1},
		{0, 1, 1, 1},
		{1, 1, 1, 2},
	}
	assert.Equal(t, 4, orangesRotting(grid))
}

func TestCase5(t *testing.T) {
	grid := [][]int{
		{1, 1, 0, 1},
		{0, 1, 1, 1},
		{0, 1, 1, 1},
		{1, 1, 1, 2},
	}
	assert.Equal(t, 6, orangesRotting(grid))
}

func TestCase6(t *testing.T) {
	grid := [][]int{
		{1, 1, 1},
		{0, 0, 1},
		{2, 1, 1},
	}
	assert.Equal(t, 6, orangesRotting(grid))
}
