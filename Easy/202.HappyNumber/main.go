package main

func isHappy(n int) bool {
	slow, fast := n, n

	for {
		slow = digitSquareSum(slow)
		fast = digitSquareSum(fast)
		fast = digitSquareSum(fast)
		if slow == fast {
			break
		}
	}

	return fast == 1
}

func digitSquareSum(n int) int {
	sum := 0

	for n != 0 {
		tmp := n % 10
		sum += tmp * tmp
		n /= 10
	}

	return sum
}

func main() {
	n := 19
	isHappy(n)
}
