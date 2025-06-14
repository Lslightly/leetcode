package main

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		// {[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		// {[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		// {[]int{1, 1, -2}, [][]int{{1, 1, -2}}},
		{[]int{0, 1, 1}, [][]int{}},
	}

	for _, tt := range tests {
		threeSum(tt.nums)
	}
}
