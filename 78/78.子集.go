package main

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	arrange := make([]int, 0, len(nums))
	return traverse(nums, 0, len(nums), arrange)
}

func traverse(nums []int, cur int, lenNums int, arrange []int) [][]int {
	if cur == lenNums-1 {
		tmp := make([]int, len(arrange))
		copy(tmp, arrange)
		return [][]int{
			append(tmp, nums[cur]),
			tmp,
		}
	}
	res := make([][]int, 0, 1<<(lenNums-cur))
	res = append(res, traverse(nums, cur+1, lenNums, append(arrange, nums[cur]))...)
	res = append(res, traverse(nums, cur+1, lenNums, arrange)...)
	return res
}

// @lc code=end
