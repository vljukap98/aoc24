package day6

import (
	"aoc24/util"
	"fmt"
	"strings"
)

func Day6() {
	input, _ := util.ReadLines("./day6/day6-input.txt")

	coordinates := make(map[coordinate]int)
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
}

func part1(mat [][]string, guard *guardState, coordinates map[coordinate]int) {
	for guard.CurrCoord.x != -1 && guard.CurrCoord.y != -1 {
		coordinates[guard.CurrCoord] += 1
		moveGuard(mat, guard)
	}

	visits := 0
	for _, coords := range coordinates {
		if coords > 0 {
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
	// simulate guard walking again and
	// for each stepped on point try putting an obstacle
	// for each obstacle put 2 guards - if they catch one another
	// guard is not in loop
	// solved with `shraddhaag` user's solution
}

func failedPart2(mat [][]string) {
	count := 0
	for y := 0; y < len(mat); y++ {
		for x := 0; x < len(mat[y]); x++ {
			//this doesn't work
			if mat[y][x] == "#" {
				if checkFirstRight(x, y, mat, directionRight, x) {
					count++
				}
				if checkFirstDown(x, y, mat, directionDown, y) {
					count++
				}
				if checkFirstUp(x, y, mat, directionUp, y) {
					count++
				}
				if checkFirstLeft(x, y, mat, directionLeft, x) {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}

func checkFirstRight(x int, y int, mat [][]string, direction int, firstX int) bool {
	if direction == directionRight {
		if x == len(mat[y])-1 || y == len(mat)-1 {
			return false
		}
		for i := x + 1; i < len(mat[y]); i++ {
			if mat[y+1][i] == "#" {
				return checkFirstRight(i, y+1, mat, directionDown, firstX)
			}
		}
	}

	if direction == directionDown {
		if x == 0 || y == len(mat)-1 {
			return false
		}
		for i := y + 1; i < len(mat); i++ {
			if mat[i][x-1] == "#" {
				return checkFirstRight(x-1, i, mat, directionLeft, firstX)
			}
		}
	}

	if direction == directionLeft {
		if x == 0 || y == 0 {
			return false
		}
		for i := x - 1; i > 0; i-- {
			if (mat[y-1][i] == "." || mat[y-1][i] == "^") && firstX-i == 1 {
				fmt.Println(i, y-1)
				return true
			}
		}
	}
	return false
}

func checkFirstDown(x int, y int, mat [][]string, direction int, firstY int) bool {
	if direction == directionDown {
		if x == 0 || y == len(mat)-1 { //check if char is on border
			return false
		}
		for i := y + 1; i < len(mat); i++ {
			if mat[i][x-1] == "#" {
				return checkFirstDown(x-1, i, mat, directionLeft, firstY)
			}
		}
	}

	if direction == directionLeft {
		if x == 0 || y == 0 { // if it's the first column being processed - there's nothing left of the first column
			return false
		}
		for i := x - 1; i > 0; i-- {
			if mat[y-1][i] == "#" { // check for first row
				return checkFirstDown(i, y-1, mat, directionUp, firstY)
			}
		}
	}

	if direction == directionUp {
		if x == len(mat[y])-1 || y == 0 {
			return false
		}
		for i := y - 1; i > 0; i-- {
			if (mat[i][x+1] == "." || mat[y-1][i] == "^") && firstY-i == -1 {
				fmt.Println(x+1, i)
				return true
			}
		}
	}

	return false
}

func checkFirstLeft(x int, y int, mat [][]string, direction int, firstX int) bool {
	if direction == directionLeft {
		if x == 0 || y == 0 { // if it's the first column being processed - there's nothing left of the first column
			return false
		}
		for i := x - 1; i > 0; i-- {
			if mat[y-1][i] == "#" { // check for first row
				return checkFirstLeft(i, y-1, mat, directionUp, firstX)
			}
		}
	}

	if direction == directionUp {
		if x == len(mat[y])-1 || y == 0 {
			return false
		}
		for i := y - 1; i > 0; i-- {
			if mat[i][x+1] == "#" {
				return checkFirstLeft(x, i, mat, directionRight, firstX)
			}
		}
	}

	if direction == directionRight {
		if x == len(mat[y])-1 || y == len(mat)-1 {
			return false
		}
		for i := x + 1; i < len(mat[y]); i++ {
			if (mat[y+1][i] == "." || mat[y-1][i] == "^") && firstX-i == -1 { //check for '^'
				fmt.Println(i, y+1)
				return true
			}
		}
	}

	return false
}

func checkFirstUp(x int, y int, mat [][]string, direction int, firstY int) bool {
	if direction == directionUp {
		if x == len(mat[y])-1 || y == 0 {
			return false
		}
		for i := y - 1; i > 0; i-- {
			if mat[i][x+1] == "#" {
				return checkFirstUp(x, i, mat, directionRight, firstY)
			}
		}
	}

	if direction == directionRight {
		for i := x + 1; i < len(mat[y+1]); i++ {
			if mat[y+1][i] == "#" {
				return checkFirstUp(i, y, mat, directionDown, firstY)
			}
		}
	}

	if direction == directionDown {
		if x == 0 || y == len(mat)-1 { //check if char is on border
			return false
		}
		for i := y + 1; i < len(mat)-1; i++ {
			if (mat[i][x-1] == "." || mat[y-1][i] == "^") && firstY-i == 1 { // check for '^'
				fmt.Println(x-1, i)
				return true
			}
		}
	}

	return false
}
