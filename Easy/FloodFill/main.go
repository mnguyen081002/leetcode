package floodfill

import (
	"strconv"
)

type Point struct {
	Sc   int
	Sr   int
	Val  int
	Done bool
}

func FloodFillLoop(image [][]int, sr int, sc int, color int) [][]int {
	c := image[sr][sc]
	stack := []*Point{
		{Val: c, Sc: sc, Sr: sr},
	}
	m := map[string]*Point{}

	for len(stack) > 0 {
		s := stack[0]

		if _, ok := m[getKey(s.Sc, s.Sr)]; ok && s.Done {
			stack = stack[1:]
			continue
		}
		if c == s.Val {
			image[s.Sr][s.Sc] = color
		}

		_, ok := m[getKey(s.Sc-1, s.Sr)] // Left
		if s.Sc-1 >= 0 && image[s.Sr][s.Sc-1] == c && !ok {
			i := &Point{Sc: s.Sc - 1, Sr: s.Sr, Val: image[s.Sr][s.Sc-1]}
			stack = append(stack, i)
			m[getKey(i.Sc, i.Sr)] = i
		}

		_, ok = m[getKey(s.Sc+1, s.Sr)] // Right
		if s.Sc+1 < len(image[0]) && image[s.Sr][s.Sc+1] == c && !ok {
			i := &Point{Sc: s.Sc + 1, Sr: s.Sr, Val: image[s.Sr][s.Sc+1]}
			stack = append(stack, i)
			m[getKey(i.Sc, i.Sr)] = i
		}

		_, ok = m[getKey(s.Sc, s.Sr-1)] // Top
		if s.Sr-1 >= 0 && image[s.Sr-1][s.Sc] == c && !ok {
			i := &Point{Sc: s.Sc, Sr: s.Sr - 1, Val: image[s.Sr-1][s.Sc]}
			stack = append(stack, i)
			m[getKey(i.Sc, i.Sr)] = i
		}

		_, ok = m[getKey(s.Sc, s.Sr+1)] // Bottom
		if s.Sr+1 < len(image) && image[s.Sr+1][s.Sc] == c && !ok {
			i := &Point{Sc: s.Sc, Sr: s.Sr + 1, Val: image[s.Sr+1][s.Sc]}
			stack = append(stack, i)
			m[getKey(i.Sc, i.Sr)] = i
		}

		s.Done = true
		m[getKey(s.Sc, s.Sr)] = s
		stack = stack[1:]
	}
	return image
}

func getKey(rc int, rr int) string {
	return strconv.Itoa(rc) + ":" + strconv.Itoa(rr)
}

func FloodFillRecursive(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}

	fill(image, sr, sc, image[sr][sc], color)
	return image
}

func fill(image [][]int, sr int, sc int, color, newColor int) [][]int {
	if sr >= len(image) || sr < 0 || sc >= len(image[0]) || sc < 0 {
		return image
	}
	if image[sr][sc] != color {
		return image
	}

	image[sr][sc] = newColor

	fill(image, sr, sc-1, color, newColor) // up
	fill(image, sr, sc+1, color, newColor) // down
	fill(image, sr-1, sc, color, newColor) // left
	fill(image, sr+1, sc, color, newColor) // right

	return image
}
