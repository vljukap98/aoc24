package day18

import (
	"aoc24/util"
	"strconv"
	"strings"
)

func Day18() {
	in, _ := util.ReadLines("./day18/day18-input-test.txt")
	mat := initMat(6)
	firstByteCount := 12
	part1(in, mat, firstByteCount)
	// firstByteCount := 1024
	// mat := initMat(70)
	// part1(in, mat)
}

func part1(lines []string, mat [][]string, firstByteCount int) {
	for i := 0; i < firstByteCount; i++ {
		// for _, v := range lines {
		// left, right, _ := strings.Cut(v, ",")
		left, right, _ := strings.Cut(lines[i], ",")
		x, _ := strconv.Atoi(left)
		y, _ := strconv.Atoi(right)
		mat[y][x] = "#"
	}
	util.PrintMat(mat)
}

func initMat(gridSize int) [][]string {
	var mat [][]string
	for i := 0; i <= gridSize; i++ {
		line := make([]string, 0)
		for j := 0; j <= gridSize; j++ {
			line = append(line, ".")
		}
		mat = append(mat, line)
	}
	return mat
}
