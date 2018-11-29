package main

import (
	"math"
)

const (
	// INPUT is the starting position.
	INPUT = 368078

	// INPUT = 1024

	// INPUT = 23 // should yield 806
)

func main() {
	size := int(math.Ceil(math.Sqrt(float64(INPUT))))

	board := make([][]uint64, size)
	for row := 0; row < size; row++ {
		board[row] = make([]uint64, size)
	}

	println("board size is", size, size)

	originTop := size >> 1
	originLeft := size >> 1

	posTop := originTop
	posLeft := originLeft

	direction := 0 // [0, 3]

	board[posTop][posLeft] = 1

	maxTop, maxBottom, maxLeft, maxRight := posTop-1, posTop+1, posLeft-1, posLeft+1
	var val uint64
	for val < INPUT {
		val = board[posTop][posLeft]
		// left
		if posLeft > 0 {
			val += board[posTop][posLeft-1]
		}
		// up
		if posTop > 0 {
			val += board[posTop-1][posLeft]
		}
		// right
		if posLeft < (size - 1) {
			val += board[posTop][posLeft+1]
		}
		//down
		if posTop < (size - 1) {
			val += board[posTop+1][posLeft]
		}
		// up left
		if posLeft > 0 && posTop > 0 {
			val += board[posTop-1][posLeft-1]
		}
		// down left
		if posLeft > 0 && posTop < (size-1) {
			val += board[posTop+1][posLeft-1]
		}
		// up right
		if posLeft < (size-1) && posTop > 0 {
			val += board[posTop-1][posLeft+1]
		}
		// downRight
		if posLeft < (size-1) && posTop < (size-1) {
			val += board[posTop+1][posLeft+1]
		}
		board[posTop][posLeft] = val

		if val >= INPUT {
			break
		}

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

	println(board[posTop][posLeft])
}
