package main

import "math/rand"

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	worklist := make([]string, 0, len(strs)*2)
	for i := range rand.Perm(len(strs)) {
		worklist = append(worklist, strs[i])
	}
	maxCommonLen := 200
	for len(worklist) > 1 {
		s1, s2 := worklist[0], worklist[1]
		worklist[0], worklist[1] = "", ""
		worklist = worklist[2:]
		common := commonOfTwo(s1, s2, maxCommonLen)
		if common == "" {
			return common
		}
		worklist = append(worklist, common)
		if len(common) < maxCommonLen {
			maxCommonLen = len(common)
		}
	}
	return worklist[0]
}

func commonOfTwo(s1, s2 string, maxLen int) string {
	s1Len, s2Len := len(s1), len(s2)
	minLen := min(s1Len, s2Len, maxLen)
	for i := range minLen {
		if s1[i] != s2[i] {
			return s1[:i]
		}
	}
	return s1[:minLen]
}

// @lc code=end
