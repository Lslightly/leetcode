package main

import (
	"slices"
)

/*
 * @lc app=leetcode.cn id=2585 lang=golang
 *
 * [2585] 获得分数的方法数
 */

// @lc code=start
const mod = 1e9 + 7

type problem struct {
	count int
	mark  int
}

func sortedProblemsByMarkDesc(types [][]int) []problem {
	ps := make([]problem, len(types))
	for i, t := range types {
		ps[i] = problem{count: t[0], mark: t[1]}
	}
	slices.SortFunc(ps, func(a, b problem) int {
		return b.mark - a.mark
	})
	return ps
}

func skipProblemsGreaterThanTarget(ps []problem, target int) []problem {
	for i, p := range ps {
		// skip problem whose mark is greater than target
		if p.mark <= target {
			return ps[i:]
		}
	}
	return nil
}

const (
	unavailable int = -1
	unknown     int = 0
)

type combineCountCache [1001][50]int

func (cache *combineCountCache) get(left, idx int) (cnt int, valid bool) {
	num := cache[left][idx]
	return num, !(num == unavailable || num == unknown)
}

func recurPickProblem(ps []problem, left, idx int, cache *combineCountCache) int {
	if cnt, valid := cache.get(left, idx); valid {
		return cnt
	} else if cnt == unavailable {
		return 0
	}
	if idx == len(ps)-1 {
		if left%ps[idx].mark == 0 && left/ps[idx].mark <= ps[idx].count {
			cache[left][idx] = 1
			return 1
		}
		cache[left][idx] = unavailable
		return 0
	}
	methodCnt := 0
	for curNum := range min(left/ps[idx].mark+1, ps[idx].count+1) {
		methodCnt = (methodCnt + recurPickProblem(ps, left-curNum*ps[idx].mark, idx+1, cache)) % mod
	}
	if methodCnt == 0 {
		cache[left][idx] = unavailable
		return 0
	}
	cache[left][idx] = methodCnt
	return methodCnt
}

func waysToReachTarget(target int, types [][]int) int {
	ps := sortedProblemsByMarkDesc(types)
	ps = skipProblemsGreaterThanTarget(ps, target)
	if ps == nil {
		return 0
	}
	var combCnt combineCountCache
	return recurPickProblem(ps, target, 0, &combCnt)
}

// @lc code=end
