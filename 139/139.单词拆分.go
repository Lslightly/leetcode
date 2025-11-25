package main

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	canBreak := make([]bool, len(s))
	for _, word := range wordDict {
		if len(word) <= len(s) && s[:len(word)] == word {
			canBreak[len(word)-1] = true
		}
	}
outer:
	for end := 0; end < len(s); end++ {
		if canBreak[end] {
			continue
		}
		for _, word := range wordDict {
			if len(word) >= end+1 {
				continue
			}
			if canBreak[end-len(word)] && s[end-len(word)+1:end+1] == word {
				canBreak[end] = true
				continue outer
			}
		}
	}
	return canBreak[len(s)-1]
}

// @lc code=end
