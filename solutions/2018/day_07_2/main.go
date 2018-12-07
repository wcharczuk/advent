package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/wcharczuk/advent/pkg/collections"
	"github.com/wcharczuk/advent/pkg/fileutil"
	"github.com/wcharczuk/advent/pkg/log"
)

const (
	file        = "./input"
	basis       = 60
	workerCount = 5
)

/*
const (
	file        = "./input_sample"
	basis       = 0
	workerCount = 2
)*/

func main() {
	var instructions []step
	err := fileutil.ReadByLines(file, func(line string) error {
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

	cost := process(roots)
	log.Solution(cost)
}

func process(roots []*node) int {
	var clock int
	started := collections.NewSetOfString()
	completed := collections.NewSetOfString()
	queue := roots

	workers := newWorkerPool()
	for len(queue) > 0 || anyBusy(workers) {
		queue = distinct(queue)
		sort.Sort(nodes(queue))

		// check if workers are done ...
		for _, worker := range workers {
			if worker.TimeLeft > 0 && worker.Active != nil {
				worker.TimeLeft--
			}
		}

		queueLen := len(queue)
		for index := 0; index < queueLen; index++ {

			available := getAvailable(workers)
			if len(available) == 0 {
				break
			}

			candidate := queue[0]
			queue = queue[1:]

			if !candidate.CanProcess(completed) {
				queue = append(queue, candidate)
				continue
			}

			started.Add(candidate.ID)
			first := available[0]
			first.TimeLeft = candidate.Cost() - 1
			first.Active = candidate

			println(clock, "starting", candidate.ID)
		}

		clock++

		for _, worker := range workers {
			if worker.TimeLeft == 0 && worker.Active != nil {
				println(clock, "completing", worker.Active.ID)
				completed.Add(worker.Active.ID)
				for _, child := range worker.Active.After {
					if !started.Contains(child.ID) {
						println(clock, "queueing", child.ID)
						queue = append(queue, child)
					}
				}
				worker.Active = nil
			}
		}
	}

	return clock
}

func newWorkerPool() []*worker {
	pool := make([]*worker, workerCount)
	for x := 0; x < workerCount; x++ {
		pool[x] = &worker{ID: x}
	}
	return pool
}

func getAvailable(pool []*worker) []*worker {
	var output []*worker
	for index := range pool {
		if pool[index].Idle() {
			output = append(output, pool[index])
		}
	}
	return output
}

func anyBusy(workers []*worker) bool {
	for index := range workers {
		if !workers[index].Idle() {
			return true
		}
	}
	return false
}

func getIdle(clock int, workers []*worker) *worker {
	for index := range workers {
		if workers[index].Idle() {
			return workers[index]
		}
	}
	return nil
}

type worker struct {
	ID       int
	TimeLeft int
	Active   *node
}

func (w *worker) Idle() bool {
	return w.Active == nil
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

type step struct {
	Before string
	After  string
}

type node struct {
	ID     string
	Before []*node
	After  []*node
}

func (n node) CanProcess(seen collections.SetOfString) bool {
	if len(n.Before) == 0 {
		return true
	}
	if seen.Contains(n.ID) {
		return false
	}
	for _, n := range n.Before {
		if !seen.Contains(n.ID) {
			return false
		}
	}
	return true
}

func (n node) Cost() int {
	lower := unicode.ToLower([]rune(n.ID)[0])
	lowerA := uint('a')
	return basis + int(uint(lower)-lowerA) + 1
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
