package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
)

func main() {
	contents, err := ioutil.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	const (
		height = 6
		width  = 25
	)

	var layers []Image
	var x, y, layer int
	image := NewImage(width, height)
	for _, c := range contents {
		image[y][x] = byteToInt(byte(c))

		x++
		if x >= len(image[y]) {
			x = 0
			y++
		}
		if y >= len(image) {
			y = 0
			x = 0
			layer++
			layers = append(layers, image)
			image = NewImage(width, height)
		}
	}

	zeroCount, zeroCountIndex := int(math.MaxInt64), 0
	for index, layer := range layers {
		stats := computeStats(layer)
		if stats[0] < zeroCount {
			zeroCount = stats[0]
			zeroCountIndex = index
		}
	}

	finalStats := computeStats(layers[zeroCountIndex])
	fmt.Printf("final stats: %#v\n", finalStats)
	fmt.Printf("ANSWER %d\n", finalStats[1]*finalStats[2])
}

func computeStats(image Image) []int {
	output := make([]int, 10)
	for y := 0; y < len(image); y++ {
		for x := 0; x < len(image[y]); x++ {
			output[image[y][x]]++
		}
	}
	return output
}

func byteToInt(b byte) int {
	value, _ := strconv.Atoi(string([]byte{b}))
	return value
}

// NewImage returns a new image
func NewImage(width, height int) Image {
	image := make(Image, height)
	for y := 0; y < height; y++ {
		image[y] = make([]int, width)
	}
	return image
}

// Image is a fixed 6 row and 25 column grid of integers.
type Image [][]int
