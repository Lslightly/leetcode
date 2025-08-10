package main

/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

func hasCycleDestory(head *ListNode) bool {
	if head == nil {
		return false
	}
	cur := head
	for cur != nil {
		next := cur.Next
		if next == cur {
			return true
		}
		cur.Next = cur
		cur = next
	}
	return false
}

// @lc code=end
