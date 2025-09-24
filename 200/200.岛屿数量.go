package main

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */
// @lc code=start

const (
	land   byte = '1'
	sea    byte = '0'
	marked byte = '2'
)

func numIslands(grid [][]byte) int {
	count := 0
	for i, row := range grid {
		for j, cell := range row {
			if cell == land {
				recursiveExploreAndMark(grid, i, j)
				count++
			}
		}
	}
	return count
}

func recursiveExploreAndMark(grid [][]byte, i, j int) {
	if !isValidLand(grid, i, j) {
		return
	}
	grid[i][j] = marked
	recursiveExploreAndMark(grid, i, j-1)
	recursiveExploreAndMark(grid, i, j+1)
	recursiveExploreAndMark(grid, i-1, j)
	recursiveExploreAndMark(grid, i+1, j)
}

type direction byte

const (
	left  direction = 0
	up    direction = 1
	right direction = 2
	down  direction = 3
)

type workItem struct {
	i, j int
}

func exploreAndMark(grid [][]byte, i, j int) {
	grid[i][j] = marked
	var wl []workItem
	wl = genNextWorkItems(grid, i, j, wl)
	for len(wl) > 0 {
		item := wl[len(wl)-1]
		wl = wl[:len(wl)-1]
		grid[item.i][item.j] = marked
		wl = genNextWorkItems(grid, item.i, item.j, wl)
	}
}

func genNextWorkItems(grid [][]byte, i, j int, wl []workItem) []workItem {
	if isValidLand(grid, i, j-1) {
		wl = append(wl, workItem{i, j - 1})
	}
	if isValidLand(grid, i-1, j) {
		wl = append(wl, workItem{i - 1, j})
	}
	if isValidLand(grid, i, j+1) {
		wl = append(wl, workItem{i, j + 1})
	}
	if isValidLand(grid, i+1, j) {
		wl = append(wl, workItem{i + 1, j})
	}
	return wl
}

func isValidLand(grid [][]byte, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] == land
}

// @lc code=end
