package main

import (
	"fmt"
	"log"

	"github.com/wcharczuk/advent/pkg/intcode"
)

func main() {
	program, err := intcode.ReadProgramFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	var max, trial int
	for _, combination := range generatePhaseSettings() {
		trial = runExperiment(program, combination)
		if trial > max {
			max = trial
		}
	}

	fmt.Println("Test Run", max)
}

func generatePhaseSettings() [][5]int {
	return generatePhaseSettingsImpl([5]int{}, 0)
}

func generatePhaseSettingsImpl(working [5]int, index int) [][5]int {
	if index == 5 {
		return [][5]int{working}
	}

	hasValue := func(value int) bool {
		for y := 0; y < index; y++ {
			if working[y] == value {
				return true
			}
		}
		return false
	}

	var output [][5]int
	for x := 0; x < 5; x++ {
		if hasValue(x) {
			continue
		}
		working[index] = x
		output = append(output, generatePhaseSettingsImpl(working, index+1)...)
	}
	return output
}

func link(from, to chan int) {
	value := <-from
	to <- value
}

func runExperiment(program []int, phaseSettings [5]int) int {
	amplifiers := [...]*intcode.Computer{
		intcode.New(program, intcode.OptName("A")),
		intcode.New(program, intcode.OptName("B")),
		intcode.New(program, intcode.OptName("C")),
		intcode.New(program, intcode.OptName("D")),
		intcode.New(program, intcode.OptName("E")),
	}

	outputs := make([]chan int, 5)
	for x := 0; x < 5; x++ {
		outputs[x] = make(chan int, 32)
		amplifiers[x].OutputHandler = func(index int) func(int) {
			return func(v int) {
				println("output", index, v)
				outputs[index] <- v
			}
		}(x)
	}

	inputs := make([]chan int, 5)
	for x := 0; x < 5; x++ {
		inputs[x] = make(chan int, 32)

		amplifiers[x].InputHandler = func(index int) func() int {
			return func() int {
				select {
				case v := <-inputs[index]:
					return v
				}
			}
		}(x)

		if x > 0 {
			go link(outputs[x-1], inputs[x])
		}
	}
	for x := 0; x < 5; x++ {
		go func(index int) {
			if err := amplifiers[index].Run(); err != nil {
				log.Fatal(err)
			}
		}(x)
	}

	go func() {
		for x := 0; x < 5; x++ {
			inputs[x] <- phaseSettings[x]
		}

		// seed the first value ...
		inputs[0] <- 0
	}()

	return <-outputs[4]
}
