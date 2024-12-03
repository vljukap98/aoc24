package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadFileAsText(path string) (string, error) {
	var fileName string = path
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Cannot read file %s", fileName)
		return "", err
	}
	fileContent := string(file)

	return fileContent, err
}

func Day3Part1() {
	mulString, _ := ReadFileAsText("./day3/day3-input.txt")
	calculateMuls(mulString)
}

func Day3Part2() {
	mulString, _ := ReadFileAsText("./day3/day3-input.txt")
	filteredDos := filterDos(mulString)
	calculateMuls(filteredDos)
}

func calculateMuls(input string) {
	re := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")
	muls := re.FindAllStringSubmatch(input, -1)

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

func filterDos(s string) string {
	var filteredDos []string
	dos := strings.Split(s, "do()")
	for _, do := range dos {
		part := strings.Split(do, "don't()")[0]
		filteredDos = append(filteredDos, part)
	}
	return strings.Join(filteredDos, "")
}
