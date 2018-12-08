package main

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/wcharczuk/advent/pkg/mathutil"

	"github.com/wcharczuk/advent/pkg/log"
)

func main() {

	raw, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	// raw := []byte("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2")

	contents := strings.Split(string(raw), " ")

	var index int
	root := visit(&index, contents)
	log.Solution(root.Sum())
}

func visit(index *int, contents []string) node {
	var node node
	var state, numChildren, numMeta int
	for ; *index < len(contents); *index = *index + 1 {
		value := contents[*index]
		println(value)
		switch state {
		case 0:
			numChildren = mustInt(value)
			state = 1
		case 1:
			numMeta = mustInt(value)
			if numChildren > 0 {
				state = 2
			} else {
				state = 3
			}
		case 2:
			node.Children = append(node.Children, visit(index, contents))
			if len(node.Children) == numChildren {
				state = 3
			}
		case 3:
			node.Metadata = append(node.Metadata, mustInt(value))
			if len(node.Metadata) == numMeta {
				return node
			}
		}
	}
	return node
}

func mustInt(value string) int {
	parsed, _ := strconv.Atoi(value)
	return parsed
}

type node struct {
	Children []node
	Metadata []int
}

func (n node) Sum() int {
	output := mathutil.SumInts(n.Metadata...)
	for _, child := range n.Children {
		output += child.Sum()
	}
	return output
}
