package main

import "testing"

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
		{3749, "MMMDCCXLIX"},
	}

	for _, tt := range tests {
		if got := intToRoman(tt.num); got != tt.want {
			t.Errorf("intToRoman(%d) = %s; want %s", tt.num, got, tt.want)
		}
	}
}
