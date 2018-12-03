package main

import (
	"fmt"

	"github.com/wcharczuk/advent/pkg/fileutil"
	"github.com/wcharczuk/advent/pkg/log"
)

type Claim struct {
	ID                  int
	X, Y, Width, Height int
}

func (a Claim) Area() int {
	return (a.B() - a.T()) * (a.R() - a.L())
}

func (a Claim) T() int {
	return a.Y
}

func (a Claim) B() int {
	return a.Y + a.Height
}

func (a Claim) L() int {
	return a.X
}

func (a Claim) R() int {
	return a.X + a.Width
}

func (a Claim) Overlaps(b Claim) bool {
	xOverlap := valueInRange(a.X, b.X, b.X+b.Width) ||
		valueInRange(b.X, a.X, a.X+a.Width)

	yOverlap := valueInRange(a.Y, b.Y, b.Y+b.Height) ||
		valueInRange(b.Y, a.Y, a.Y+a.Height)

	return xOverlap && yOverlap
}

func (a Claim) String() string {
	return fmt.Sprintf("%d: [%d, %d, %d, %d]", a.ID, a.L(), a.R(), a.T(), a.B())
}

func valueInRange(value, min, max int) bool {
	return (value >= min) && (value < max)
}

func main() {
	var claims []Claim
	err := fileutil.ReadByLines("./input", func(line string) error {
		var claim Claim
		_, err := fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &claim.ID, &claim.X, &claim.Y, &claim.Width, &claim.Height)
		if err != nil {
			return err
		}
		claims = append(claims, claim)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	for index, claim := range claims {
		var hasOverlap bool
		for otherIndex, other := range claims {
			if index != otherIndex {
				if claim.Overlaps(other) {
					hasOverlap = true
					break
				}
			}
		}
		if !hasOverlap {
			log.Solution(claim.ID)
		}
	}
	log.Fatal("no solution found")
}
