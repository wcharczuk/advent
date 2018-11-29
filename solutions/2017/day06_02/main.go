package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	util "github.com/blendlabs/go-util"
	"github.com/blendlabs/go-util/collections"
)

type bank struct {
	blocks []int
}

func (b *bank) max() (maxIndex, maxValue int) {
	for index := 0; index < len(b.blocks); index++ {
		if b.blocks[index] > maxValue {
			maxValue = b.blocks[index]
			maxIndex = index
		}
	}
	return
}

func (b *bank) redistribute() {
	index, count := b.max()
	// zero the index
	b.blocks[index] = 0

	index = (index + 1) % len(b.blocks)
	for count > 0 {
		b.blocks[index]++
		count--
		index = (index + 1) % len(b.blocks)
	}
}

func (b *bank) hash() string {
	buffer := bytes.NewBuffer(nil)

	for _, v := range b.blocks {
		buffer.WriteString(fmt.Sprintf("%d", v))
		buffer.WriteRune(rune('-'))
	}

	return buffer.String()
}

func main() {
	// read file
	contents, err := ioutil.ReadFile("./testdata/input")
	if err != nil {
		log.Fatal(err)
	}
	// break into blocks
	words := util.String.SplitOnSpace(string(contents))

	// create bank
	b := &bank{}
	for _, w := range words {
		v, err := strconv.Atoi(w)
		if err != nil {
			log.Fatal(err)
		}
		b.blocks = append(b.blocks, v)
	}

	// keep redistributing until we have a collision
	hashLookup := collections.NewSetOfString(b.hash())

	var hash string
	for {
		b.redistribute()
		hash = b.hash()
		if hashLookup.Contains(hash) {
			break
		}
		hashLookup.Add(hash)
	}
	count := 1
	checkHash := hash
	for {
		b.redistribute()
		hash = b.hash()
		if hash == checkHash {
			break
		}
		count++
	}
	println(count)
}
