package main

func findTheDifference(s string, t string) byte {
	var sums byte = 0
	var sumt byte = 0

	for i := 0; i < len(t); i++ {
		sumt += t[i]
		if i < len(s) {
			sums += s[i]
		}
	}

	return sumt - sums
}
