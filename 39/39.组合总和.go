package main

import (
	"slices"
)

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start

type selectedUnit struct {
	unit int
	cnt  int
}

func combinationSum(candidates []int, target int) [][]int {
	slices.SortFunc(candidates, func(a int, b int) int {
		if a < b {
			return 1
		} else if a > b {
			return -1
		}
		return 0
	})
	selected := make([]selectedUnit, 0, len(candidates))
	return genCombination(candidates, selected, target, 0)
}

func genCombination(cands []int, selected []selectedUnit, left int, cur int) [][]int {
	curElem := cands[cur]
	if cur == len(cands)-1 {
		if left%curElem == 0 {
			res := make([]int, 0, len(selected))
			for _, sunit := range selected {
				for range sunit.cnt {
					res = append(res, sunit.unit)
				}
			}
			for range left / curElem {
				res = append(res, curElem)
			}
			return [][]int{res}
		}
		return nil
	}
	maxLimit := left / curElem
	res := make([][]int, 0)
	for tryCnt := range maxLimit + 1 {
		tmpRes := genCombination(
			cands,
			append(selected, selectedUnit{
				unit: curElem,
				cnt:  tryCnt,
			}),
			left-curElem*tryCnt,
			cur+1,
		)
		if tmpRes != nil {
			res = append(res, tmpRes...)
		}
	}
	return res
}

// @lc code=end
