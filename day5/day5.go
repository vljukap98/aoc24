package day5

import (
	"aoc24/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func getRules(input string) []string {
	ruleRe, _ := regexp.Compile(`\d{2}\|\d{2}`)
	return ruleRe.FindAllString(input, -1)
}

func getUpdates(input string) []string {
	updateRe, _ := regexp.Compile(`(\d+,)+\d+`)
	return updateRe.FindAllString(input, -1)
}

func Day5() {
	input, _ := util.ReadFileAsText("./day5/day5-input.txt")

	rules := getRules(input)
	updates := getUpdates(input)

	part1(rules, updates)
	part2(rules, updates)
}

func part1(rules []string, updates []string) {
	sum := 0
	for _, u := range updates {
		update := strings.Split(u, ",")
		if updateCompliesWithRules(update, rules) {
			middleElem, _ := strconv.Atoi(update[len(update)/2])
			sum += middleElem
		}
	}

	fmt.Println(sum)
}

func part2(rules []string, updates []string) {
	sum := 0

	var nonCompliant []string

	for _, u := range updates {
		update := strings.Split(u, ",")
		if !updateCompliesWithRules(update, rules) {
			nonCompliant = append(nonCompliant, u)
		}
	}

	for _, ncu := range nonCompliant {
		nonCompliantUpdate := strings.Split(ncu, ",")
		for !updateCompliesWithRules(nonCompliantUpdate, rules) {
			fixNonCompliant(nonCompliantUpdate, rules)
		}
		middleElem, _ := strconv.Atoi(nonCompliantUpdate[len(nonCompliantUpdate)/2])
		sum += middleElem
	}

	fmt.Println(sum)
}

func updateCompliesWithRules(update []string, rules []string) bool {
	for i := 0; i < len(update)-1; i++ {
		if !compliesWithRules(update[i], update[i+1], rules) {
			return false
		}
	}
	return true
}

func compliesWithRules(num1 string, num2 string, rules []string) bool {
	for _, rule := range rules {
		ruleFirstNum := strings.Split(rule, "|")[0]
		ruleSecondNum := strings.Split(rule, "|")[1]
		if num1 == ruleSecondNum && num2 == ruleFirstNum {
			return false
		}
	}

	return true
}

func fixNonCompliant(update []string, rules []string) {
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			if !compliesWithRules(update[i], update[j], rules) {
				temp := update[i]
				update[i] = update[i+1]
				update[i+1] = temp
			}
		}
	}
}
