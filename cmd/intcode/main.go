package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/wcharczuk/advent/pkg/intcode"
)

var debug = flag.Bool("debug", false, "If we should show debug output")

func main() {
	flag.Parse()

	var program []int
	var err error
	if len(os.Args) == 1 {
		program, err = intcode.Parse(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		program, err = intcode.Parse(f)
		if err != nil {
			log.Fatal(err)
		}
	}
	computer := intcode.New(program, intcode.OptDebug(*debug), intcode.OptDebugLog(os.Stderr))
	computer.OutputHandlers = []func(int){
		func(v int) {
			fmt.Println("Output: ", v)
		},
	}
	if err := computer.Run(); err != nil {
		log.Fatal(err)
	}
}
