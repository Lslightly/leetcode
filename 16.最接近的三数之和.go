package main

import "sort"

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	fi, se, th := 0, 0, 0
	minDiff := 1<<31 - 1
	res := 0
	for ; fi < len(nums)-2; fi++ {
		se = fi + 1
		th = len(nums) - 1
		for se < th {
			diff, sum := dist(nums, fi, se, th, target)
			if diff < minDiff {
				minDiff = diff
				res = sum
			}
			if diff == 0 {
				return target
			}
			if nums[fi]+nums[se]+nums[th] < target {
				se++
			} else if nums[fi]+nums[se]+nums[th] > target {
				th--
			}
		}
	}
	return res
}

func dist(nums []int, i, j, k, target int) (int, int) {
	sum := nums[i] + nums[j] + nums[k]
	diff := sum - target
	if diff < 0 {
		return -diff, sum
	}
	return diff, sum
}

// @lc code=end
