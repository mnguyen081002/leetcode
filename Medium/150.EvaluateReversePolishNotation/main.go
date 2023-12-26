package main

import (
	"fmt"
	"strconv"
)

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
