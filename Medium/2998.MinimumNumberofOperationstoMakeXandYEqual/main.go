package main

import "slices"

func minimumOperationsToMakeEqual(x int, y int) int {
	if x <= y {
		return y - x
	}

	ops := make([]int, x+1, x+1)

	for i := 0; i < y; i++ {
		ops[i] = y - i
	}

	ops[y] = 0

	tmp := make([]int, 3, 3)
	for i := y + 1; i <= x; i++ {
		rem := i % 11 // divide by 11
		if rem < 6 {
			tmp[0] = ops[i/11] + rem + 1 // decrement rem times
		} else {
			tmp[0] = ops[(i/11)+1] + (11 - rem) + 1 // increment 11-rem times
		}

		rem = i % 5 // divide by 5
		if rem < 3 {
			tmp[1] = ops[i/5] + rem + 1 // decrement rem times
		} else {
			tmp[1] = ops[(i/5)+1] + (5 - rem) + 1 // increment 5-rem times
		}
		tmp[2] = i - y
		ops[i] = slices.Min(tmp)
	}

	return ops[x]
}

func main() {
	minimumOperationsToMakeEqual(54, 2)
}
