package main

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	right := head
	for range n - 1 {
		right = right.Next
		if right == nil { // len(list) < n
			return head
		}
	}
	var delPrev *ListNode = nil
	del := head
	for right.Next != nil {
		right = right.Next
		delPrev = del
		del = del.Next
	}
	if delPrev == nil {
		return head.Next
	}
	delPrev.Next = del.Next
	return head
}

// @lc code=end
