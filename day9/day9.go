package day9

import (
	"aoc24/util"
	"fmt"
	"strconv"
	"strings"
)

func Day9() {
	input, _ := util.ReadFileAsText("./day9/day9-input-test.txt")

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

	// part1(blocks, totalFilled)
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

	totalNewUsed := 0
	for i := len(disk) - 1; i > 0; i-- {
		for j := 0; j < len(disk); j++ {
			if len(disk[i].Used) <= len(disk[j].Free) {
				for k := 0; k < len(disk[i].Used); k++ {
					disk[j].NewUsed = append(disk[j].NewUsed, disk[i].Used[k]) //append to newly used
					disk[i].Used[k] = "."                                      // mark previously used with '.'
					disk[j].Free = util.RemoveS(disk[j].Free, 0)
					totalNewUsed += len(disk[j].Used)    // remove free from left block
					totalNewUsed += len(disk[j].NewUsed) // remove free from left block
				}
				break
			}
		}
	}

	util.WriteString(diskString(disk))

	// I think below code doesn't calculate things right - because I convert it to string, but it needs to remain a list
	// 9999999 is 1 number not 7 nines - this is why test input works but real puzzle doesn't
	diskMap := make(map[int]rune, totalNewUsed)
	for i, j := range diskString(disk) {
		if j != '.' {
			diskMap[i] = j
		}
	}
	total := 0
	for i, v := range diskMap {
		vv := string(v)
		idx, _ := strconv.Atoi(vv)
		total += i * idx
	}
	fmt.Println(total)
}

func diskString(disk []Block2) string {
	var diskString string
	for _, v := range disk {
		diskString += strings.Join(v.Used, "")
		diskString += strings.Join(v.NewUsed, "")
		diskString += strings.Join(v.Free, "")
	}

	return diskString
}
