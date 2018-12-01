package main

import (
	"crypto/md5"
	"fmt"
	"strings"
	"sync"
)

const (
	AGENTS = 4
)

func hash(base, extra string) [16]byte {
	data := []byte(fmt.Sprintf("%s%s", base, extra))
	return md5.Sum(data)
}

func hashStartsWithZeros(hashData [16]byte) bool {
	hashAsString := fmt.Sprintf("%x", hashData)
	return strings.HasPrefix(hashAsString, "000000")
}

func main() {
	base := "bgvyzdsv"

	isDone := make(chan bool, 4)
	defer close(isDone)

	wg := sync.WaitGroup{}
	wg.Add(AGENTS)
	for i := 0; i < AGENTS; i++ {
		go func(agentId int) {
			defer wg.Done()
			hashAgent(base, agentId, AGENTS, isDone)
		}(i)
	}
	wg.Wait()
}

func hashAgent(base string, agentId, agents int, isDone chan bool) {
	extra := agentId
	for {
		select {
		case <-isDone:
			return
		default:
			extra = extra + agents
			hashData := hash(base, fmt.Sprintf("%d", extra))
			if hashStartsWithZeros(hashData) {
				fmt.Printf("%s%d\n", base, extra)

				for i := 0; i < agents; i++ {
					isDone <- true
				}
				break
			}
		}

	}
}
