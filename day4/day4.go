package day4

import (
	"aoc24/util"
	"fmt"
	"strings"
)

func Day4() {
	lines, _ := util.ReadLines("./day4/day4-input.txt")
	var matrix [][]string
	for _, line := range lines {
		lineChars := strings.Split(line, "")
		matrix = append(matrix, (lineChars))
	}

	part1(matrix)
	part2(matrix)
}

func part1(mat [][]string) {
	count := 0

	for y := 3; y < len(mat)-2; y++ {
		for x := 3; x < len(mat[y])-2; x++ {
			//right
			if mat[y][x] == "X" && mat[y][x+1] == "M" && mat[y][x+2] == "A" && mat[y][x+3] == "S" {
				count++
			}
			//left
			if mat[y][x] == "X" && mat[y][x-1] == "M" && mat[y][x-2] == "A" && mat[y][x-3] == "S" {
				count++
			}
			//up
			if mat[y][x] == "X" && mat[y-1][x] == "M" && mat[y-2][x] == "A" && mat[y-3][x] == "S" {
				count++
			}
			//down
			if mat[y][x] == "X" && mat[y+1][x] == "M" && mat[y+2][x] == "A" && mat[y+3][x] == "S" {
				count++
			}
			//down right
			if mat[y][x] == "X" && mat[y+1][x+1] == "M" && mat[y+2][x+2] == "A" && mat[y+3][x+3] == "S" {
				count++
			}
			//up right
			if mat[y][x] == "X" && mat[y-1][x+1] == "M" && mat[y-2][x+2] == "A" && mat[y-3][x+3] == "S" {
				count++
			}
			//down left
			if mat[y][x] == "X" && mat[y+1][x-1] == "M" && mat[y+2][x-2] == "A" && mat[y+3][x-3] == "S" {
				count++
			}
			//up left
			if mat[y][x] == "X" && mat[y-1][x-1] == "M" && mat[y-2][x-2] == "A" && mat[y-3][x-3] == "S" {
				count++
			}
		}
	}

	fmt.Println(count)
}

func part2(mat [][]string) {
	count := 0

	for y := 3; y < len(mat)-2; y++ {
		for x := 3; x < len(mat[y])-2; x++ {
			if mat[y][x] == "A" {

				// up left M
				if mat[y-1][x-1] == "M" && mat[y+1][x+1] == "S" {
					// down left MAS
					downLeftMAS := mat[y+1][x-1] == "M" && mat[y-1][x+1] == "S"
					downLeftSAM := mat[y+1][x-1] == "S" && mat[y-1][x+1] == "M"
					if downLeftMAS || downLeftSAM {
						count++
					}
				}
				// up left S
				if mat[y-1][x-1] == "S" && mat[y+1][x+1] == "M" {
					downLeftMAS := mat[y+1][x-1] == "M" && mat[y-1][x+1] == "S"
					downLeftSAM := mat[y+1][x-1] == "S" && mat[y-1][x+1] == "M"
					if downLeftMAS || downLeftSAM {
						count++
					}
				}
			}
		}
	}

	fmt.Println(count)
}
