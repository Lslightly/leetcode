package main

import (
	"slices"
	"strings"
)

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start

type Res struct {
	id   int
	line string
}

// [start, end)
type Range struct {
	start, end int
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	if numRows < 10 && len(s) < 1000 {
		lines := make([]string, 0, numRows)
		for i := range numRows {
			lines = append(lines, getSubStr(s, numRows, i))
		}
		return strings.Join(lines, "")
	}

	numThreads := 3

	lines := make([]string, 0, numThreads)
	partition := make([]Range, 0, numThreads)
	for i := range numThreads {
		partition = append(partition, Range{
			start: i * numRows / numThreads,
			end:   (i + 1) * numRows / numThreads,
		})
	}
	partition[numThreads-1].end = numRows

	ch := make(chan Res, numThreads)
	defer close(ch)

	for i := range numThreads {
		go func(s string, id int, part Range) {
			res := ""
			for i := part.start; i < part.end; i++ {
				res += getSubStr(s, numRows, i)
			}
			ch <- Res{
				id:   id,
				line: res,
			}
		}(s, i, partition[i])
	}
	resSlice := make([]Res, 0, numThreads)
	for res := range ch {
		resSlice = append(resSlice, res)
		if len(resSlice) == numThreads {
			break
		}
	}
	slices.SortFunc(resSlice, func(a, b Res) int {
		if a.id < b.id {
			return -1
		} else if a.id > b.id {
			return 1
		}
		return 0
	})
	for _, res := range resSlice {
		lines = append(lines, res.line)
	}
	return strings.Join(lines, "")
}

func getSubStr(s string, numRows, row int) string {
	if row == 0 || row == numRows-1 {
		return writeOneIdx(s, numRows, row)
	}
	return writeTwoIdx(s, numRows, row, 2*(numRows-1)-row)
}

func writeTwoIdx(s string, numRows, start1, start2 int) string {
	sLen := len(s)
	stride := 2 * (numRows - 1)
	res := make([]byte, 0, len(s)/numRows)
	idx1 := start1
	idx2 := start2
	for {
		added := false
		if idx1 < sLen {
			res = append(res, s[idx1])
			idx1 += stride
			added = true
		}
		if idx2 < sLen {
			res = append(res, s[idx2])
			idx2 += stride
			added = true
		}
		if !added {
			break
		}
	}
	return string(res)
}

func writeOneIdx(s string, numRows int, start int) string {
	res := make([]byte, 0, len(s)/numRows)
	stride := 2 * (numRows - 1)
	sLen := len(s)
	idx := start
	for idx < sLen {
		res = append(res, s[idx])
		idx += stride
	}
	return string(res)
}

// @lc code=end
