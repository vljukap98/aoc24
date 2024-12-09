package util

import (
	"strconv"
	"unicode"
)

func generatePermutations(n int, current []string, result *[][]string) {
	if n == 0 {
		permutation := make([]string, len(current))
		copy(permutation, current)
		*result = append(*result, permutation)
		return
	}
	generatePermutations(n-1, append(current, "*"), result)
	generatePermutations(n-1, append(current, "+"), result)
	generatePermutations(n-1, append(current, "|"), result)
}

func Generate(n int) [][]string {
	var result [][]string
	generatePermutations(n, []string{}, &result)
	return result
}

func EvaluateExpression(s string) int {
	token := ""
	currOperator := "*"
	result := 0
	temp := s + "."

	for _, c := range temp {
		if unicode.IsDigit(c) {
			token += string(c)
		} else if c != ' ' {
			val, _ := strconv.Atoi(token)
			if currOperator == "*" {
				result *= val
			} else if currOperator == "+" {
				result += val
			}

			token = ""
			currOperator = string(c)
		}
	}

	return result
}

func EvaluateExpression2(s string) int {
	token := ""
	currOperator := "+"
	result := 0
	temp := s + "."
	operands := []int{}

	for _, c := range temp {
		if unicode.IsDigit(c) {
			token += string(c)
		} else if c != ' ' {
			if token != "" {
				val, _ := strconv.Atoi(token)
				if currOperator == "*" {
					if len(operands) > 0 {
						operands[len(operands)-1] *= val
					} else {
						result *= val
					}
				} else if currOperator == "+" {
					if len(operands) > 0 {
						operands[len(operands)-1] += val
					} else {
						result += val
					}
				} else if currOperator == "|" {
					strResult := strconv.Itoa(result)
					strVal := strconv.Itoa(val)
					concatenated := strResult + strVal
					result, _ = strconv.Atoi(concatenated)
				}
			}
			token = ""
			currOperator = string(c)
		}
	}

	return result
}
