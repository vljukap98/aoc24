package day6

import (
	"aoc24/util"
	"fmt"
	"strings"
)

const (
	_ = iota
	directionUp
	directionRight
	directionDown
	directionLeft
)

type guardState struct {
	CurrCoord     coordinate
	CurrDirection int
}

type coordinate struct {
	x int
	y int
}

func Day6() {
	input, _ := util.ReadLines("./day6/day6-input-test.txt")

	coordinates := make(map[coordinate]bool)
	guard := guardState{}
	mat := make([][]string, 0)

	for y, lines := range input {
		points := strings.Split(lines, "")
		mat = append(mat, make([]string, 0))
		for x, point := range points {
			mat[y] = append(mat[y], point)
			coordinate := coordinate{x, y}
			if point == "^" {
				guard = guardState{coordinate, directionUp}
			}
		}
	}

	part1(mat, &guard, coordinates)
	part2(mat)
}

func part1(mat [][]string, guard *guardState, coordinates map[coordinate]bool) {
	// for guard.CurrCoord.x != -1 && guard.CurrCoord.y != -1 {
	// 	coordinates[guard.CurrCoord] = true
	// 	moveGuard(mat, guard)
	// }

	visits := 0
	for _, coords := range coordinates {
		if coords {
			visits += 1
		}
	}

	fmt.Println(visits)
}

func moveGuard(mat [][]string, guard *guardState) {
	if guard.CurrDirection == directionUp {
		nextPos := mat[guard.CurrCoord.y-1][guard.CurrCoord.x]

		if nextPos == "#" { //guard comes across an obstacle - change direction
			guard.CurrDirection = directionRight
		} else if nextPos == "." || nextPos == "^" { //guard comes across a path - keep moving in the direction
			guard.CurrCoord.y -= 1
		} else { //guard will leave the mapped area
			guard.CurrCoord.y = -1
		}
	} else if guard.CurrDirection == directionRight {
		nextPos := mat[guard.CurrCoord.y][guard.CurrCoord.x+1]

		if nextPos == "#" { //guard comes across an obstacle - change direction
			guard.CurrDirection = directionDown
		} else if nextPos == "." || nextPos == "^" { //guard comes across a path - keep moving in the direction
			guard.CurrCoord.x += 1
		} else { //guard will leave the mapped area
			guard.CurrCoord.x = -1
		}
	} else if guard.CurrDirection == directionDown {
		nextPos := mat[guard.CurrCoord.y+1][guard.CurrCoord.x]

		if nextPos == "#" { //guard comes across an obstacle - change direction
			guard.CurrDirection = directionLeft
		} else if nextPos == "." || nextPos == "^" { //guard comes across a path - keep moving in the direction
			guard.CurrCoord.y += 1
		} else { //guard will leave the mapped area
			guard.CurrCoord.y = -1
		}
	} else if guard.CurrDirection == directionLeft {
		nextPos := mat[guard.CurrCoord.y][guard.CurrCoord.x-1]

		if nextPos == "#" { //guard comes across an obstacle - change direction
			guard.CurrDirection = directionUp
		} else if nextPos == "." || nextPos == "^" { //guard comes across a path - keep moving in the direction
			guard.CurrCoord.x -= 1
		} else { //guard will leave the mapped area
			guard.CurrCoord.x = -1
		}
	}
}

func part2(mat [][]string) {
	count := 0
	for y := 1; y < len(mat)-1; y++ {
		for x := 1; x < len(mat[y])-1; x++ {
			if mat[y][x] == "#" && (checkFirstRight(x, y, mat, directionRight, x) ||
				checkFirstUp(x, y, mat, directionUp) ||
				checkFirstLeft(x, y, mat, directionLeft)) {
				count++
			}
		}
	}

	fmt.Println(count)
}

func checkFirstRight(x int, y int, mat [][]string, direction int, firstX int) bool {
	if direction == directionRight {
		for i := x + 1; i < len(mat[y+1]); i++ {
			if mat[y+1][i] == "#" {
				return checkFirstRight(i, y, mat, directionDown, x)
			}
		}
	}

	if direction == directionDown {
		for i := y + 1; i < len(mat)-1; i++ {
			if mat[i][x-1] == "#" {
				return checkFirstRight(x, i, mat, directionLeft, x)
			}
		}
	}

	//TODO: check if from current x to first x are not obstacles
	// this means I can put an obstacle to first x+-1 to put the guard on the loop

	return false
}

func checkFirstUp(x int, y int, mat [][]string, direction int) bool {
	if direction == directionUp {
		for i := y - 1; i > 0; i-- {
			if mat[i][x+1] == "#" {
				return checkFirstUp(x, i, mat, directionRight)
			}
		}
	}

	if direction == directionRight {
		for i := x + 1; i < len(mat[y+1]); i++ {
			if mat[y+1][i] == "#" {
				return checkFirstUp(i, y, mat, directionDown)
			}
		}
	}

	//TODO: check if from current y to first one are not obstacles

	return false
}

func checkFirstLeft(x int, y int, mat [][]string, direction int) bool {
	if direction == directionLeft {
		for i := x - 1; i > 0; i++ {
			if mat[y-1][i] == "#" { // check for first row
				return checkFirstLeft(i, y, mat, directionUp)
			}
		}
	}

	if direction == directionUp {
		for i := y - 1; i > 0; i-- {
			if mat[i][x+1] == "#" {
				return checkFirstUp(x, i, mat, directionRight)
			}
		}
	}

	//TODO: check if from current x there are no obstacles

	return false
}
