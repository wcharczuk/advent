package main

import (
	"fmt"
	"strings"

	"github.com/wcharczuk/advent/pkg/log"
)

/*
const (
	initialState = "#..#.#..##......###...###"
	inputRules   = `...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #`
)
*/

const (
	initialState = "###..#...####.#..###.....####.######.....##.#####.##.##..###....#....##...##...##.#..###..#.#...#..#"
	inputRules   = `.###. => .
..#.. => .
.#### => .
.##.. => #
#.#.# => .
..#.# => #
#.##. => #
#...# => #
..... => .
##..# => #
.#.#. => .
..##. => #
##.#. => .
###.. => .
.#... => #
..### => .
#..## => .
...#. => .
###.# => #
.##.# => .
.#.## => .
....# => .
##### => .
#.#.. => #
...## => #
#.... => .
#.### => #
##... => #
.#..# => .
####. => .
#..#. => #
##.## => #`
)

const (
	maxGen = 50000000000
)

func main() {
	rules := ParseRules(inputRules)
	row := ParsePlantRow(initialState)

	var gen, current, previous, change, previousChange int
	for ; gen < maxGen; gen++ {
		current, _ = row.Apply(rules)
		change = current - previous
		if change == previousChange && gen > 50 {
			break
		}
		previous = current
		previousChange = change
	}

	generationsLeft := maxGen - (gen + 1)
	log.Solution(current + (generationsLeft * change))
}

func ParseRules(rawRules string) []Rule {
	lines := strings.Split(rawRules, "\n")
	var rules []Rule
	for _, line := range lines {
		rules = append(rules, ParseRule(line))
	}
	return rules
}

func ParseRule(line string) Rule {
	runes := []rune(line)
	return Rule{
		L1: runes[0] == '#',
		L0: runes[1] == '#',
		C:  runes[2] == '#',
		R0: runes[3] == '#',
		R1: runes[4] == '#',
		N:  runes[9] == '#',
	}
}

type Rule struct {
	L0, L1 bool
	C      bool
	R0, R1 bool
	N      bool
}

func (r Rule) String() string {
	output := make([]string, 6)
	if r.L1 {
		output[0] = "#"
	} else {
		output[0] = "."
	}
	if r.L0 {
		output[1] = "#"
	} else {
		output[1] = "."
	}
	if r.C {
		output[2] = "#"
	} else {
		output[2] = "."
	}
	if r.R0 {
		output[3] = "#"
	} else {
		output[3] = "."
	}
	if r.R1 {
		output[4] = "#"
	} else {
		output[4] = "."
	}
	if r.N {
		output[5] = "#"
	} else {
		output[5] = "."
	}

	return fmt.Sprintf("%s%s%s%s%s => %s", output[0], output[1], output[2], output[3], output[4], output[5])
}

func (r Rule) HasNext(section []bool) (applies, hasNext bool) {
	if section[0] != r.L1 {
		return
	}
	if section[1] != r.L0 {
		return
	}
	if section[2] != r.C {
		return
	}
	if section[3] != r.R0 {
		return
	}
	if section[4] != r.R1 {
		return
	}

	applies = true
	hasNext = r.N
	return
}

func ParsePlantRow(line string) PlantRow {
	runes := []rune(line)
	plants := make([]bool, len(runes)+10)
	for index, r := range runes {
		if r == '#' {
			plants[index+5] = true
		}
	}

	return PlantRow{
		CenterIndex: 5,
		Plants:      plants,
	}
}

type PlantRow struct {
	CenterIndex int
	Plants      []bool
}

func (pr *PlantRow) Apply(rules []Rule) (sum, flipped int) {
	currentGeneration := copy(pr.Plants)
	pr.Plants = make([]bool, len(currentGeneration))
	for index := 2; index < len(currentGeneration)-2; index++ {
		for _, rule := range rules {
			applies, hasNext := rule.HasNext(currentGeneration[index-2 : index+3])
			if !applies {
				continue
			}
			pr.Plants[index] = hasNext
		}
	}

	for index := 0; index < len(currentGeneration); index++ {
		if pr.Plants[index] {
			sum += (index - pr.CenterIndex)
			flipped++
		}
	}
	pr.Grow()
	return
}

func (pr *PlantRow) Grow() {
	if pr.Plants[0] {
		pr.Plants = expandLeft(pr.Plants, 5)
		pr.CenterIndex += 5
	} else if pr.Plants[1] {
		pr.Plants = expandLeft(pr.Plants, 4)
		pr.CenterIndex += 4
	} else if pr.Plants[2] {
		pr.Plants = expandLeft(pr.Plants, 3)
		pr.CenterIndex += 3
	} else if pr.Plants[3] {
		pr.Plants = expandLeft(pr.Plants, 2)
		pr.CenterIndex += 2
	} else if pr.Plants[4] {
		pr.Plants = expandLeft(pr.Plants, 1)
		pr.CenterIndex++
	}

	plen := len(pr.Plants)
	if pr.Plants[plen-1] {
		pr.Plants = expandRight(pr.Plants, 5)
	} else if pr.Plants[plen-2] {
		pr.Plants = expandRight(pr.Plants, 4)
	} else if pr.Plants[plen-3] {
		pr.Plants = expandRight(pr.Plants, 3)
	} else if pr.Plants[plen-4] {
		pr.Plants = expandRight(pr.Plants, 2)
	} else if pr.Plants[plen-5] {
		pr.Plants = expandRight(pr.Plants, 1)
	}
}

func (pr PlantRow) String() string {
	output := make([]string, len(pr.Plants))
	for index, p := range pr.Plants {
		if p {
			output[index] = "#"
		} else {
			output[index] = "."
		}
		if index == pr.CenterIndex {
			output[index] = "(" + output[index] + ")"
		}
	}
	return strings.Join(output, " ")
}

func (pr PlantRow) Score() int {
	var output int
	for index := 0; index < len(pr.Plants); index++ {
		if pr.Plants[index] {
			output = output + (index - pr.CenterIndex)
		}
	}
	return output
}

func copy(values []bool) []bool {
	output := make([]bool, len(values))
	for x := 0; x < len(values); x++ {
		output[x] = values[x]
	}
	return output
}

func expandLeft(values []bool, expandBy int) []bool {
	output := make([]bool, len(values)+expandBy)
	for x := 0; x < len(values)-expandBy; x++ {
		output[x+expandBy] = values[x]
	}
	return output
}

func expandRight(values []bool, expandBy int) []bool {
	output := make([]bool, len(values)+expandBy)
	for x := 0; x < len(values)-expandBy; x++ {
		output[x] = values[x]
	}
	return output
}
