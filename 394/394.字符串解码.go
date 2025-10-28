package main

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start

type PrefixElem struct {
	repeatTimes int
	scannedStr  string
}

type PrefixStack []PrefixElem

func (ps *PrefixStack) push(times int) {
	*ps = append(*ps, PrefixElem{
		repeatTimes: times,
		scannedStr:  "",
	})
}

func (ps *PrefixStack) pop() {
	pslen := len(*ps)
	top := (*ps)[pslen-1]
	(*ps)[pslen-2].scannedStr += strings.Repeat(top.scannedStr, top.repeatTimes)
	*ps = (*ps)[:pslen-1]
}

func (ps *PrefixStack) updateTop(substr string) {
	pslen := len(*ps)
	(*ps)[pslen-1].scannedStr += substr
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func getRepeatTimes(startIdx int, s string) (int, int) {
	idx := startIdx + 1
	for isDigit(s[idx]) {
		idx++
	}

	// s[idx] is [
	times, _ := strconv.Atoi(s[startIdx:idx])
	return times, idx
}

func decodeString(s string) string {
	// init stack
	ps := make(PrefixStack, 0)
	ps.push(1)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if isDigit(ch) {
			// 2321321[
			var times int
			times, i = getRepeatTimes(i, s)
			// s[i] is ]
			ps.push(times)
		} else if ch == ']' {
			/*
				....]
					^
			*/
			ps.pop()
		} else {
			startIdx := i
			for i < len(s) && !isDigit(s[i]) && s[i] != ']' {
				i++
			}
			ps.updateTop(s[startIdx:i])
			i--
		}
	}
	return ps[0].scannedStr
}

// @lc code=end
