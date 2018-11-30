package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/blend/go-sdk/util"
)

func test() {
	i1 := Ingredient{
		Name:       "Butterscotch",
		Capacity:   -1,
		Durability: -2,
		Flavor:     6,
		Texture:    3,
		Calories:   8,
	}

	i2 := Ingredient{
		Name:       "Cinnamon",
		Capacity:   2,
		Durability: 3,
		Flavor:     -2,
		Texture:    -1,
		Calories:   3,
	}

	ingredients := []Ingredient{i1, i2}
	distribution := []int{44, 56}

	score, _ := calculateScore(distribution, ingredients)
	if score != 62842880 {
		fmt.Println("Test :: Scores Wrong!", score, 62842880)
		os.Exit(1)
	}
}

type Ingredient struct {
	Name       string
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

func parseEntry(input string) Ingredient {
	inputParts := strings.Split(input, " ")
	i := Ingredient{}
	i.Name = strings.Replace(inputParts[0], ":", "", 1)
	i.Capacity, _ = strconv.Atoi(strings.Replace(inputParts[2], ",", "", 1))
	i.Durability, _ = strconv.Atoi(strings.Replace(inputParts[4], ",", "", 1))
	i.Flavor, _ = strconv.Atoi(strings.Replace(inputParts[6], ",", "", 1))
	i.Texture, _ = strconv.Atoi(strings.Replace(inputParts[8], ",", "", 1))
	i.Calories, _ = strconv.Atoi(inputParts[10])
	return i
}

func main() {
	test()

	codeFile := "../testdata/day15"
	ingredients := []Ingredient{}
	util.File.ReadByLines(codeFile, func(line string) error {
		i := parseEntry(line)
		ingredients = append(ingredients, i)
		return nil
	})

	distributions := permuteDistributions(100, len(ingredients))

	bestScore := 0
	bestDistribution := []int{}
	for _, distribution := range distributions {
		score, calories := calculateScore(distribution, ingredients)
		if calories == 500 && score > bestScore {
			bestScore = score
			bestDistribution = make([]int, len(ingredients))
			copy(bestDistribution, distribution)
		}
	}

	fmt.Println("Best Score:", bestScore)
	fmt.Println("Distribution:", fmt.Sprintf("%#v", bestDistribution))
}

func calculateScore(distribution []int, ingredients []Ingredient) (int, int) {
	capacity := 0
	durability := 0
	flavor := 0
	texture := 0
	calories := 0

	for index, value := range distribution {
		i := ingredients[index]

		capacity = capacity + (value * i.Capacity)
		durability = durability + (value * i.Durability)
		flavor = flavor + (value * i.Flavor)
		texture = texture + (value * i.Texture)
		calories = calories + (value * i.Calories)
	}

	subTotal := util.Ternary.OfInt(capacity > 0, capacity, 0) *
		util.Ternary.OfInt(durability > 0, durability, 0) *
		util.Ternary.OfInt(flavor > 0, flavor, 0) *
		util.Ternary.OfInt(texture > 0, texture, 0)
	return subTotal, calories
}

func permuteDistributions(total, buckets int) [][]int {
	return permuteDistributionsFromExisting(total, buckets, []int{})
}

func permuteDistributionsFromExisting(total, buckets int, existing []int) [][]int {
	output := [][]int{}
	existingLength := len(existing)
	existingSum := sum(existing)
	remainder := total - existingSum

	if buckets == 1 {
		newExisting := make([]int, existingLength+1)
		copy(newExisting, existing)
		newExisting[existingLength] = remainder
		output = append(output, newExisting)
		return output
	}

	for x := 0; x <= remainder; x++ {
		newExisting := make([]int, existingLength+1)
		copy(newExisting, existing)
		newExisting[existingLength] = x

		results := permuteDistributionsFromExisting(total, buckets-1, newExisting)
		output = append(output, results...)
	}

	return output
}

func sum(values []int) int {
	total := 0
	for x := 0; x < len(values); x++ {
		total += values[x]
	}

	return total
}
