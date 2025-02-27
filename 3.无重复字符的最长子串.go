package main

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start

type Set struct {
	m map[rune]bool
}

func NewSet() *Set {
	m := make(map[rune]bool)
	return &Set{m}
}

func (s *Set) Add(c rune) {
	s.m[c] = true
}

func (s *Set) Remove(c rune) {
	delete(s.m, c)
}

func (s *Set) Contains(c rune) bool {
	return s.m[c]
}

func (s *Set) Len() int {
	return len(s.m)
}

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	set := NewSet()
	beginPos := 0
	for _, ch := range s {
		startRemove := true
		for set.Contains(ch) {
			if startRemove {
				if set.Len() > maxLen {
					maxLen = set.Len()
				}
				startRemove = false
			}
			set.Remove(rune(s[beginPos]))
			beginPos++
		}
		set.Add(ch)
	}
	if set.Len() > maxLen {
		maxLen = set.Len()
	}
	return maxLen
}

// @lc code=end
