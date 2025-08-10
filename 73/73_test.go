package main

import "testing"

func testcase(t *testing.T, matrix [][]int, want [][]int) {
	setZeroes(matrix)
	if len(matrix) != len(want) || len(matrix[0]) != len(want[0]) {
		t.Fatalf("expected matrix of size %dx%d, got %dx%d", len(want), len(want[0]), len(matrix), len(matrix[0]))
	}
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != want[i][j] {
				t.Errorf("at (%d, %d): expected %d, got %d", i, j, want[i][j], matrix[i][j])
			}
		}
	}
}

func TestCase1(t *testing.T) {
	testcase(
		t,
		[][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		},
		[][]int{
			{1, 0, 1},
			{0, 0, 0},
			{1, 0, 1},
		},
	)
}
func TestCase2(t *testing.T) {
	testcase(
		t,
		[][]int{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		},
		[][]int{
			{0, 0, 0, 0},
			{0, 4, 5, 0},
			{0, 3, 1, 0},
		},
	)
}
