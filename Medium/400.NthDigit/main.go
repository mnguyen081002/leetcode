package main

// TODO: handle this
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	x, i, num := 0, 1, 1
	for i <= n {
		x = num / 10
		for x != 0 && i <= n {
			x = num / 10
			i++
		}
		i++
		num++
	}

	return x
}
