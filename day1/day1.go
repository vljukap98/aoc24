package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day1() {
	input, _ := os.Open("./day1/day1-input.txt")
	defer input.Close()

	var locationIds []string

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		locationIds = append(locationIds, scanner.Text())
	}

	var locationIdList1, locationIdList2 []string

	for _, locationId := range locationIds {
		line := strings.Split(locationId, "   ")
		locationIdList1 = append(locationIdList1, line[0])
		locationIdList2 = append(locationIdList2, line[1])
	}

	part1(locationIdList1, locationIdList1, len(locationIds))
	part2(locationIdList1, locationIdList2)
}

func part1(locationIdList1 []string, locationIdList2 []string, size int) {

	sort.Slice(locationIdList1, func(i, j int) bool {
		return locationIdList1[i] < locationIdList1[j]
	})

	sort.Slice(locationIdList2, func(i, j int) bool {
		return locationIdList2[i] < locationIdList2[j]
	})

	totalSum := 0.0

	for i := 0; i < size; i++ {
		list1Member, _ := strconv.Atoi(locationIdList1[i])
		list2Member, _ := strconv.Atoi(locationIdList2[i])

		memberDistance := math.Abs(float64(list1Member - list2Member))
		totalSum += memberDistance
	}

	fmt.Printf("%f\n", totalSum)
}

func part2(locationIdList1 []string, locationIdList2 []string) {

	freq := make(map[int]int)
	for _, num := range locationIdList2 {
		n, _ := strconv.Atoi(num)
		freq[n] = freq[n] + 1
	}

	totalSum := 0

	for _, number := range locationIdList1 {
		n, _ := strconv.Atoi(number)
		occurrences, ok := freq[n]
		if ok && occurrences > 0 {
			totalSum += n * occurrences
		}
	}

	fmt.Printf("%d\n", totalSum)
}
