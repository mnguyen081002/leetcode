package main

func LoopChannel(n int) int {
	sum := 0
	c := make(chan int)
	for i := 0; i < n; i++ {
		go func(i int) {
			c <- i
		}(i)
	}
	sum += <-c
	return sum
}
