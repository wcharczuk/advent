package main

import (
	"log"
	"os"

	"github.com/wcharczuk/advent/pkg/intcode"
)

func main() {
	program, err := intcode.ReadProgramFile("../input")
	if err != nil {
		log.Fatal(err)
	}
	computer := intcode.New(program,
		intcode.OptName("boost"),
		intcode.OptDebug(false),
		intcode.OptDebugLog(os.Stdout),
	)
	if err = computer.Run(); err != nil {
		log.Fatal(err)
	}
}
