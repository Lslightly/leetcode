package main

/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start

func rotate(matrix [][]int) {
	transpose(matrix)
	reverse(matrix)
}

func transpose(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func reverse(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j, k := 0, len(matrix[i])-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}

// @lc code=end
