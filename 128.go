package main

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */
// @lc code=start

func longestConsecutive1(nums []int) int {
	input := make(map[int]bool)
	ks := make(map[int]bool)
	for _, num := range nums {
		input[num] = true
	}
	for num := range input {
		if !input[num-1] {
			ks[num] = true
		}
	}
	longest := 0
	for k := range ks {
		l := 1
		for input[k+l-1] {
			l++
		}
		if l > longest {
			longest = l - 1
		}
	}
	return longest
}

// TODO: TLE
func longestConsecutive(nums []int) int {
	input := make(map[int]bool)
	ks := make(map[int]bool)
	for _, num := range nums {
		input[num] = true
	}
	for num := range input {
		if _, ok := input[num-1]; !ok { // only add num where num-1 does not exist in nums
			ks[num] = true
		}
	}
	maxPossibleLen := len(input)
	for l := 2; l <= maxPossibleLen; l++ {
		dels := make([]int, 0, len(ks))
		exist := false
		for k := range ks {
			if _, ok := input[k+l-1]; !ok {
				dels = append(dels, k)
			} else {
				exist = true
			}
		}
		if !exist { // all buckets[k+l-1] do not exist
			return l - 1
		}
		for _, k := range dels {
			delete(ks, k)
		}
	}
	return maxPossibleLen
}

// @lc code=end
