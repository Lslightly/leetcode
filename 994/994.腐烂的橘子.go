package main

/*
 * @lc app=leetcode.cn id=994 lang=golang
 *
 * [994] 腐烂的橘子
 */

// @lc code=start

type CellStatus int

const (
	Blank CellStatus = iota
	Fresh
	Rotten
)

type point struct {
	i, j int
}

func orangesRotting(grid [][]int) int {
	freshCnt := freshCntOfGrid(grid)
	rottens := rottensOfGrid(grid)
	queue := rottens
	distances := make([][]int, len(grid))
	for i := range grid {
		distances[i] = make([]int, len(grid[0]))
	}
	for len(queue) != 0 {
		item := queue[0]
		queue = queue[1:]
		i, j := item.i, item.j
		propogate := func(tgti, tgtj int) {
			if grid[tgti][tgtj] == int(Fresh) {
				freshCnt--
				grid[tgti][tgtj] = int(Rotten)
				distances[tgti][tgtj] = distances[i][j] + 1
				queue = append(queue, point{tgti, tgtj})
			}
		}
		if i > 0 {
			propogate(i-1, j)
		}
		if i < len(grid)-1 {
			propogate(i+1, j)
		}
		if j > 0 {
			propogate(i, j-1)
		}
		if j < len(grid[0])-1 {
			propogate(i, j+1)
		}
	}
	if freshCnt != 0 {
		return -1
	}
	maxDist := 0
	for _, row := range distances {
		for _, elem := range row {
			if elem > maxDist {
				maxDist = elem
			}
		}
	}
	return maxDist
}

func freshCntOfGrid(grid [][]int) int {
	cnt := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == int(Fresh) {
				cnt++
			}
		}
	}
	return cnt
}

func rottensOfGrid(grid [][]int) []point {
	rottens := make([]point, 0)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == int(Rotten) {
				rottens = append(rottens, point{i, j})
			}
		}
	}
	return rottens
}

// @lc code=end
