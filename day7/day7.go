package day7

import (
	"aoc24/util"
	"fmt"
	"strconv"
	"strings"
)

func Day7() {
	input, _ := util.ReadLines("./day7/day7-input.txt")

	part1(input)
	part2(input)
}

func part1(input []string) {
	sum := 0
	for _, line := range input {
		before, after, _ := strings.Cut(line, ": ")
		afterPerms := getPermutations(strings.Split(after, " "))
		for _, ap := range afterPerms {
			eval := util.EvaluateExpression(ap)
			beforeInt, _ := strconv.Atoi(before)
			if beforeInt == eval {
				sum += eval
				break
			}
		}
	}
	fmt.Println(sum)
}

func part2(input []string) {
	sum := 0
	for _, line := range input {
		before, after, _ := strings.Cut(line, ": ")
		afterPerms := getPermutations(strings.Split(after, " "))
		for _, ap := range afterPerms {
			eval := util.EvaluateExpression2(ap)
			beforeInt, _ := strconv.Atoi(before)
			if beforeInt == eval {
				sum += eval
				break
			}
		}
	}
	fmt.Println(sum)
}

func getPermutations(nums []string) []string {
	operatorPerms := util.Generate(len(nums) - 1)

	var perms []string
	for _, opPer := range operatorPerms {
		var infixString string

		for i, op := range opPer {
			infixString += nums[i]
			infixString += op
		}
		infixString += nums[len(nums)-1]
		perms = append(perms, infixString)
	}

	return perms
}
