package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	black       = 0
	white       = 1
	transparent = 2
)
const (
	height = 6
	width  = 25
)

func main() {
	contents, err := ioutil.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

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

	fmt.Println(Layers(layers).Visible().String())
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

// Layers are a set of images.
type Layers []Image

// Visible flattens the image into visible pixels.
func (l Layers) Visible() Image {
	final := NewImage(width, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			for z := 0; z < len(l); z++ {
				if l[z][y][x] == transparent {
					continue
				}
				final[y][x] = l[z][y][x]
				break
			}
		}
	}
	return final
}

// Image is a fixed 6 row and 25 column grid of integers.
type Image [][]int

func (i Image) String() string {
	var lines []string
	for y := 0; y < len(i); y++ {
		var line []string
		for x := 0; x < len(i[y]); x++ {
			if i[y][x] == black {
				line = append(line, " ")
			} else {
				line = append(line, "X")
			}
		}
		lines = append(lines, strings.Join(line, ""))
	}
	return strings.Join(lines, "\n")
}
