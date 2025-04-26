package main

import "container/heap"

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	return maxLeftMaxRightMaxArea(height)
}

func naiveMaxArea(height []int) int {
	max := 0
	for i := range height {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * min(height[i], height[j])
			if area > max {
				max = area
			}
		}
	}
	return max
}

func maxLeftMaxArea(height []int) int {
	max := 0
	maxAreaMaxLeft := 0
	for i := range height {
		if height[i] <= maxAreaMaxLeft {
			continue
		}
		// this is slow, can by optimzied to O(log n)
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * min(height[i], height[j])
			if area > max {
				max = area
				maxAreaMaxLeft = height[i]
			}
		}
	}
	return max
}

// HSet contains heights and their index
// HSet is a max heap, the max height is at the top
type HSet struct {
	data            []heapItem
	index2dataIndex map[int]int
}

// heapItem represents a height and its index
type heapItem struct {
	height int
	index  int
}

// create a new HSet from heights
func newHSet(heights []int) HSet {
	hset := HSet{
		data:            make([]heapItem, len(heights)),
		index2dataIndex: make(map[int]int),
	}
	for i, h := range heights {
		hset.data[i] = heapItem{height: h, index: i}
		hset.index2dataIndex[i] = i // Store the index of the item in the heap
	}
	heap.Init(&hset)
	return hset
}

// remove the height at index idx
func (h *HSet) remove(idx int) {
	// 获取要删除的元素在堆中的索引
	dataIdx, exists := h.index2dataIndex[idx]
	if !exists {
		return // 如果索引不存在，直接返回
	}

	// 将该元素的高度设置为一个足够大的值（大于所有可能的高度）
	h.data[dataIdx].height = int(^uint(0) >> 1) // 设置为 int 的最大值

	// 修复堆结构
	heap.Fix(h, dataIdx)

	// 弹出堆顶元素（即刚刚设置为最大值的元素）
	heap.Pop(h)
}

// return maxHeight, maxIdx
func (h *HSet) max() (int, int) {
	if len(h.data) == 0 {
		return 0, 0 // Return invalid values if the heap is empty
	}
	maxH := h.data[0].height
	maxIdx := 0
	for _, hi := range h.data {
		if hi.height < maxH {
			return maxH, maxIdx
		}
		if hi.index > maxIdx {
			maxIdx = hi.index
		}
	}

	return maxH, maxIdx
}

// return the index of the height >= ht
func (heap *HSet) findLargerEqThanIdxs(ht int) []int {
	res := make([]int, 0)
	for _, item := range heap.data {
		if item.height >= ht {
			res = append(res, item.index) // 记录索引
		}
		if item.height < ht {
			break // 因为是最大堆，后面的元素都小于 ht
		}
	}
	return res
}

// Implement heap.Interface for HSet
func (h *HSet) Len() int           { return len(h.data) }
func (h *HSet) Less(i, j int) bool { return h.data[i].height > h.data[j].height } // Max heap
func (h *HSet) Swap(i, j int) {
	h.index2dataIndex[h.data[i].index] = j
	h.index2dataIndex[h.data[j].index] = i
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *HSet) Push(x any) {
	h.data = append(h.data, x.(heapItem))
	h.index2dataIndex[x.(heapItem).index] = len(h.data) - 1
}

func (h *HSet) Pop() any {
	item := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	delete(h.index2dataIndex, item.index) // 删除索引映射
	return item
}

// TODO failed
func maxLeftMaxRightMaxArea(heights []int) int {
	maxArea := 0
	maxAreaMaxLeft := 0
	for i, hi := range heights {
		if hi <= maxAreaMaxLeft {
			continue
		}
		for j := i + 1; j < len(heights); j++ {
			area := (j - i) * min(heights[i], heights[j])
			if area > maxArea {
				maxArea = area
				maxAreaMaxLeft = heights[i]
			}
		}
	}
	return maxArea
}

func calcArea(i, j, height int) int {
	return (j - i) * height
}

// @lc code=end
