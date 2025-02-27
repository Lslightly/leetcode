package main

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start

func findMedianSortedArrays(a []int, b []int) float64 {
	m := len(a)
	n := len(b)
	if (m+n)%2 == 0 { // even
		a, b = trimElems(a, b, (m+n)/2-1)
		left, a, b := selFirst(a, b)
		right, a, b := selFirst(a, b)
		return float64(left+right) / 2
	}
	// odd
	a, b = trimElems(a, b, (m+n)/2)
	mid, _, _ := selFirst(a, b)
	return float64(mid)
}

func selFirst(a, b []int) (res int, resA []int, resB []int) {
	if len(a) == 0 {
		return b[0], a, b[1:]
	} else if len(b) == 0 {
		return a[0], a[1:], b
	}
	if a[0] < b[0] {
		return a[0], a[1:], b
	}
	return b[0], a, b[1:]
}

type IdxPair struct {
	aIdx int
	bIdx int
}

func newIdxPair(aLen int, sum int) IdxPair {
	return IdxPair{
		aIdx: aLen - 1,
		bIdx: sum - aLen - 1,
	}
}

func trimElems(a, b []int, k int) (resA []int, resB []int) {
	if k == 0 {
		return a, b
	}
	if len(a) > len(b) {
		b, a = trimElems(b, a, k)
		return a, b
	}
	if k == 1 {
		_, a, b = selFirst(a, b)
		return a, b
	} else if k == 2 {
		_, a, b = selFirst(a, b)
		_, a, b = selFirst(a, b)
		return a, b
	}

	// len(a) <= len(b)
	mid := k / 2

	// compare the last elem of a with b[]
	aLen := len(a)
	if aLen < mid {
		pair := newIdxPair(aLen, k/2)
		if a[pair.aIdx] < b[pair.bIdx] {
			// trim all elems of a and front of b
			return a[pair.aIdx+1:], b[newIdxPair(aLen, k).bIdx+1:]
		}
		// remove mid elems of b
		b = b[mid:]
		return trimElems(a, b, k-mid)
	}

	// remove the smaller part
	if a[mid-1] < b[mid-1] {
		a = a[mid:]
		return trimElems(a, b, k-mid)
	}
	b = b[mid:]
	return trimElems(a, b, k-mid)
}

// @lc code=end
