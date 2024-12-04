package day2

import (
	"aoc24/util"
	"fmt"
	"math"
	"strings"
)

func Day2() {}

func Day2Part1() {
	inputRows, _ := util.ReadLines("./day2/day2-input.txt")

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
	inputRows, _ := util.ReadLines("./day2/day2-input.txt")

	safeLevels := 0

	for _, row := range inputRows {
		rowElems := util.ArrStrToI(strings.Split(row, " "))
		fmt.Println(rowElems)
		if isSafe2(rowElems) || canBeSafeByRemovingOne(rowElems) {
			safeLevels += 1
			fmt.Println("safe")
		} else {
			fmt.Println("unsafe")
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

func isSafe2(report []int) bool {
	if len(report) < 2 {
		return true
	}
	increasing := true
	decreasing := true

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff < 1 || diff > 3 {
			return false
		}
		if diff > 0 {
			decreasing = false
		} else if diff < 0 {
			increasing = false
		}
	}
	return increasing || decreasing
}

// Checks if the report can be made safe by removing a single element
func canBeSafeByRemovingOne(report []int) bool {
	for i := range report {
		newReport := make([]int, 0, len(report)-1)
		newReport = append(newReport, report[:i]...)
		newReport = append(newReport, report[i+1:]...)
		if isSafe(newReport) {
			return true
		}
	}
	return false
}
