package main

func canVisitAllRooms(rooms [][]int) bool {
	visit := make([]bool, len(rooms))

	q := []int{0}

	for len(q) > 0 {
		f := q[0]
		q = q[1:]
		visit[f] = true
		for _, v := range rooms[f] {
			if visit[v] {
				continue
			}
			q = append(q, v)
		}
	}

	for i := 0; i < len(rooms); i++ {
		if !visit[i] {
			return false
		}
	}
	return true
}

func main() {
	rooms := [][]int{
		{1, 3}, {3, 0, 1}, {2}, {0},
	}
	canVisitAllRooms(rooms)
}
