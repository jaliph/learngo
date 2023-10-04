package maths

import "fmt"

func IsRobotBounded(instructions string) bool {
	dirX, dirY := 0, 1
	x, y := 0, 0

	for _, v := range instructions {
		if v == 'G' {
			x += dirX
			y += dirY
		} else if v == 'L' {
			dirX, dirY = -1*dirY, dirX
		} else {
			dirX, dirY = dirY, -1*dirX
		}
		fmt.Println(x, y, dirX, dirY)
	}

	if (x == 0 && y == 0) || (dirX != 0 || dirY != 1) {
		return true
	}

	return false
}
