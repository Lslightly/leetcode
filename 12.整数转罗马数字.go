package main

import "strings"

/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
func intToRoman(num int) string {
	digits := make([]int, 0, 4)
	for range 4 {
		digits = append(digits, num%10)
		num /= 10
	}
	res := ""
	res += toRoman(digits[3], "M", "", "")
	res += toRoman(digits[2], "C", "D", "M")
	res += toRoman(digits[1], "X", "L", "C")
	res += toRoman(digits[0], "I", "V", "X")
	return res
}

func toRoman(digit int, one, five, ten string) string {
	if digit == 0 {
		return ""
	}
	if digit == 4 {
		return one + five
	}
	if digit == 9 {
		return one + ten
	}
	if digit < 4 {
		return strings.Repeat(one, digit)
	}
	return five + strings.Repeat(one, digit-5)
}

// @lc code=end
