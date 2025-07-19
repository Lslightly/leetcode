package main

/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */
// @lc code=start

func maxSlidingWindow(nums []int, k int) []int {
	idxQ := make([]int, 0)
	res := make([]int, 0, len(nums)-k+1)
	for i := range k {
		for len(idxQ) != 0 && nums[idxQ[len(idxQ)-1]] <= nums[i] {
			idxQ = idxQ[:len(idxQ)-1]
		}
		idxQ = append(idxQ, i)
	}
	res = append(res, nums[idxQ[0]])
	for i := k; i < len(nums); i++ {
		num := nums[i]
		if idxQ[0] == i-k {
			idxQ = idxQ[1:]
		}
		for len(idxQ) != 0 && nums[idxQ[len(idxQ)-1]] <= num {
			idxQ = idxQ[:len(idxQ)-1]
		}
		idxQ = append(idxQ, i)
		res = append(res, nums[idxQ[0]])
	}
	return res
}

// @lc code=end
