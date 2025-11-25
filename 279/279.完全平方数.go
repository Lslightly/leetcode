package main

/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	if n <= 4 {
		return map[int]int{
			0: 0,
			1: 1,
			2: 2,
			3: 3,
			4: 1,
		}[n]
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	dp[4] = 1
	for curn := 5; curn <= n; curn++ {
		minCount := curn
		for j := 1; j*j <= curn; j++ {
			withj2Count := dp[curn-j*j] + 1
			if withj2Count < minCount {
				minCount = withj2Count
			}
		}
		dp[curn] = minCount
	}
	return dp[n]
}

// @lc code=end
