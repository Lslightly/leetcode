package main

/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
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

func recoverList(headA *ListNode) {
	for headA != nil {
		headA.Val = -headA.Val
		headA = headA.Next
	}
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	oldA := headA
	for headA != nil {
		headA.Val = -headA.Val
		headA = headA.Next
	}
	defer recoverList(oldA)
	for headB != nil {
		if headB.Val < 0 {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// @lc code=end
