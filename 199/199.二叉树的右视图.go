package main

/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	LastLayer := []*TreeNode{root}
	for len(LastLayer) > 0 {
		res = append(res, LastLayer[0].Val)
		curLayer := []*TreeNode{}
		for _, node := range LastLayer {
			if node.Right != nil {
				curLayer = append(curLayer, node.Right)
			}
			if node.Left != nil {
				curLayer = append(curLayer, node.Left)
			}
		}
		LastLayer = curLayer
	}
	return res
}

// @lc code=end
