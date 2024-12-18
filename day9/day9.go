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
		blocks = append(blocks, Block{id, filledSpace + freeSpace, filledSpace, freeSpace})
	}

	part1(blocks, totalFilled)
	part2(blocks)
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
		blockId, _ := strconv.Atoi(string(disk[i]))
		total += i * blockId
	}
	fmt.Println(total)
}

func part2(blocks []Block) {
	var disk []Block2
	for _, b := range blocks {
		var used []string
		for i := 0; i < b.FilledSpace; i++ {
			used = append(used, b.Id)
		}
		var free []string
		for i := 0; i < b.FreeSpace; i++ {
			free = append(free, ".")
		}
		disk = append(disk, Block2{used, make([]string, 0), free})
	}

	for i := len(disk) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if len(disk[j].Free) >= len(disk[i].Used) {
				for k := 0; k < len(disk[i].Used); k++ {
					disk[j].NewUsed = append(disk[j].NewUsed, disk[i].Used[k]) //append to newly used
					disk[i].Used[k] = "."                                      // mark previously used with '.'
					disk[j].Free = util.RemoveS(disk[j].Free, 0)
				}
				break
			}
		}
	}

	var diskList []string
	for _, block := range disk {
		diskList = append(diskList, block.Used...)
		diskList = append(diskList, block.NewUsed...)
		diskList = append(diskList, block.Free...)
	}

	total := 0

	for i, b := range diskList {
		if b != "." {
			idx, _ := strconv.Atoi(b)
			total += i * idx

		}
	}

	fmt.Println(total)
}

func stringDisk(disk []Block2) string {
	var stringDisk string
	for _, block := range disk {
		for _, u := range block.Used {
			stringDisk += u
		}
		for _, u := range block.NewUsed {
			stringDisk += u
		}
		for _, u := range block.Free {
			stringDisk += u
		}
	}
	return stringDisk
}
