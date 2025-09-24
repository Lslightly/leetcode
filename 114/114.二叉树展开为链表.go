package main

/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			cur = cur.Right
			continue
		}
		pre := cur.Left
		for pre.Right != nil {
			pre = pre.Right
		}
		pre.Right = cur.Right
		cur.Right = cur.Left
		cur.Left = nil
		cur = cur.Right
	}
}

// @lc code=end
