package main

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

func reverseInts(heights []int) func(yield func(int, int) bool) {
	return func(yield func(int, int) bool) {
		for i := len(heights) - 1; i >= 0; i-- {
			if !yield(i, heights[i]) {
				return
			}
		}
	}
}

// @lc code=start
func trap(heights []int) int {
	leftMaxIdxChain := make([]int, len(heights))
	rightMaxIdxChain := make([]int, len(heights))
	leftMaxH := 0
	leftMaxHIdx := 0
	leftMaxIdxChain[0] = 0
	for i, h := range heights {
		if h >= leftMaxH {
			leftMaxIdxChain[i] = leftMaxHIdx
			if h == leftMaxH {
				// fast path, don't waste time in same height
				if heights[leftMaxIdxChain[leftMaxHIdx]] == h {
					leftMaxIdxChain[i] = leftMaxIdxChain[leftMaxHIdx]
				}
			}
			leftMaxHIdx = i
			leftMaxH = h
		}
	}
	rightMaxH := 0
	rightMaxHIdx := len(heights) - 1
	for i := len(heights) - 1; i >= leftMaxHIdx; i-- {
		h := heights[i]
		if h >= rightMaxH {
			rightMaxIdxChain[i] = rightMaxHIdx
			if h == rightMaxH {
				// fast path, don't waste time in same height
				if heights[rightMaxIdxChain[rightMaxHIdx]] == h {
					rightMaxIdxChain[i] = rightMaxIdxChain[rightMaxHIdx]
				}
			}
			rightMaxHIdx = i
			rightMaxH = h
		}
	}
	idx := leftMaxHIdx
	cover := 0
	for idx != leftMaxIdxChain[idx] {
		cover += heights[leftMaxIdxChain[idx]] * (idx - leftMaxIdxChain[idx])
		idx = leftMaxIdxChain[idx]
	}
	idx = leftMaxHIdx
	for idx != rightMaxIdxChain[idx] {
		cover += heights[rightMaxIdxChain[idx]] * (rightMaxIdxChain[idx] - idx)
		idx = rightMaxIdxChain[idx]
	}
	cover += heights[leftMaxHIdx]
	sum := 0
	for _, h := range heights {
		sum += h
	}
	return cover - sum
}

// @lc code=end
