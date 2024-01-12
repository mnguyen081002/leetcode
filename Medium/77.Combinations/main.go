package main

func combine(n int, k int) [][]int {
	res := [][]int{}
	recursive(k, n, 0, []int{}, &res)
	return res
}

func recursive(k int, n int, i int, cur []int, res *[][]int) {
	if i != 0 {
		cur = append(cur, i)
	}

	if len(cur) == k {
		*res = append(*res, append([]int{}, cur...))
		return
	}
	for y := i + 1; y <= n; y++ {
		recursive(k, n, y, cur, res)
	}
}

func main() {
	combine(5, 4)
}

// k = 2 n = 4
// 1. c = 0; cur = []
// recur(k=2,n=4,i=0,cur=[],r=[])
// i == 0 => false NO DO cur = append(0)
// for 1 -> n
// ->1.1 recur(k=2,n=4,i=1,cur=[],r=[])
// cur = append(1)
// => for 1 > 4
// -> 1.1.1 recur(k=2,n=4,i=2,cur=[1],r=[])
// cur = append(2)
// => len(cur == k) => r append cur[1,2]
// -> 1.1.2 recur(k=2, n=4,i=4,cur=[1])
// cur = append(4)
// => len(cur == k) => r append cur[1,3]
//...
// 1.2 recur(k=2,n=4,i=)
