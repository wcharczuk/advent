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

	//raw = []byte("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2")

	contents := strings.Split(string(raw), " ")

	var index, nodeID int
	root := visit(&index, &nodeID, contents)
	log.Solution(root.Sum())
}

func visit(index, nodeID *int, contents []string) node {
	var node node
	node.ID = *nodeID
	var state, numChildren, numMeta int
	for ; *index < len(contents); *index = *index + 1 {
		value := contents[*index]
		switch state {
		case 0:
			numChildren = mustInt(value)
			state = 1
			continue
		case 1:
			numMeta = mustInt(value)
			if numChildren == 0 {
				println(node.ID, "has no children")
				state = 3
			} else {
				state = 2
			}
			continue
		case 2:
			*nodeID = *nodeID + 1
			node.Children = append(node.Children, visit(index, nodeID, contents))
			if len(node.Children) == numChildren {
				state = 3
			}
			continue
		case 3:
			node.Metadata = append(node.Metadata, mustInt(value))
			if len(node.Metadata) == numMeta {
				return node
			}
			continue
		}
	}
	return node
}

func mustInt(value string) int {
	parsed, _ := strconv.Atoi(value)
	return parsed
}

type node struct {
	ID       int
	Children []node
	Metadata []int
}

func (n node) Sum() int {
	if len(n.Children) == 0 {
		leafValue := mathutil.SumInts(n.Metadata...)
		println(n.ID, "node value", leafValue)
		return leafValue
	}

	var output int
	for _, index := range n.Metadata {
		if index <= len(n.Children) {
			child := n.Children[index-1]
			childSum := child.Sum()
			println(n.ID, "processing child", child.ID, childSum)
			output = output + childSum
		} else {
			println(n.ID, "skipping index", index-1, len(n.Children))
		}
	}
	return output
}
