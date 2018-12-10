package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wcharczuk/advent/pkg/log"
)

const (
	input      = "464 players; last marble is worth 71730 points"
	players    = 464
	lastMarble = 71730
)

type testCase struct {
	Players    int
	LastMarble int
	Expected   int
}

func main() {
	testCases := []testCase{
		{Players: 9, LastMarble: 25, Expected: 32},
		{Players: 10, LastMarble: 1618, Expected: 8317},
		{Players: 13, LastMarble: 7999, Expected: 146373},
		{Players: 17, LastMarble: 1104, Expected: 2764},
		{Players: 21, LastMarble: 6111, Expected: 54718},
		{Players: 30, LastMarble: 5807, Expected: 37305},
		{Players: 464, LastMarble: 71730, Expected: 0},
	}

	for _, tc := range testCases {
		actual := simulate(tc.Players, tc.LastMarble)
		if actual != tc.Expected {
			log.Context("error").Printf("expected: %d, actual: %d", tc.Expected, actual)
		} else {
			log.Context("success").Printf("players: %d, last: %d, highest: %d", tc.Players, tc.LastMarble, actual)
		}
	}
}

// simulate runs the simulation
func simulate(players, lastMarble int) int {
	board := NewRing(0)
	scores := make([]int, players)
	var player int
	for next := 1; next <= lastMarble; next++ {
		if next > 0 && next%23 == 0 {
			removed := board.Remove()
			scores[player] = scores[player] + next + removed
		} else {
			board.Add(next)
		}
		player = (player + 1) % players
	}

	var highestScore int
	for _, score := range scores {
		if highestScore < score {
			highestScore = score
		}
	}
	return highestScore
}

func NewRing(values ...int) *Ring {
	return &Ring{
		Values: values,
	}
}

type Ring struct {
	Cursor int
	Values []int
}

func (r *Ring) Add(value int) {
	r.Cursor = cw(r.Cursor, len(r.Values))
	var before, after []int
	if r.Cursor < len(r.Values) {
		before = r.Values[:r.Cursor]
		after = r.Values[r.Cursor:]
	} else {
		before = r.Values
	}
	r.Values = append(before, append([]int{value}, after...)...)
	return
}

func (r *Ring) Remove() int {
	r.Cursor = ccw(r.Cursor, len(r.Values), 7)
	output := r.Values[r.Cursor]
	r.Values = removeAt(r.Values, r.Cursor)
	return output
}

func (r Ring) String() string {
	var values []string
	for index, value := range r.Values {
		if index == r.Cursor {
			values = append(values, fmt.Sprintf("(%d)", value))
		} else {
			values = append(values, strconv.Itoa(value))
		}
	}
	return strings.Join(values, " ")
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func cw(cursor, length int) int {
	if length == 0 {
		return 0
	} else if length == 1 {
		return 1
	} else if length == 2 {
		return 1
	} else if length == 3 {
		return 3
	} else if length == 4 {
		return 1
	}
	return ((cursor + 1) % length) + 1
}

func ccw(cursor, length, change int) int {
	intermediate := cursor - change
	for intermediate < 0 {
		intermediate = intermediate + length
	}
	return intermediate % length
}

func removeAt(values []int, index int) []int {
	if index == len(values)-1 {
		return values[:index]
	}
	before := values[:index]
	after := values[index+1:]
	return append(before, after...)
}
