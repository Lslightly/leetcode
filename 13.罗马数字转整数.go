package main

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	var roman2int map[string]int = map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	var oneSet map[string]bool = map[string]bool{
		"I": true,
		"V": true,
		"X": true,
		"C": true,
	}

	res := 0
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		ch := s[i : i+1]
		if _, ok := oneSet[ch]; ok {
			if i+1 < sLen {
				if value, ok := roman2int[s[i:i+2]]; ok {
					res += value
					i++ // Skip the next character as it has been processed
					continue
				}
			}
		}
		res += roman2int[ch]
	}
	return res
}

// @lc code=end
