package day10

import "aoc24/util"

const (
	_ = iota
	up
	right
	down
	left
)

type Pos struct {
	X int
	Y int
}

type Hiker struct {
	currPos     Pos
	currNum     int
	currDir     int
	visitedPos  []Pos
	visitedDirs []int
}

func (h *Hiker) visited(x int, y int) bool {
	nextPos := Pos{x, y}
	return contains(h.visitedPos, nextPos)
}

func (h *Hiker) visit(x int, y int) {
	h.currNum++
	h.currPos = Pos{x, y}
	h.visitedPos = append(h.visitedPos, h.currPos)
	if h.currDir == right {
		h.visitedDirs = []int{left}
	} else if h.currDir == down {
		h.visitedDirs = []int{up}
	} else if h.currDir == left {
		h.visitedDirs = []int{right}
	} else {
		h.visitedDirs = []int{down}
	}
}

func (h *Hiker) changeDirection() {
	if h.currDir == right {
		h.currDir = down
	} else if h.currDir == down {
		h.currDir = left
	} else if h.currDir == left {
		h.currDir = up
	} else {
		h.currDir = right
	}
	h.visitedDirs = append(h.visitedDirs, h.currDir)
}

func (h *Hiker) allDirectionsChecked() bool {
	if util.ContainsI(h.visitedDirs, right) &&
		util.ContainsI(h.visitedDirs, down) &&
		util.ContainsI(h.visitedDirs, left) &&
		util.ContainsI(h.visitedDirs, up) {
		return true
	}
	return false
}

func contains(s []Pos, e Pos) bool {
	for _, a := range s {
		if a.X == e.X && a.Y == e.Y {
			return true
		}
	}
	return false
}
