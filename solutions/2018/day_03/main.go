package main

import (
	"fmt"
	"sort"

	"github.com/wcharczuk/advent/pkg/fileutil"
	"github.com/wcharczuk/advent/pkg/log"
)

type Claim struct {
	ID                  int
	X, Y, Width, Height int
}

func (c Claim) Area() int {
	return (c.B() - c.T()) * (c.R() - c.L())
}

func (c Claim) T() int {
	return c.Y
}

func (c Claim) B() int {
	return c.Y + c.Height
}

func (c Claim) L() int {
	return c.X
}

func (c Claim) R() int {
	return c.X + c.Width
}

func (c Claim) Intersects(other Claim) bool {
	return !(c.L() > other.R() || other.L() > c.R() || other.B() > c.T() || c.B() > other.T())
}

func (c Claim) Overlap(other Claim) int {
	if !c.Intersects(other) {
		return 0
	}

	horiz := []int{c.L(), c.R(), other.R(), other.L()}
	sort.Ints(horiz)

	vert := []int{c.T(), c.B(), other.T(), other.B()}
	sort.Ints(vert)

	return Claim{X: horiz[1], Y: vert[1], Width: horiz[2] - horiz[1], Height: vert[2] - vert[1]}.Area()
}

func (c Claim) String() string {
	return fmt.Sprintf("[%d, %d, %d, %d]", c.T(), c.B(), c.L(), c.R())
}

func main() {
	var sheet [][]int

	var claims []Claim

	var mx, my int

	err := fileutil.ReadByLines("./input", func(line string) error {
		var claim Claim
		_, err := fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &claim.ID, &claim.X, &claim.Y, &claim.Width, &claim.Height)
		if err != nil {
			return err
		}
		if mx < claim.R() {
			mx = claim.R()
		}

		if my < claim.B() {
			my = claim.B()
		}

		claims = append(claims, claim)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	sheet = make([][]int, my)
	for index := 0; index < my; index++ {
		sheet[index] = make([]int, mx)
	}

	for _, claim := range claims {
		for y := claim.T(); y < claim.B(); y++ {
			for x := claim.L(); x < claim.R(); x++ {
				sheet[y][x] = sheet[y][x] + 1
			}
		}
	}

	var total int
	for y := 0; y < len(sheet); y++ {
		for x := 0; x < len(sheet[y]); x++ {
			if sheet[y][x] > 1 {
				total = total + 1
			}
		}
	}

	log.Solution(total)
}
