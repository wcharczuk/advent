package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/wcharczuk/advent/pkg/intcode"
)

func main() {
	contents, err := ioutil.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	rawValues := strings.Split(string(contents), ",")
	opCodes := make([]int, len(rawValues))

	for x := 0; x < len(rawValues); x++ {
		opCodes[x], err = strconv.Atoi(strings.TrimSpace(rawValues[x]))
		if err != nil {
			log.Fatal(err)
		}
	}

	computer := intcode.New(opCodes...)

	err = computer.Run()
	if err != nil {
		log.Fatal(err)
	}
	for _, logItem := range computer.Log {
		fmt.Println(logItem.String())
	}
}
