package main

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dpim2 := nums[0]
	dpim1 := max(nums[1], nums[0])
	for i := 2; i < len(nums); i++ {
		dpi := max(dpim2+nums[i], dpim1)
		dpim2, dpim1 = dpim1, dpi
	}
	return dpim1
}

// @lc code=end
