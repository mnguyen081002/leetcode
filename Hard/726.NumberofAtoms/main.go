package main

import (
	"sort"
	"strconv"
	"strings"
)

func countOfAtoms(formula string) string {
	multiplier := stack{}
	count := make(map[string]int)
	current, length := 1, 0
	for i := len(formula) - 1; i >= 0; i-- {
		c := formula[i]
		if c >= '0' && c <= '9' {
			if length == 0 {
				current = int(c - '0')
			} else {
				current += int(c-'0') * pow10(length)
			}
			length++
			continue
		}
		if c == '(' {
			multiplier.pop()
		} else if c == ')' {
			if multiplier.empty() {
				multiplier.push(current)
			} else {
				multiplier.push(current * multiplier.top())
			}
		} else {
			var element string
			if c >= 'a' && c <= 'z' {
				element = formula[i-1 : i+1]
				i--
			} else {
				element = formula[i : i+1]
			}
			if multiplier.empty() {
				count[element] += current
			} else {
				count[element] += multiplier.top() * current
			}
		}
		current, length = 1, 0
	}
	keys := []string{}
	for k := range count {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	var result strings.Builder
	for _, k := range keys {
		result.WriteString(k)
		if count[k] > 1 {
			result.WriteString(strconv.Itoa(count[k]))
		}
	}
	return result.String()
}

func pow10(i int) int {
	res := 1
	for i > 0 {
		res *= 10
		i--
	}
	return res
}

type stack []int

func (s *stack) push(v int) {
	*s = append(*s, v)
}
func (s *stack) pop() int {
	if len(*s) == 0 {
		return -1
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
func (s *stack) top() int {
	if len(*s) == 0 {
		return -1
	}
	return (*s)[len(*s)-1]
}
func (s *stack) empty() bool {
	return len(*s) == 0
}
