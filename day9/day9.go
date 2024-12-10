package day9

import (
	"aoc24/util"
	"fmt"
	"strconv"
	"strings"
)

func Day9() {
	input, _ := util.ReadFileAsText("./day9/day9-input.txt")

	numbers := strings.Split(input, "")
	var blocks []Block
	totalFilled := 0
	totalFree := 0
	for i, j := 0, 0; i < len(numbers)-1; i, j = i+2, j+1 {
		id := strconv.Itoa(j)
		filledSpace, _ := strconv.Atoi(numbers[i])
		freeSpace, _ := strconv.Atoi(numbers[i+1])
		totalFilled += filledSpace
		totalFree += freeSpace
		blocks = append(blocks, Block{id, filledSpace, freeSpace})
	}

	part1(blocks, totalFilled)
}

func part1(blocks []Block, totalFilled int) {
	var disk []string
	for _, b := range blocks {
		for i := 0; i < b.FilledSpace; i++ {
			disk = append(disk, b.Id)
		}
		for i := 0; i < b.FreeSpace; i++ {
			disk = append(disk, ".")
		}
	}

	j := len(disk) - 1

	for i := 0; i < totalFilled; i++ {
		if disk[i] == "." {
			for disk[j] == "." {
				j--
			}

			disk[i] = disk[j]
			disk[j] = "."
		}

	}

	total := 0

	for i := 0; i < totalFilled; i++ {
		blockId, _ := strconv.Atoi(disk[i])
		total += i * blockId
	}
	fmt.Println(total)
}
