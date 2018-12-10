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
		//{Players: 9, LastMarble: 25, Expected: 32},
		//{Players: 10, LastMarble: 1618, Expected: 8317},
		//{Players: 13, LastMarble: 7999, Expected: 146373},
		//{Players: 17, LastMarble: 1104, Expected: 2764},
		//{Players: 21, LastMarble: 6111, Expected: 54718},
		//{Players: 30, LastMarble: 5807, Expected: 37305},
		{Players: 464, LastMarble: 71730 * 100, Expected: 0},
	}

	for _, tc := range testCases {
		actual := simulate(tc.Players, tc.LastMarble)
		log.Solution(actual)
	}
}

// simulate runs the simulation
func simulate(players, lastMarble int) int {
	current := NewRing()
	scores := make([]int, players)
	var player, removed int
	for next := 1; next <= lastMarble; next++ {
		if next > 0 && next%23 == 0 {
			removed, current = current.Remove()
			scores[player] = scores[player] + next + removed
		} else {
			current = current.Add(next)
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

func NewRing() *RingNode {
	r := &RingNode{Value: 0}
	r.CW = r
	r.CCW = r
	return r
}

type RingNode struct {
	Value   int
	CW, CCW *RingNode
}

func (r *RingNode) Add(value int) *RingNode {
	if value == 0 {
		return r
	}

	newNode := &RingNode{Value: value}
	if value == 1 {
		r.CCW = newNode
		r.CW = newNode
		newNode.CCW = r
		newNode.CW = r
		return newNode
	}

	ccw := r.CW
	cw := ccw.CW

	ccw.CW = newNode
	newNode.CCW = ccw
	newNode.CW = cw
	cw.CCW = newNode

	return newNode
}

func (r *RingNode) Remove() (int, *RingNode) {
	cursor := r
	for x := 0; x < 7; x++ {
		cursor = cursor.CCW
	}

	value := cursor.Value
	ccw := cursor.CCW
	cw := cursor.CW
	ccw.CW = cw
	cw.CCW = ccw
	return value, cw
}

func (r *RingNode) String() string {
	cursor := r
	for cursor.Value != 0 {
		cursor = cursor.CCW
	}

	var output []string
	if cursor.Value == r.Value {
		output = append(output, fmt.Sprintf("(%d)", cursor.Value))
	} else {
		output = append(output, strconv.Itoa(cursor.Value))
	}
	cursor = cursor.CW
	for cursor.Value != 0 {
		if cursor.Value == r.Value {
			output = append(output, fmt.Sprintf("(%d)", cursor.Value))
		} else {
			output = append(output, strconv.Itoa(cursor.Value))
		}
		cursor = cursor.CW
	}
	return strings.Join(output, " ")
}
