package main

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	x32 := int32(x)
	if x32 == 0 {
		return 0
	}
	inputPos := x32 > 0
	res := newInt(x32)
	if res == 0 {
		return int(res)
	}
	resPos := res > 0
	if (resPos && !inputPos) || (!resPos && inputPos) {
		return 0
	}
	return int(res)
}

func yieldDigits(x int32) func(func(int32) bool) { // generator for digits from back to front
	return func(yield func(int32) bool) {
		for x != 0 {
			if !yield(x % 10) {
				return
			}
			x /= 10
		}
	}
}

func newInt(x int32) int32 {
	res := int32(0)
	for digit := range yieldDigits(x) {
		old_res := res
		res = 10*res + digit
		if (res-digit)/10 != old_res {
			return 0 // overflow
		}
	}
	return res
}

// func main() {
// 	reverse(123)
// }

// @lc code=end
