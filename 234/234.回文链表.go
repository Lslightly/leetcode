package main

/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
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

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	n := listLen(head)
	if n == 2 {
		return head.Val == head.Next.Val
	}
	mid := n / 2
	var rightListHead *ListNode
	if n%2 == 0 {
		/*
			1 2 2 1
			  |
			  cur
			1 2 nil
			nil 2 1
		*/
		cur := head
		for range mid - 1 {
			cur = cur.Next
		}
		next := cur.Next
		cur.Next = nil
		rightListHead = reverseList(next, nil)
	} else {
		/*
			1 2 1
			  |
			  cur
			1 2 nil
			nil 2 1
		*/
		cur := head
		for range mid {
			cur = cur.Next
		}
		rightListHead = reverseList(cur, nil)
	}
	leftCur := head
	rightCur := rightListHead
	for leftCur != nil && rightCur != nil {
		if leftCur.Val != rightCur.Val {
			return false
		}
		leftCur = leftCur.Next
		rightCur = rightCur.Next
	}
	if leftCur != nil || rightCur != nil {
		return false
	}
	return true
}

func listLen(head *ListNode) int {
	n := 0
	for head != nil {
		n++
		head = head.Next
	}
	return n
}

func reverseList(head, headPrev *ListNode) *ListNode {
	prev := headPrev
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

// @lc code=end
