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
	CurrPosX      int
	CurrPosY      int
	CurrDirection int
}

type coordinate struct {
	x       int
	y       int
	visited bool
}

func Day6() {
	input, _ := util.ReadLines("./day6/day6-input.txt")

	var coordinates []coordinate
	var startPos coordinate
	mat := make([][]string, 0)

	for y, lines := range input {
		points := strings.Split(lines, "")
		mat = append(mat, make([]string, 0))
		for x, point := range points {
			mat[y] = append(mat[y], point)
			if point == "^" {
				startPos = coordinate{x, y, true}
			}
			coordinates = append(coordinates, coordinate{x, y, false})
		}
	}

	fmt.Println(startPos)
	fmt.Println(coordinates)
	//TODO: start walking
}

func moveGuard(mat [][]string, guard guardState) {
	if guard.CurrDirection == directionUp {
		nextPos := mat[guard.CurrPosY-1][guard.CurrPosX]

		if nextPos == "#" { //guard comes across an obstacle - change direction
			guard.CurrDirection = directionRight
		} else if nextPos == "." { //guard comes across a path - keep moving in the direction
			guard.CurrPosY -= 1
		} else { //guard will leave the mapped area
			guard.CurrPosY = -1
		}
	} else if guard.CurrDirection == directionRight {
		nextPos := mat[guard.CurrPosY][guard.CurrPosX+1]

		if nextPos == "#" { //guard comes across an obstacle - change direction
			guard.CurrDirection = directionDown
		} else if nextPos == "." { //guard comes across a path - keep moving in the direction
			guard.CurrPosX += 1
		} else { //guard will leave the mapped area
			guard.CurrPosX = -1
		}
	} else if guard.CurrDirection == directionDown {
		nextPos := mat[guard.CurrPosY+1][guard.CurrPosX]

		if nextPos == "#" { //guard comes across an obstacle - change direction
			guard.CurrDirection = directionLeft
		} else if nextPos == "." { //guard comes across a path - keep moving in the direction
			guard.CurrPosY += 1
		} else { //guard will leave the mapped area
			guard.CurrPosY = -1
		}
	} else if guard.CurrDirection == directionLeft {
		nextPos := mat[guard.CurrPosY][guard.CurrPosX-1]

		if nextPos == "#" { //guard comes across an obstacle - change direction
			guard.CurrDirection = directionUp
		} else if nextPos == "." { //guard comes across a path - keep moving in the direction
			guard.CurrPosX -= 1
		} else { //guard will leave the mapped area
			guard.CurrPosX = -1
		}
	}
}
