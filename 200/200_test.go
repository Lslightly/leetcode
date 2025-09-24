package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	expected := 1
	assert.Equal(t, expected, numIslands(grid))
}

func TestCase2(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	expected := 3
	assert.Equal(t, expected, numIslands(grid))
}

func TestCase3(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '0', '1', '1'},
		{'1', '1', '1', '1', '0'},
		{'0', '0', '0', '0', '0'},
		{'1', '0', '1', '1', '1'},
	}
	expected := 3
	assert.Equal(t, expected, numIslands(grid))
}
