package main

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start

func longestValidParentheses(s string) int {
	sLen := len(s)
	if sLen <= 1 {
		return 0
	}
	dp := make([]int, sLen)
	dp[0] = 0
	if s[0] == '(' && s[1] == ')' {
		dp[1] = 2
	} else {
		dp[1] = 0
	}
	for i := 2; i < sLen; i++ {
		if s[i] == ')' && s[i-1] == '(' {
			dp[i] = dp[i-2] + 2
		} else if s[i] == ')' && s[i-1] == ')' {
			if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				dp[i] = dp[i-1] + 2
				prev := 0 // prev before ((...))
				if i-dp[i-1]-2 >= 0 {
					prev = dp[i-dp[i-1]-2]
				}
				dp[i] += prev
			}
		}
	}
	max := 0
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}

// @lc code=end
// )((

// func main() {
// 	fmt.Println(longestValidParentheses(")()())"))
// }
