package main

func generateParenthesis(n int) []string {
	var res []string
	backtrack(&res, "", 0, 0, n)
	return res
}

func backtrack(res *[]string, cur string, open, close, max int) {
	if len(cur) == max*2 {
		*res = append(*res, cur)
		return
	}

	if open < max {
		backtrack(res, cur+"(", open+1, close, max)
	}
	if close < open {
		backtrack(res, cur+")", open, close+1, max)
	}
}

func main() {
	generateParenthesis(3)
}
