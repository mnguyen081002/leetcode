package main

import "slices"

type Robot struct {
	index     int
	health    int
	position  int
	direction byte
}

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	stack := []Robot{}
	robots := make([]Robot, len(positions))
	for i := range positions {
		robots[i] = Robot{i, healths[i], positions[i], directions[i]}
	}

	slices.SortFunc(robots, func(i, j Robot) int {
		return i.position - j.position
	})

	for _, robot := range robots {
		isRemain := true
		for len(stack) > 0 && stack[len(stack)-1].direction == 'R' && robot.direction == 'L' {
			if stack[len(stack)-1].health < robot.health {
				stack = stack[:len(stack)-1]
				robot.health--
			} else if stack[len(stack)-1].health > robot.health {
				stack[len(stack)-1].health--
				isRemain = false
				break
			} else {
				stack = stack[:len(stack)-1]
				isRemain = false
				break
			}
		}

		if isRemain {
			stack = append(stack, robot)
		}
	}
	result := []int{}

	slices.SortFunc(stack, func(i, j Robot) int {
		return i.index - j.index
	})

	for i := range stack {
		result = append(result, stack[i].health)
	}

	return result
}
