package main

/*
 * @lc app=leetcode.cn id=1080 lang=golang
 *
 * [1080] 根到叶路径上的不足节点
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
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return root
	}
	revReduceOrd, deleteSet := calcPathSum(root, limit)
	for i := len(revReduceOrd) - 1; i >= 0; i-- {
		node := revReduceOrd[i]
		exist := false
		if node.Left != nil {
			if deleteSet.Contains(node.Left) {
				delete(deleteSet, node.Left)
				node.Left = nil
			} else {
				exist = true
			}
		}
		if node.Right != nil {
			if deleteSet.Contains(node.Right) {
				delete(deleteSet, node.Right)
				node.Right = nil
			} else {
				exist = true
			}
		}
		if !exist {
			deleteSet.Add(node)
		}
	}
	if deleteSet.Contains(root) {
		return nil
	}
	return root
}

type PathSumItem struct {
	node      *TreeNode
	prefixSum int
}

type NodeSet map[*TreeNode]struct{}

func (ns NodeSet) Contains(n *TreeNode) bool {
	_, ok := ns[n]
	return ok
}

func (ns NodeSet) Add(n *TreeNode) {
	ns[n] = struct{}{}
}

/*
calcPathSum calculates path sum for each leaf node, and returns
reduceRevOrder recording order of nodes in the reversed order of reduce
*/
func calcPathSum(root *TreeNode, limit int) (revReduceOrd []*TreeNode, deleteSet NodeSet) {
	wl := []PathSumItem{
		{
			node:      root,
			prefixSum: 0,
		},
	}
	revReduceOrd = make([]*TreeNode, 0)
	deleteSet = make(NodeSet)
	for len(wl) > 0 {
		top := wl[len(wl)-1]
		wl = wl[:len(wl)-1]
		prefixSum := top.prefixSum + top.node.Val
		appendNode := func(n *TreeNode) {
			if n != nil {
				wl = append(wl, PathSumItem{
					node:      n,
					prefixSum: prefixSum,
				})
			}
		}
		appendNode(top.node.Left)
		appendNode(top.node.Right)
		if top.node.Left == nil && top.node.Right == nil {
			if prefixSum < limit {
				deleteSet.Add(top.node)
			}
		} else {
			revReduceOrd = append(revReduceOrd, top.node)
		}
	}
	return
}

// @lc code=end
