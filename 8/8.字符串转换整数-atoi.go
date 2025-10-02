package main

import (
	"fmt"
	"math"
	"strings"
)

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	if len(s) == 0 {
		return 0
	}
	s, pos := signOfStr(s)
	s = strings.TrimLeft(s, "0")
	return getNumber(s, pos)
}

// true: +, false: -
func signOfStr(s string) (string, bool) {
	if s[0] == '-' {
		return s[1:], false
	}
	if s[0] == '+' {
		return s[1:], true
	}
	return s, true
}

func byteToInt(b byte) (int, bool) {
	return int(b - '0'), b >= '0' && b <= '9'
}

func getNumber(s string, pos bool) int {
	var collect func(prev int32, i int32) (int32, bool)
	res := 0
	if pos {
		collect = func(prev, digit int32) (int32, bool) {
			res := 10*prev + digit
			if (res-digit > 0 && res < 0) || (res-digit)/10 != prev {
				return math.MaxInt32, true
			}
			return res, false
		}
	} else {
		collect = func(prev, digit int32) (int32, bool) {
			res := 10*prev - digit
			if (res+digit < 0 && res > 0) || (res+digit)/10 != prev {
				return math.MinInt32, true
			}
			return res, false
		}
	}
	for _, ch := range s {
		ch := byte(ch)
		digit, isInt := byteToInt(ch)
		if !isInt {
			return res
		}
		tmp, overflow := collect(int32(res), int32(digit))
		res = int(tmp)
		if overflow {
			return res
		}
	}
	return res
}

type pair struct {
	in  string
	out int
}

// @lc code=end

func main() {
	s := []pair{
		pair{
			in:  " -042,",
			out: -42,
		},
		pair{
			in:  "-91283472332",
			out: -2147483648,
		},
		pair{
			in:  "2147483648",
			out: 2147483647,
		},
	}
	for _, p := range s {
		if myAtoi(p.in) != p.out {
			fmt.Println(p)
		}
	}
}
