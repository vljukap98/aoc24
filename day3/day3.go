package day3

import (
	"aoc24/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day3() {
	mulString, _ := util.ReadFileAsText("./day3/day3-input.txt")
	part1(mulString)
	part2(mulString)
}

func part1(mulString string) {
	re := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")
	muls := re.FindAllStringSubmatch(mulString, -1)

	var numStrings []string
	for _, mul := range muls {
		re = regexp.MustCompile("\\d{1,3},\\d{1,3}")
		numToMulString := re.FindAllStringSubmatch(mul[0], -1)
		for _, numString := range numToMulString {
			numStrings = append(numStrings, numString[0])
		}
	}

	totalSum := 0

	for _, num := range numStrings {
		nums := strings.Split(num, ",")
		first, _ := strconv.Atoi(nums[0])
		second, _ := strconv.Atoi(nums[1])

		mul := first * second
		totalSum += mul
	}

	fmt.Printf("%d", totalSum)
}

func part2(mulString string) {
	filteredDos := filterDos(mulString)
	part1(filteredDos)
}

func filterDos(s string) string {
	var filteredDos []string
	dos := strings.Split(s, "do()")
	for _, do := range dos {
		part := strings.Split(do, "don't()")[0]
		filteredDos = append(filteredDos, part)
	}
	return strings.Join(filteredDos, "")
}
