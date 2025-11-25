package main

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	lis := make([]int, len(nums))
	lis[0] = 1
	for end := 1; end < len(nums); end++ {
		maxLen := 1
		for i := 0; i < end; i++ {
			if nums[i] < nums[end] && 1+lis[i] > maxLen {
				maxLen = 1 + lis[i]
			}
		}
		lis[end] = maxLen
	}
	maxLen := 0
	for _, l := range lis {
		if l > maxLen {
			maxLen = l
		}
	}
	return maxLen
}

// @lc code=end
