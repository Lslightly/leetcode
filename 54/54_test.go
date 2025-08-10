package main

import "testing"

func testcase(t *testing.T, m [][]int, want []int) {
	got := spiralOrder(m)
	if len(got) != len(want) {
		t.Fatalf("expected length %d, got %d", len(want), len(got))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("at index %d: expected %d, got %d", i, want[i], got[i])
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
		[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
	)
}
func TestCase2(t *testing.T) {
	testcase(
		t,
		[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
		[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
	)
}

func TestCase3(t *testing.T) {
	testcase(
		t,
		[][]int{
			{1},
			{2},
			{3},
			{4},
			{5},
		},
		[]int{1, 2, 3, 4, 5},
	)
}
