package main

import (
	"fmt"
	"log"
	"sync"

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
		if hasValue(x + 5) {
			continue
		}
		working[index] = x + 5
		output = append(output, generatePhaseSettingsImpl(working, index+1)...)
	}
	return output
}

func link(from, to chan int) {
	for {
		select {
		case value := <-from:
			to <- value
		}
	}
}

func runExperiment(program []int, phaseSettings [5]int) int {
	bufferSize := 128

	amplifiers := [...]*intcode.Computer{
		intcode.New(program, intcode.OptName("A"), intcode.OptDebug(true)),
		intcode.New(program, intcode.OptName("B"), intcode.OptDebug(true)),
		intcode.New(program, intcode.OptName("C"), intcode.OptDebug(true)),
		intcode.New(program, intcode.OptName("D"), intcode.OptDebug(true)),
		intcode.New(program, intcode.OptName("E"), intcode.OptDebug(true)),
	}

	outputs := make([]chan int, 5)
	for x := 0; x < 5; x++ {
		outputs[x] = make(chan int, bufferSize)
		amplifiers[x].OutputHandlers = append(amplifiers[x].OutputHandlers, func(index int) func(int) {
			return func(v int) {
				outputs[index] <- v
			}
		}(x))
	}

	var last int
	amplifiers[4].OutputHandlers = append(amplifiers[4].OutputHandlers, func(value int) {
		last = value
	})

	inputs := make([]chan int, 5)
	for x := 0; x < 5; x++ {
		inputs[x] = make(chan int, bufferSize)
		amplifiers[x].InputHandler = func(index int) func() int {
			return func() int {
				return <-inputs[index]
			}
		}(x)

		if x > 0 {
			go link(outputs[x-1], inputs[x])
		}
	}
	go link(outputs[4], inputs[0]) // feedback loop

	go func() {
		for x := 0; x < 5; x++ {
			inputs[x] <- phaseSettings[x]
		}
		// seed the first value as well
		inputs[0] <- 0
	}()

	wg := sync.WaitGroup{}
	wg.Add(5)
	for x := 0; x < 5; x++ {
		go func(index int) {
			defer wg.Done()
			if err := amplifiers[index].Run(); err != nil {
				log.Fatal(err)
			}
		}(x)
	}

	wg.Wait()
	return last
}
