package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/wcharczuk/advent/pkg/collections"
	"github.com/wcharczuk/advent/pkg/fileutil"
	"github.com/wcharczuk/advent/pkg/log"
)

func main() {
	var instructions []step
	err := fileutil.ReadByLines("./input", func(line string) error {
		var s step
		_, err := fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &s.Before, &s.After)
		if err != nil {
			return err
		}

		instructions = append(instructions, s)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	nodes := map[string]*node{}
	var ok bool
	for _, s := range instructions {
		var before, after *node
		if before, ok = nodes[s.Before]; !ok {
			before = &node{ID: s.Before}
			nodes[s.Before] = before
		}
		if after, ok = nodes[s.After]; !ok {
			after = &node{ID: s.After}
			nodes[s.After] = after
		}

		before.After = append(before.After, after)
		after.Before = append(after.Before, before)
	}

	var roots []*node
	for _, node := range nodes {
		if len(node.Before) == 0 {
			roots = append(roots, node)
		}
	}

	order := process(roots)
	log.Solution(strings.Join(order, ""))
}

func process(roots []*node) []string {
	var output []string
	seen := collections.NewSetOfString()
	working := roots
	sort.Sort(nodes(working))

	for len(working) > 0 {
		found := working[0]
		working = working[1:]

		// check if we've processed all the parents
		if !hasAll(seen, found.Before) {
			working = append(working, found)
			continue
		}

		// don't reprocess nodes
		if seen.Contains(found.ID) {
			continue
		}

		// process the node
		seen.Add(found.ID)
		output = append(output, found.ID)

		for _, after := range found.After {
			working = append(working, after)
		}

		working = distinct(working)
		sort.Sort(nodes(working))
	}
	return output
}

func hasAll(seen collections.SetOfString, nodes []*node) bool {
	if len(nodes) == 0 {
		return true
	}
	for _, n := range nodes {
		if !seen.Contains(n.ID) {
			return false
		}
	}
	return true
}

func distinct(nodes []*node) []*node {
	seen := collections.NewSetOfString()
	var output []*node
	for _, n := range nodes {
		if !seen.Contains(n.ID) {
			output = append(output, n)
			seen.Add(n.ID)
		}
	}
	return output
}

type nodes []*node

func (n nodes) Len() int {
	return len(n)
}

func (n nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n nodes) Less(i, j int) bool {
	return n[i].ID < n[j].ID
}

func (n nodes) String() string {
	var ids []string
	for _, c := range n {
		ids = append(ids, c.ID)
	}
	return strings.Join(ids, "")
}

type step struct {
	Before string
	After  string
}

type node struct {
	ID     string
	Before []*node
	After  []*node
}
