package util

import (
	"strconv"
	"strings"
	"unicode"
)

// generatePermutations recursively generates all permutations of '*' and '+' of length n.
func generatePermutations(n int, current string, result *[]string) {
	if n == 0 {
		*result = append(*result, current)
		return
	}
	generatePermutations(n-1, current+"*", result)
	generatePermutations(n-1, current+"+", result)
}

func Generate(n int) []string {
	var result []string
	generatePermutations(n, "", &result)
	return result
}

// precedence returns the precedence of the given operator
func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}

// applyOp applies the operator on the two operands
func applyOp(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}

// isOperator checks if the character is an operator
func isOperator(c string) bool {
	return c == "+" || c == "-" || c == "*" || c == "/"
}

// evaluatePostfix evaluates the postfix expression
func EvaluatePostfix(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		if isOperator(token) {
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result := applyOp(a, b, token)
			stack = append(stack, result)
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

// infixToPostfix converts an infix expression to a postfix expression
func InfixToPostfix(expression string) []string {
	output := []string{}
	operators := []string{}
	tokens := strings.Fields(expression)

	for _, token := range tokens {
		if unicode.IsDigit(rune(token[0])) {
			output = append(output, token)
		} else if isOperator(token) {
			for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(token) {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		}
	}

	for len(operators) > 0 {
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return output
}
