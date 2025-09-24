package main

import (
	"log"
)

/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start

const (
	NotVisited int = iota
	Visiting
	Finished
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	status := make([]int, numCourses)
	outEdges := make([][]int, numCourses)
	for _, p := range prerequisites {
		dst, src := p[0], p[1]
		outEdges[src] = append(outEdges[src], dst)
	}
	stack := make([]int, 0, numCourses)
	/*
		a,  b,	...
			1,
			2,
		a
	*/
	for src, edges := range outEdges {
		if len(edges) != 0 {
			stack = append(stack, src)
		}
	}
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		if status[top] == Visiting {
			status[top] = Finished
			stack = stack[:len(stack)-1]
			continue
		}
		status[top] = Visiting
		for _, dst := range outEdges[top] {
			switch status[dst] {
			case Visiting:
				// find loop
				return false
			case Finished:
				// won't be loop
			case NotVisited:
				stack = append(stack, dst)
			}
		}
	}
	return true
}

// @lc code=end

func main() {
	if !canFinish(2, [][]int{
		{1, 0},
	}) {
		log.Fatalln("error1")
	}
	if canFinish(2, [][]int{
		{1, 0},
		{0, 1},
	}) {
		log.Fatalln("error2")
	}
	if !canFinish(3, [][]int{
		{0, 1},
		{1, 2},
	}) {
		log.Fatalln("error3")
	}
	log.Println("Success")
}
