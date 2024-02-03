package main

func sequentialDigits(low int, high int) []int {
	// low = 100, high = 300

	adigit, sum, mindigit, ans, start, c := 1, 0, 0, []int{}, 1, low
	for {
		c /= 10
		if c == 0 {
			break
		}
		mindigit++
	}
	for i := 0; i < mindigit; i++ {
		adigit = adigit * 10
	}

	for {
		sum = 0
		digit := adigit
		for i := start; ; i++ {
			// 0 + 100 * 1 = 100
			// 100 + 10 * 2 = 120
			// 120 + 1 * 3 = 123
			sum += digit * i
			digit /= 10

			if digit == 0 {
				if i == 9 {
					adigit *= 10
					start = 0
				}
				break
			}
		}

		if sum > high {
			return ans
		} else if sum >= low {
			ans = append(ans, sum)
		}
		start++
	}
}

func main() {
	println(sequentialDigits(58, 155))
}
