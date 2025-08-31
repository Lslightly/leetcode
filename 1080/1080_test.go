package main

import (
	"lc/assert"
	"testing"
)

/*
genRoot generates a binary tree from a list of integers.
0 as nil

Example:

			5
		   / \
		  4   8
		 /   / \
		11  13  4
		/ \      \
	   7  2       1

will be represented as:
[

	[5],
	[4,8],
	[11,0,13,4],
	[7,2,0,0,0,1]

]
*/
func genRoot(layers [][]int) *TreeNode {
	root := &TreeNode{
		Val: layers[0][0],
	}
	lastLayer := []*TreeNode{root}
	for _, layer := range layers[1:] {
		newLayer := make([]*TreeNode, 0)
		parentIdx := 0
		for i, num := range layer {
			isLeft := i%2 == 0
			var newNode *TreeNode = nil
			if num != 0 {
				newNode = &TreeNode{
					Val: num,
				}
			}
			if isLeft {
				lastLayer[parentIdx].Left = newNode
			} else {
				lastLayer[parentIdx].Right = newNode
				parentIdx++
			}
			if newNode != nil {
				newLayer = append(newLayer, newNode)
			}
		}
		lastLayer = newLayer
	}
	return root
}

func MatchTree(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a != nil && b == nil {
		return false
	}
	if a == nil && b != nil {
		return false
	}
	return MatchTree(a.Left, b.Left) && MatchTree(a.Right, b.Right) && a.Val == b.Val
}

func TestCase1(t *testing.T) {
	root := genRoot(
		[][]int{
			{1},
			{2, 3},
			{4, -99, -99, 7},
			{8, 9, -99, -99, 12, 13, -99, 14},
		},
	)
	expect := genRoot(
		[][]int{
			{1},
			{2, 3},
			{4, 0, 0, 7},
			{8, 9, 0, 14},
		},
	)
	assert.Sastisfy(t, MatchTree(sufficientSubset(root, 22), expect), "Case 1")
}

func TestCase2(t *testing.T) {
	root := genRoot(
		[][]int{
			{5},
			{4, 8},
			{11, 0, 17, 4},
			{7, 1, 0, 0, 5, 3},
		},
	)
	expect := genRoot(
		[][]int{
			{5},
			{4, 8},
			{11, 0, 17, 4},
			{7, 0, 0, 0, 5, 0},
		},
	)

	assert.Sastisfy(t, MatchTree(sufficientSubset(root, 22), expect), "Case 2")
}

func TestCase3(t *testing.T) {
	root := genRoot(
		[][]int{
			{1},
			{2, -3},
			{-5, 0, 4, 0},
		},
	)
	expect := genRoot(
		[][]int{
			{1},
			{0, -3},
			{4, 0},
		},
	)
	res := sufficientSubset(root, -1)
	assert.Sastisfy(t, MatchTree(res, expect), "Case 3")
}
