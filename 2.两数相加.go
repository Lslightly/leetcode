package main

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
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
func listNodeLen(l1 *ListNode) int {
	length := 0
	for l1 != nil {
		length++
		l1 = l1.Next
	}
	return length
}
func newNode(l1 int, l2 int, carry int) (node *ListNode, newCarray int) {
	node = new(ListNode)
	sum := l1 + l2 + carry
	node.Val = sum % 10
	return node, sum / 10
}
func connPrev(prev, cur *ListNode) *ListNode {
	prev.Next = cur
	prev = cur
	return prev
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head, carry := newNode(l1.Val, l2.Val, 0)
	prev := head
	l1 = l1.Next
	l2 = l2.Next
	for l1 != nil && l2 != nil {
		cur, newCarry := newNode(l1.Val, l2.Val, carry)
		carry = newCarry
		prev = connPrev(prev, cur)
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		cur, newCarry := newNode(l1.Val, 0, carry)
		carry = newCarry
		prev = connPrev(prev, cur)
		l1 = l1.Next
		if carry == 0 {
			cur.Next = l1
			break
		}
	}
	for l2 != nil {
		cur, newCarry := newNode(l2.Val, 0, carry)
		carry = newCarry
		prev = connPrev(prev, cur)
		l2 = l2.Next
		if carry == 0 {
			cur.Next = l2
			break
		}
	}
	if carry != 0 {
		prev.Next = new(ListNode)
		prev.Next.Val = carry
	}
	return head
}

func newList(src []int) (head *ListNode) {
	head = new(ListNode)
	head.Val = src[0]
	prev := head
	for i := 1; i < len(src); i++ {
		prev.Next = new(ListNode)
		prev.Next.Val = src[i]
		prev = prev.Next
	}
	return head
}

// @lc code=end
