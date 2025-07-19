package main

import (
	"slices"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	s := "cbaebabacd"
	p := "abc"
	expected := []int{0, 6}
	result := findAnagrams(s, p)
	if !slices.Equal(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
