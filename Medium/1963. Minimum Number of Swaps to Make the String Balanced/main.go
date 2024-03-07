package main

func minSwaps(s string) int {
	op, cl := 0, 0
	for _, c := range s {
		if c == '[' {
			op++
		} else if op > 0 {
			op--
		} else {
			cl++
		}
	}

	return (cl + 1) / 2
}
