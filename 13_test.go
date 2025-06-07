package main

import "testing"

func TestRomanToInt(t *testing.T) {
	if romanToInt("III") != 3 {
		t.Fail()
	}
	if romanToInt("LVIII") != 58 {
		t.Fail()
	}
	if romanToInt("MCMXCIV") != 1994 {
		t.Fail()
	}
}
