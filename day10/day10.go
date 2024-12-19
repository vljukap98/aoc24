package day10

import (
	"aoc24/util"
	"fmt"
)

func Day10() {
	mat, _ := util.ReadMatI("./day10/day10-input-test.txt")
	part1(mat)
}

func part1(mat [][]int) {
	h := len(mat)
	w := len(mat[0])
	count := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			currNum := mat[y][x]
			hikingPaths := make([]int, 0)
			hiker := Hiker{Pos{x, y}, currNum, right, []Pos{{x, y}}, []int{left}}
			hikingPaths = append(hikingPaths, findHikingPath(hiker, mat, x, y)...)
			count += len(hikingPaths)
		}
	}

	fmt.Println(count)
}

func findHikingPath(hiker Hiker, mat [][]int, x int, y int) []int {
	// hiker might find 1 path to the 9
	// but there could be paths to other nines
	if hiker.currNum == 9 {
		return append(make([]int, 0), hiker.currNum)
	}

	h := len(mat)
	w := len(mat[0])

	if x < 0 {
		return make([]int, 0)
	}
	if y < 0 {
		return make([]int, 0)
	}
	if x >= w {
		return make([]int, 0)
	}
	if y >= h {
		return make([]int, 0)
	}

	hiker.currNum++

	hikingPaths := make([]int, 0)
	hikingPaths = append(hikingPaths, findHikingPath(hiker, mat, x+1, y)...)
	hikingPaths = append(hikingPaths, findHikingPath(hiker, mat, x, y+1)...)
	hikingPaths = append(hikingPaths, findHikingPath(hiker, mat, x-1, y)...)
	hikingPaths = append(hikingPaths, findHikingPath(hiker, mat, x, y-1)...)

	return hikingPaths
}
