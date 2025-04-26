package main

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	heights := []int{1, 2}
	maxArea(heights)
}

func TestMaxArea2(t *testing.T) {
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(heights))
}

func TestMaxArea3(t *testing.T) {
	heights := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(maxArea(heights))
}
