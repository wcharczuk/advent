package main

import (
	"math"
)

const (
	// INPUT is the starting position.
	INPUT = 368078

	// INPUT = 1024 // should yield 31
)

func main() {
	size := int(math.Ceil(math.Sqrt(float64(INPUT))))

	println("board size is", size, size)

	originTop := size >> 1
	originLeft := size >> 1

	posTop := originTop
	posLeft := originLeft

	direction := 0 // [0, 3]

	maxTop, maxBottom, maxLeft, maxRight := posTop-1, posTop+1, posLeft-1, posLeft+1
	for val := 1; val < INPUT; val++ {
		// move
		switch direction {
		case 0:
			posLeft = posLeft + 1
		case 1:
			posTop = posTop - 1
		case 2:
			posLeft = posLeft - 1
		case 3:
			posTop = posTop + 1
		}

		// at corner
		if (direction == 0 && posLeft == maxRight) ||
			(direction == 1 && posTop == maxTop) ||
			(direction == 2 && posLeft == maxLeft) ||
			(direction == 3 && posTop == maxBottom) {

			// switch direction
			direction = (direction + 1) % 4

			if direction == 0 {
				// grow the square
				maxLeft = maxLeft - 1
				maxRight = maxRight + 1
				maxTop = maxTop - 1
				maxBottom = maxBottom + 1
			}

		}
	}

	println("final", posTop, posLeft)
	println(abs(posTop-originTop) + abs(posLeft-originLeft))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
