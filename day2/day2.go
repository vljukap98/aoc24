package day2

import (
	"aoc24/util"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func readFile(inputPath string) []string {
	input, _ := os.Open(inputPath)
	defer input.Close()

	var rows []string

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		row := scanner.Text()
		rows = append(rows, row)
	}
	return rows
}

func Day2Part1() {
	inputRows := readFile("./day2/day2-input.txt")

	safeLevels := 0

	for _, row := range inputRows {
		rowNumbers := strings.Split(row, " ")

		if isSafe(util.ArrStrToI(rowNumbers)) {
			safeLevels += 1
		} else {
			continue
		}
	}

	fmt.Printf("%d", safeLevels)
}

func Day2Part2() {
	inputRows := readFile("./day2/day2-input.txt")

	safeLevels := 0

	for _, row := range inputRows {
		rowElems := util.ArrStrToI(strings.Split(row, " "))

		if isSafe2(rowElems, 0) {
			safeLevels += 1
		} else {
			continue
		}
	}

	fmt.Printf("%d", safeLevels)
}

func isSafe(numbers []int) bool {
	if numbers[0] == numbers[1] {
		return false
	}

	var isIncreasing1 bool
	var isIncreasing2 bool

	for i := 0; i < len(numbers)-1; i++ {
		if i > 0 {
			if numbers[i] < numbers[i+1] {
				isIncreasing2 = true
			} else {
				isIncreasing2 = false
			}

			if isIncreasing1 != isIncreasing2 {
				return false
			}
		} else {
			if numbers[i] < numbers[i+1] {
				isIncreasing1 = true
			} else {
				isIncreasing1 = false
			}
		}

		distanceInt := numbers[i] - numbers[i+1]
		distance := math.Abs(float64(distanceInt))
		if distance < 1 || distance > 3 {
			return false
		}
	}

	return true
}

func isSafe2(numbers []int, iter int) bool {
	badCount := 0
	var isIncreasing1 bool
	var isIncreasing2 bool

	for i := 0; i < len(numbers); i++ {
		if numbers[i] == numbers[i+1] {
			if badCount > 0 {
				return false
			} else {
				badCount++
				i--
				//remove item
				continue
			}
		}
		if i > 0 {
			if numbers[i] < numbers[i+1] {
				isIncreasing2 = true
			} else {
				isIncreasing2 = false
			}

			if isIncreasing1 != isIncreasing2 {
				if badCount > 0 {
					return false
				} else {
					badCount++
					i--
					//remove item
					continue
				}
			}
		} else {
			if numbers[i] < numbers[i+1] {
				isIncreasing1 = true
			} else {
				isIncreasing1 = false
			}
		}

		distanceInt := numbers[i] - numbers[i+1]
		distance := math.Abs(float64(distanceInt))
		if distance < 1 || distance > 3 {
			if badCount > 0 {
				return false
			} else {
				badCount++
				i--
				//remove item
				continue
			}
		}

	}

	return badCount == 0
}
