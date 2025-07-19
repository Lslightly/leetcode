package main

/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
type PatCounter [26]int

func (pc *PatCounter) add(c byte) {
	(*pc)[c-'a']++
}

func (pc *PatCounter) remove(c byte) {
	(*pc)[c-'a']--
}

func (pc *PatCounter) get(c byte) int {
	return (*pc)[c-'a']
}

func (pc PatCounter) isEqual(other PatCounter) bool {
	for k, v := range pc {
		if other[k] != v {
			return false
		}
	}
	return true
}

func createPatCounter(p string) PatCounter {
	res := PatCounter{}
	for i := range p {
		res.add(p[i])
	}
	return res
}

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(s) < len(p) {
		return res
	}
	batchSize := 50000
	if len(s) < batchSize {
		return batch(s, p, 0)
	}
	ch := make(chan []int, len(s)/batchSize)
	cnt := 0
	for i := 0; i < len(s); i += batchSize {
		cnt++
		end := min(i+batchSize+len(p)-1, len(s))
		go func(start, end int) {
			ch <- batch(s[start:end], p, start)
		}(i, end)
	}
loop:
	for {
		select {
		case tmp := <-ch:
			res = append(res, tmp...)
			cnt--
		default:
			if cnt == 0 {
				break loop
			}
		}
	}
	return res
}

func batch(s string, p string, start int) []int {
	pCnt := createPatCounter(p)
	res := make([]int, 0)
	if len(s) < len(p) {
		return res
	}
	curCnt := createPatCounter(s[:len(p)])
	if curCnt.isEqual(pCnt) {
		res = append(res, start)
	}
	pLen, sLen := len(p), len(s)
	for i := pLen; i < sLen; i++ {
		curCnt.add(s[i])
		curCnt.remove(s[i-pLen])
		if curCnt.isEqual(pCnt) {
			res = append(res, i-pLen+1+start)
		}
	}
	return res
}

// @lc code=end
