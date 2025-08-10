package main

import "slices"

/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
func ge(a, b int) bool { return a >= b }

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix[0])
	col := slices.IndexFunc(matrix[0], func(elem int) bool {
		return ge(elem, target)
	})
	if col == -1 {
		// matrix[0][n-1] < target
		col = n - 1
	} else if matrix[0][col] == target {
		return true
	} else {
		// matrix[0][col] > target
		col--
	}
	curRow := 1
	for curRow < len(matrix) && col >= 0 {
		if matrix[curRow][col] == target {
			return true
		}
		if matrix[curRow][col] < target {
			for col < n && matrix[curRow][col] < target {
				col++
			}
			if col == n {
				// matrix[curRow][n-1] < target
				col = n - 1
				curRow++
				continue
			} else if matrix[curRow][col] == target {
				return true
			}
			// matrix[curRow][col] > target
			curRow++
			col--
			continue
		}
		for col >= 0 && matrix[curRow][col] > target {
			// col could not < 0
			col--
		}
		if col == -1 {
			return false
		}
		if matrix[curRow][col] == target {
			return true
		}
		// matrix[curRow][col] < target
		curRow++
	}
	return false
}

// @lc code=end
