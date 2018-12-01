package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/wcharczuk/advent/pkg/fileutil"
)

type scanString string

func (ss *scanString) Scan(state fmt.ScanState, verb rune) error {
	token, err := state.Token(true, unicode.IsLetter)
	if err != nil {
		return err
	}
	*ss = scanString(token)
	return nil
}

func newNode() *node {
	return &node{
		children: map[string]*node{},
	}
}

type node struct {
	name       string
	weight     int
	childNames []string
	children   map[string]*node
}

func (n *node) height() int {
	var maxHeight int
	for _, n := range n.children {
		childHeight := n.height()
		if childHeight > maxHeight {
			maxHeight = childHeight
		}
	}
	return maxHeight + 1
}

func (n *node) totalWeight() int {
	var total int
	for _, n := range n.children {
		total += n.totalWeight()
	}
	return total + n.weight
}

func main() {
	nodes := map[string]*node{}

	err := fileutil.ReadByLines("./testdata/input", func(line string) error {
		n := newNode()
		var nameAndWeight string
		if strings.Contains(line, "->") {
			parts := strings.SplitN(line, "->", 2)
			nameAndWeight = strings.TrimSpace(parts[0])
			for _, id := range strings.Split(strings.TrimSpace(parts[1]), ",") {
				n.childNames = append(n.childNames, strings.TrimSpace(id))
			}
		} else {
			nameAndWeight = strings.TrimSpace(line)
		}
		var name scanString
		var weight int
		_, err := fmt.Sscanf(nameAndWeight, "%s (%d)", &name, &weight)
		if err != nil {
			return err
		}
		n.name = string(name)
		n.weight = weight
		nodes[n.name] = n
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, n := range nodes {
		for _, childID := range n.childNames {
			if childNode, hasChild := nodes[childID]; hasChild {
				n.children[childID] = childNode
			}
		}
	}

	var nodeHeight, maxHeight int
	var rootNode string
	for _, n := range nodes {
		nodeHeight = n.height()
		if nodeHeight > maxHeight {
			maxHeight = nodeHeight
			rootNode = n.name
		}
	}

	root := nodes[rootNode]
	weightCounts := map[int]int{}
	for _, c := range root.children {
		weight := c.totalWeight()
		weightCounts[weight]++
	}
}
