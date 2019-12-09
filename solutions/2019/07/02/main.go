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

	var max, trial int64
	for _, combination := range generatePhaseSettings() {
		trial = runExperiment(program, combination)
		if trial > max {
			max = trial
		}
	}
	fmt.Println("Test Run", max)
}

func generatePhaseSettings() [][5]int64 {
	return generatePhaseSettingsImpl([5]int64{}, 0)
}

func generatePhaseSettingsImpl(working [5]int64, index int) [][5]int64 {
	if index == 5 {
		return [][5]int64{working}
	}

	hasValue := func(value int64) bool {
		for y := 0; y < index; y++ {
			if working[y] == value {
				return true
			}
		}
		return false
	}

	var output [][5]int64
	for x := 0; x < 5; x++ {
		if hasValue(int64(x) + 5) {
			continue
		}
		working[index] = int64(x) + 5
		output = append(output, generatePhaseSettingsImpl(working, index+1)...)
	}
	return output
}

func link(from, to chan int64) {
	for {
		select {
		case value := <-from:
			to <- value
		}
	}
}

func runExperiment(program []int64, phaseSettings [5]int64) int64 {
	bufferSize := 128

	amplifiers := [...]*intcode.Computer{
		intcode.New(program, intcode.OptName("A"), intcode.OptDebug(true)),
		intcode.New(program, intcode.OptName("B"), intcode.OptDebug(true)),
		intcode.New(program, intcode.OptName("C"), intcode.OptDebug(true)),
		intcode.New(program, intcode.OptName("D"), intcode.OptDebug(true)),
		intcode.New(program, intcode.OptName("E"), intcode.OptDebug(true)),
	}

	outputs := make([]chan int64, 5)
	for x := 0; x < 5; x++ {
		outputs[x] = make(chan int64, bufferSize)
		amplifiers[x].OutputHandlers = append(amplifiers[x].OutputHandlers, func(index int) func(int64) {
			return func(v int64) {
				outputs[index] <- v
			}
		}(x))
	}

	var last int64
	amplifiers[4].OutputHandlers = append(amplifiers[4].OutputHandlers, func(value int64) {
		last = value
	})

	inputs := make([]chan int64, 5)
	for x := 0; x < 5; x++ {
		inputs[x] = make(chan int64, bufferSize)
		amplifiers[x].InputHandler = func(index int) func() int64 {
			return func() int64 {
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
