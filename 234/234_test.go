package main

import "testing"

func testcase(t *testing.T, list []int, want bool) {
	head := &ListNode{Val: list[0]}
	cur := head
	for _, v := range list[1:] {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	got := isPalindrome(head)
	if got != want {
		t.Errorf("isPalindrome(%v) = %v, want %v", list, got, want)
	}
}
func TestCase1(t *testing.T) {
	testcase(t, []int{1, 2, 2, 1}, true)
}
func TestCase2(t *testing.T) {
	testcase(t, []int{1, 2, 3, 4, 5}, false)
}

func TestCase3(t *testing.T) {
	testcase(t, []int{1, 2}, false)
}

func TestCase4(t *testing.T) {
	testcase(t, []int{1, 2, 1}, true)
}
