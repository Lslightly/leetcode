package main

import "testing"

func testcase(t *testing.T, matrix [][]int, target int, expected bool) {
	if searchMatrix(matrix, target) != expected {
		t.Errorf("searchMatrix(%v, %d) = %v; want %v", matrix, target, !expected, expected)
	}
}

func TestCase1(t *testing.T) {
	testcase(
		t,
		[][]int{
			{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		},
		5,
		true,
	)
}

func TestCaes2(t *testing.T) {
	testcase(
		t,
		[][]int{
			{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		},
		20,
		false,
	)
}

func TestCase3(t *testing.T) {
	testcase(
		t,
		[][]int{
			{1},
			{3},
			{5},
		},
		5,
		true,
	)
}
