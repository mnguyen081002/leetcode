package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	q := []int{}
	degre := make([]int, numCourses)
	graph := make([][]int, numCourses)
	ans := []int{}
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		degre[p[0]]++
	}

	for i := 0; i < numCourses; i++ {
		if degre[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		f := q[0]
		q = q[1:]
		numCourses--
		ans = append(ans, f)
		for _, e := range graph[f] {
			degre[e]--
			if degre[e] == 0 {
				q = append(q, e)
			}
		}
	}

	if numCourses != 0 {
		return []int{}
	}
	return ans
}
