package main

import (
	"slices"
)

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
type NumCounter struct {
	count   map[int]int // key -> cnt
	negKeys []int       // key < 0
	posKeys []int       // key > 0
	zeros   int         // num of 0
}

type TupleSet map[[3]int]bool

func tuple2slice(a [3]int) []int {
	return []int{a[0], a[1], a[2]}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func threeSum(nums []int) [][]int {
	nc := createCounter(nums)
	res := make([][]int, 0)
	if nc.case1() {
		res = append(res, []int{0, 0, 0})
	}
	if tuples := nc.case2(); len(tuples) > 0 {
		res = append(res, mergeTuples(tuples)...)
	}
	tuples := make([][3]int, 0)
	tuples = append(tuples, nc.case3()...)
	tuples = append(tuples, nc.case4()...)
	res = append(res, mergeTuples(tuples)...)
	return res
}

func createCounter(nums []int) NumCounter {
	nc := NumCounter{
		count:   make(map[int]int),
		negKeys: make([]int, 0, len(nums)/2),
		posKeys: make([]int, 0, len(nums)/2),
		zeros:   0,
	}
	for _, k := range nums {
		if k == 0 {
			nc.zeros++
			continue
		}
		if k > 0 {
			nc.posKeys = append(nc.posKeys, k)
			nc.count[k]++
		} else {
			nc.negKeys = append(nc.negKeys, k)
			nc.count[k]++
		}
	}

	slices.Sort(nc.posKeys)
	slices.Sort(nc.negKeys)
	slices.Reverse(nc.negKeys)
	return nc
}

/*
0 0 0
*/
func (nc *NumCounter) case1() bool {
	return nc.zeros >= 3
}

func (nc *NumCounter) case2() [][3]int {
	if nc.zeros == 0 {
		return nil
	}
	tuples := make([][3]int, 0)
	for _, neg := range nc.negKeys {
		if nc.count[abs(neg)] > 0 {
			tuples = append(tuples, [3]int{neg, 0, abs(neg)})
		}
	}
	return tuples
}

// -, +, +
func (nc *NumCounter) case3() [][3]int {
	if len(nc.negKeys) < 1 || len(nc.posKeys) < 2 {
		return nil
	}
	return oneMatchTwo(nc.negKeys, nc.posKeys, nc.count)
}

// -, -, +
func (nc *NumCounter) case4() [][3]int {
	if len(nc.negKeys) < 2 || len(nc.posKeys) < 1 {
		return nil
	}
	return oneMatchTwo(nc.posKeys, nc.negKeys, nc.count)
}

// oneArr[?] + twoArr[?] + twoArr[?] = 0
func oneMatchTwo(oneArr []int, twoArr []int, count map[int]int) [][3]int {
	res := make([][3]int, 0)
	for _, expect := range oneArr {
		if abs(twoArr[0]) >= abs(expect) {
			continue
		}
		for _, cur := range twoArr {
			last := -expect - cur
			if abs(cur) > abs(last) {
				break
			}
			if last == cur { // expect + a + a = 0
				if count[cur] >= 2 {
					res = append(res, createOrderedTuple(expect, cur, last))
				}
				break // no need for further iteration
			}
			if count[last] > 0 {
				res = append(res, createOrderedTuple(expect, cur, last))
			}
		}
	}
	return res
}

func createOrderedTuple(expect, cur, last int) [3]int {
	if expect < 0 {
		return [3]int{expect, cur, last}
	}
	return [3]int{last, cur, expect}
}

func mergeTuples(inTuples [][3]int) [][]int {
	tupleSet := make(TupleSet)
	for _, tuple := range inTuples {
		tupleSet[tuple] = true
	}
	res := make([][]int, 0, len(tupleSet))
	for tuple := range tupleSet {
		res = append(res, tuple2slice(tuple))
	}
	return res
}

// @lc code=end
