package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wcharczuk/advent/pkg/intcode"
)

func main() {
	program, err := intcode.ReadProgramFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	computer := intcode.New(program,
		intcode.OptName("day11"),
		intcode.OptDebug(true),
		intcode.OptDebugLog(os.Stdout),
	)

	input := make(chan int64, 32)
	output := make(chan int64, 32)

	computer.InputHandler = func() int64 {
		return <-input
	}
	computer.OutputHandlers = append(computer.OutputHandlers, func(value int64) {
		output <- value
	})

	board := make(Grid)
	robot := NewRobot()

	done := make(chan struct{})

	go func() {
		defer close(done)
		if err := computer.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	func() {
		for {
			input <- board.Get(robot.Pos.X, robot.Pos.Y)
			select {
			case <-done:
				return
			case color := <-output:
				robot.Paint(board, color)
				robot.Turn(<-output)
				robot.Move()
			}
		}
	}()

	fmt.Printf("ANSWER: %d\n", len(board))
}

// NewRobot returns a new robot.
func NewRobot() *Robot {
	return &Robot{
		Dir: Vect{0, 1},
	}
}

// Robot is the state machine we're going to use to
type Robot struct {
	Pos Coord
	Dir Vect
}

// Paint paints the grid with a given value at the
// robot's current position.
func (r *Robot) Paint(grid Grid, v int64) {
	grid[r.Pos] = v == 1
}

// Turn tuns the vector.
func (r *Robot) Turn(d int64) error {
	switch d {
	case 0:
		r.Dir = r.Dir.TurnLeft()
		return nil
	case 1:
		r.Dir = r.Dir.TurnRight()
		return nil
	default:
		return fmt.Errorf("invalid turn direction: %d", d)
	}
}

// Move applies the vector to the position.
func (r *Robot) Move() {
	r.Pos.Y = r.Pos.Y + r.Dir.Y
	r.Pos.X = r.Pos.X + r.Dir.X
}

// Grid is a map of corrdinates.
type Grid map[Coord]bool

// Get gets a value at a given coordinate.
// It returns false if there is no corresponding coordinate.
func (g Grid) Get(x, y int) int64 {
	if value, ok := g[Coord{x, y}]; ok {
		if value {
			return 1
		}
		return 0
	}
	return 0
}

// Coord is an X,Y position.
type Coord struct {
	X, Y int
}

// Vect is a unit distance (0,1), (0,-1), (1,0), (-1,0)
type Vect struct {
	X, Y int
}

// TurnLeft rotates the vector by 90 ccw.
func (v Vect) TurnLeft() Vect {
	if v.X == 0 {
		if v.Y == 1 {
			return Vect{-1, 0}
		}
		if v.Y == -1 {
			return Vect{1, 0}
		}
	}
	if v.X == -1 {
		return Vect{0, -1}
	}
	return Vect{0, 1}
}

// TurnRight rotates the vector by 90 cw.
func (v Vect) TurnRight() Vect {
	if v.X == 0 {
		if v.Y == 1 {
			return Vect{1, 0}
		}
		if v.Y == -1 {
			return Vect{-1, 0}
		}
	}
	if v.X == -1 {
		return Vect{0, 1}
	}
	return Vect{0, -1}
}
