package main

func minAddToMakeValid(s string) int {
	count, need := 0, 0
	for _, v := range s {
		if v == '(' {
			need++
		} else if need > 0 {
			need--
		} else {
			count++
		}
	}
	return count + need
}
