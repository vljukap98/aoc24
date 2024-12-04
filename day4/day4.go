package day4

import (
	"aoc24/util"
	"strings"
)

func Day4() {
	lines, _ := util.ReadLines("./day4/day4-input.txt")
	var matrix [][]int
	for _, line := range lines {
		lineChars := strings.Split(line, "\"")
		matrix = append(matrix, util.ArrStrToI(lineChars))
	}

	part1(matrix)
}

func part1(mat [][]int) {
	xmasCount := 0
	for x, y := 0, 0; x < len(mat[0]); x, y = x+1, y+1 {
		xmasCount += checkHorizontalLeft()
		xmasCount += checkHorizontalRight()
		xmasCount += checkVerticalDown()
		xmasCount += checkVerticalUp()
		xmasCount += checkDiagonalLeftDown()
		xmasCount += checkDiagonalLeftUp()
		xmasCount += checkDiagonalRightDown()
		xmasCount += checkDiagonalRightUp()
	}
}

func checkHorizontalLeft() int {
	panic("unimplemented")
}

func checkHorizontalRight() int {
	panic("unimplemented")
}

func checkVerticalDown() int {
	panic("unimplemented")
}

func checkVerticalUp() int {
	panic("unimplemented")
}

func checkDiagonalLeftDown() int {
	panic("unimplemented")
}

func checkDiagonalLeftUp() int {
	panic("unimplemented")
}

func checkDiagonalRightDown() int {
	panic("unimplemented")
}

func checkDiagonalRightUp() int {
	panic("unimplemented")
}
