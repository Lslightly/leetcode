package main

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */
// @lc code=start

type Direction int

const (
	Right Direction = iota
	Down
	Left
	Up
)

type State struct {
	x, y                         int
	direction                    Direction
	LeftB, RightB, TopB, BottomB int
}

func (s *State) rd() {
	s.direction = Down
	s.TopB++
	s.y++
}

func (s *State) dl() {
	s.direction = Left
	s.RightB--
	s.x--
}

func (s *State) lu() {
	s.direction = Up
	s.BottomB--
	s.y--
}

func (s *State) ur() {
	s.direction = Right
	s.LeftB++
	s.x++
}

func initState(m [][]int) *State {
	return &State{
		x:         -1,
		y:         0,
		direction: Right,
		LeftB:     0,
		RightB:    len(m[0]) - 1,
		TopB:      0,
		BottomB:   len(m) - 1,
	}
}

func (s *State) next() bool {
	switch s.direction {
	case Right:
		if s.x == s.RightB {
			s.rd()
			if s.TopB > s.BottomB {
				return false
			}
		} else {
			s.x++
		}
	case Down:
		if s.y == s.BottomB {
			s.dl()
			if s.LeftB > s.RightB {
				return false
			}
		} else {
			s.y++
		}
	case Left:
		if s.x == s.LeftB {
			s.lu()
			if s.TopB > s.BottomB {
				return false
			}
		} else {
			s.x--
		}
	case Up:
		if s.y == s.TopB {
			s.ur()
			if s.LeftB > s.RightB {
				return false
			}
		} else {
			s.y--
		}
	default:
		return false
	}
	return true
}

func spiralOrder(matrix [][]int) []int {
	s := initState(matrix)
	res := make([]int, 0, len(matrix)*len(matrix[0]))
	for s.next() {
		res = append(res, matrix[s.y][s.x])
	}
	return res
}

// @lc code=end
