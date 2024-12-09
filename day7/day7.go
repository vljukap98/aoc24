package day7

import (
	"aoc24/util"
	"strconv"
	"strings"
)

func Day7() {
	input, _ := util.ReadLines("./day7/day7-input.txt")

	part1(input)
}

func part1(input []string) {
	count := 0
	for _, line := range input {
		before, after, _ := strings.Cut(line, ": ")
		//TODO: generate permutations for 'after'
		// check if any permutation result equals 'before'
		afterPerms := getPermutations(strings.Split(after, " "))
		for _, ap := range afterPerms {
			postfix := util.InfixToPostfix(ap)
			eval := util.EvaluatePostfix(postfix)
			beforeInt, _ := strconv.Atoi(before)
			if beforeInt == eval {
				count += eval
				break
			}
		}
	}
}

func getPermutations(nums []string) []string {
	operatorPerms := util.Generate(len(nums) - 1)

	var perms []string
	for _, opPer := range operatorPerms {
		var infixString string

		opArr := strings.Split(opPer, "")
		for i, op := range opArr {
			infixString += nums[i]
			infixString += op
		}
		infixString += nums[len(nums)-1]
		perms = append(perms, infixString)
	}

	return perms
}
