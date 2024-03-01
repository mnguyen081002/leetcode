package main

func colorTheArray(n int, queries [][]int) []int {
	count := 0
	arr := make([]int, n)
	ans := make([]int, len(queries))
	for i, v := range queries {
		if v[1] == arr[v[0]] {
			ans[i] = count
			continue
		}
		if v[0]+1 < n && arr[v[0]+1] != 0 {
			// nếu màu bên phải giống
			if arr[v[0]+1] == v[1] {
				count++
				// nếu màu hiện tại đang bị đổi giống màu bên phải thì giảm
			} else if arr[v[0]] == arr[v[0]+1] {
				count--
			}
		}
		// nếu màu bên trái giống
		if v[0]-1 >= 0 && arr[v[0]-1] != 0 {
			if arr[v[0]-1] == v[1] {
				count++
			} else if arr[v[0]] == arr[v[0]-1] {
				count--
			}
		}
		ans[i] = count
		arr[v[0]] = v[1]
	}

	return ans
}

// 2 0 0 0 0 0 - 0
// 2 2 0 0 0 0 - 1
// 2 2 0 1 0 0 - 1
// 2 2 1 1 0 0 - 2
// 2 2 1 1 1 0
// 2 2 1 1 1 1
