package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Reindeer struct {
	Name        string
	FlyingSpeed int
	FlyingTime  int
	RestingTime int
}

func parseEntry(input string) *Reindeer {
	r := Reindeer{}
	inputParts := strings.Split(input, " ")
	r.Name = inputParts[0]

	speedStr := inputParts[3]
	timeStr := inputParts[6]
	restingStr := inputParts[13]

	speed, _ := strconv.Atoi(speedStr)
	flyingTime, _ := strconv.Atoi(timeStr)
	restingTime, _ := strconv.Atoi(restingStr)

	r.FlyingSpeed = speed
	r.FlyingTime = flyingTime
	r.RestingTime = restingTime

	fmt.Printf("reindeer: %#v\n", r)
	return &r
}

func main() {
	dataFile := "../testdata/day14"
	simulationTime := 2503

	deer := map[string]*Reindeer{}
	if f, err := os.Open(dataFile); err == nil {
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			entry := scanner.Text()
			r := parseEntry(entry)
			deer[r.Name] = r
		}
	}

	progress := map[string]int{}
	inAir := map[string]int{}
	grounded := map[string]int{}

	toGround := []string{}
	toPutInAir := []string{}

	//put all the deer in the air
	for name, _ := range deer {
		r := deer[name]
		inAir[r.Name] = 0
	}

	time := 0
	tick := 1
	for time < simulationTime {
		//advance in air deer
		for name, _ := range inAir {
			ref := deer[name]
			inAir[name] = inAir[name] + tick

			progress[name] = progress[name] + (tick * ref.FlyingSpeed)
			if inAir[name] >= ref.FlyingTime {
				toGround = append(toGround, name)
			}
		}

		//advance grounded deer
		for name, timeOnGround := range grounded {
			ref := deer[name]
			grounded[name] = timeOnGround + tick
			if grounded[name] >= ref.RestingTime {
				toPutInAir = append(toPutInAir, name)
			}
		}

		//ground deer
		for _, name := range toGround {
			delete(inAir, name)
			grounded[name] = 0
		}
		toGround = []string{}

		for _, name := range toPutInAir {
			delete(grounded, name)
			inAir[name] = 0
		}
		toPutInAir = []string{}

		time = time + tick
	}

	bestDistance := 0
	bestDeer := ""
	for name, distance := range progress {
		if distance > bestDistance {
			bestDistance = distance
			bestDeer = name
		}
	}

	fmt.Printf("%s travelled %d\n", bestDeer, bestDistance)
}
