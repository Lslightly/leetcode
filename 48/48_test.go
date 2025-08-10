package main

import "testing"

func testcase(t *testing.T, m [][]int, want []int) {
	rotate(m)
	for i := range m {
		for j := range m[i] {
			if m[i][j] != want[i*len(m)+j] {
				t.Errorf("rotate(%v) = %v, want %v", m, m, want)
			}
		}
	}
}

func TestCase1(t *testing.T) {
	testcase(
		t,
		[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		[]int{
			7, 4, 1,
			8, 5, 2,
			9, 6, 3,
		},
	)
}
func TestCase2(t *testing.T) {
	testcase(
		t,
		[][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		},
		[]int{
			15, 13, 2, 5,
			14, 3, 4, 1,
			12, 6, 8, 9,
			16, 7, 10, 11,
		},
	)
}
