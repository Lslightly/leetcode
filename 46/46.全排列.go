package main

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	lenNums := len(nums)
	unvisited := make(map[int]bool)
	for i := range lenNums {
		unvisited[i] = true
	}
	arrange := make([]int, lenNums)
	idxss := genIdxss(0, lenNums, arrange, unvisited)
	res := idxss
	for k, idxs := range res {
		for i, idx := range idxs {
			res[k][i] = nums[idx]
		}
	}
	return res
}

func factorial(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

func genIdxss(cur int, lenNums int, arrange []int, unvisited map[int]bool) [][]int {
	if cur == lenNums-1 {
		for v := range unvisited {
			newArrange := make([]int, lenNums)
			copy(newArrange, arrange)
			newArrange[lenNums-1] = v
			return [][]int{newArrange}
		}
		return nil
	}
	res := make([][]int, 0, factorial(lenNums-cur))
	keys := make([]int, 0, len(unvisited))
	for v := range unvisited {
		keys = append(keys, v)
	}
	for _, v := range keys {
		delete(unvisited, v)
		arrange[cur] = v
		res = append(res, genIdxss(cur+1, lenNums, arrange, unvisited)...)
		unvisited[v] = true
	}
	return res
}

// @lc code=end
