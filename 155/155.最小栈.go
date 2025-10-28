package main

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start

type MinStack struct {
	realStack []int
	minStack  []int
}

func Constructor() MinStack {
	return MinStack{
		realStack: make([]int, 0),
		minStack:  make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.realStack = append(this.realStack, val)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
		return
	}
	oldMin := this.minStack[len(this.minStack)-1]
	if val < oldMin {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, oldMin)
	}
}

func (this *MinStack) Pop() {
	this.realStack = this.realStack[:len(this.realStack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.realStack[len(this.realStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
