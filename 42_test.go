package main

import "testing"

func Test42_1(t *testing.T) {
	hs := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	if trap(hs) != 6 {
		t.Fail()
	}
}

func Test42_2(t *testing.T) {
	hs := []int{4, 2, 0, 3, 2, 5}
	if trap(hs) != 9 {
		t.Fail()
	}
}
