package main

import (
	"fmt"
	"strconv"
)

// evalRPN evaluates the given reverse polish notation expression and returns the result.
// It uses a stack to perform the calculations.
// The time complexity of evalRPN is O(n), where n is the number of tokens in the input.
func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, v := range tokens {
		switch v {
		case "*":
			stack[len(stack)-2] = stack[len(stack)-2] * stack[len(stack)-1]
		case "/":
			stack[len(stack)-2] = stack[len(stack)-2] / stack[len(stack)-1]
		case "+":
			stack[len(stack)-2] = stack[len(stack)-2] + stack[len(stack)-1]
		case "-":
			stack[len(stack)-2] = stack[len(stack)-2] - stack[len(stack)-1]
		default:
			i, _ := strconv.Atoi(v)
			stack = append(stack, i)
			continue
		}

		stack = stack[:len(stack)-1]
	}

	return stack[0]
}

func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))
}
