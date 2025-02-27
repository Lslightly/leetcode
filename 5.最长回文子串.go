package main

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start

func getMaxIfExist(l int, s string, preds []bool) (string, bool) {
	sLen := len(s)
	for i := 0; i+l <= sLen; i++ {
		if preds[i] {
			return s[i : i+l], true
		}
	}
	return "", false
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	sLen := len(s)
	dp := make([][]bool, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]bool, sLen)
	}
	for i := 0; i < sLen; i++ {
		dp[0][i] = true
		dp[1][i] = true
	}
	maxStr := string(s[0])
	for l := 2; l <= sLen; l++ {
		for i := 0; i+l-1 < sLen; i++ {
			dp[l%2][i] = (s[i] == s[i+l-1]) && dp[(l-2)%2][i+1]
		}
		tmpStr, exist := getMaxIfExist(l, s, dp[l%2])
		if exist {
			maxStr = tmpStr
		}
	}
	return maxStr
}

// func main() {
// 	longestPalindrome("babad")
// }

// @lc code=end
