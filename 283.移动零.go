package main

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	zeroCnt := 0
	for i, num := range nums {
		if num == 0 {
			zeroCnt++
			continue
		}
		nums[i-zeroCnt] = num
	}
	for i := len(nums) - zeroCnt; i < len(nums); i++ {
		nums[i] = 0
	}
}

// @lc code=end
