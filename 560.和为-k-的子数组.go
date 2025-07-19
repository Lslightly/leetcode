package main

/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	cdfCnt := make(map[int]int)
	cdfCnt[0] = 1 // 初始值，处理从头开始的子数组
	cdf := 0
	count := 0
	for _, num := range nums {
		cdf += num
		count += cdfCnt[cdf-k]
		cdfCnt[cdf]++
	}
	return count
}

// @lc code=end
