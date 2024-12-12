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
	// for each member in mat[y] if it's a 0
	// try finding the next number up, left, right, down - recursively

	hikingPaths := 0
	for y := 0; y < len(mat); y++ {
		for x := 0; x < len(mat[y]); x++ {
			currNum := mat[y][x]
			//TODO: track previously visited nine positions
			// sum all paths that lead to the visited nines?
			if currNum == 0 && canFindHikingPath(currNum, x, y, mat) {
				hikingPaths += 1
			}
		}
	}

	fmt.Println(hikingPaths)
}

func canFindHikingPath(currNum int, x int, y int, mat [][]int) bool {
	hikingDirection := hikingDirectionRight
	if hikingDirection == hikingDirectionRight &&
		!canFindHikingPathRight(currNum, x, y, mat) {
		hikingDirection = hikingDirectionDown
	} else if hikingDirection == hikingDirectionDown &&
		!canFindHikingPathDown(currNum, x, y, mat) {
		hikingDirection = hikingDirectionLeft
	} else if hikingDirection == hikingDirectionLeft &&
		!canFindHikingPathLeft(currNum, x, y, mat) {
		hikingDirection = hikingDirectionUp
	} else if hikingDirection == hikingDirectionUp &&
		!canFindHikingPathUp(currNum, x, y, mat) {
		return false
	}
	return true
}

func canFindHikingPathRight(currNum int, x int, y int, mat [][]int) bool {
	if currNum == 9 {
		return true
	}
	if x == len(mat[y])-1 {
		return false
	}
	if currNum+1 == mat[y][x+1] {
		currNum := mat[y][x+1]
		return canFindHikingPathRight(currNum, x+1, y, mat)
	}
	return false
}

func canFindHikingPathDown(currNum int, x int, y int, mat [][]int) bool {
	if currNum == 9 {
		return true
	}
	if y == len(mat)-1 {
		return false
	}
	if currNum+1 == mat[y+1][x] {
		currNum := mat[y+1][x]
		return canFindHikingPathDown(currNum, x, y+1, mat)
	}
	return false
}

func canFindHikingPathLeft(currNum int, x int, y int, mat [][]int) bool {
	if currNum == 9 {
		return true
	}
	if x == 0 {
		return false
	}
	if currNum+1 == mat[y][x-1] {
		currNum := mat[y][x-1]
		return canFindHikingPathLeft(currNum, x-1, y, mat)
	}
	return false
}

func canFindHikingPathUp(currNum int, x int, y int, mat [][]int) bool {
	if currNum == 9 {
		return true
	}
	if y == 0 {
		return false
	}
	if currNum+1 == mat[y-1][x] {
		currNum := mat[y-1][x]
		return canFindHikingPathUp(currNum, x, y-1, mat)
	}
	return false
}
