package main

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	anm1 := 2
	anm2 := 1
	for i := 3; i < n; i++ {
		anm1, anm2 = anm1+anm2, anm1
	}
	return anm1 + anm2
}

// @lc code=end
