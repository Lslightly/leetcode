package main

/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	rows := make(map[int]bool)
	cols := make(map[int]bool)
	for rowIdx, row := range matrix {
		for colIdx, val := range row {
			if val == 0 {
				rows[rowIdx] = true
				cols[colIdx] = true
			}
		}
	}

	m := len(matrix)
	n := len(matrix[0])
	for rowIdx := range rows {
		for i := range n {
			matrix[rowIdx][i] = 0
		}
	}
	leftRows := make([]int, 0, m-len(rows))
	for i := range m {
		if _, ok := rows[i]; !ok {
			leftRows = append(leftRows, i)
		}
	}
	for colIdx := range cols {
		for _, rowIdx := range leftRows {
			matrix[rowIdx][colIdx] = 0
		}
	}
}

// @lc code=end
