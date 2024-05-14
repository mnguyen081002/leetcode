package main

func checkValidString(s string) bool {
	stack := []rune{}
	count := 0
	for _, v := range s {
		if v == '*' {
			count++
			continue
		}
		if v == ')' {
			if len(stack) == 0 {
				if count > 0 {
					count--
					continue
				} else {
					return false
				}
			}
			if stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return count >= len(stack)
}

func main() {
	// checkValidString("((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()")
	checkValidString("(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())")
}
