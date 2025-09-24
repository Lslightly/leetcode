package main

/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isValid(root.Left, -1<<63, root.Val) && isValid(root.Right, root.Val, 1<<63-1)
}

func isValid(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}
	if node.Val <= lower || node.Val >= upper {
		return false
	}
	return isValid(node.Left, lower, node.Val) && isValid(node.Right, node.Val, upper)
}

// @lc code=end
