package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/wcharczuk/advent/pkg/fileutil"
	"github.com/wcharczuk/advent/pkg/log"
)

const (
	width  = 256
	height = 64
)

func main() {
	var rawPoints []*point
	err := fileutil.ReadByLines("input", func(line string) error {
		var point point
		var position, velocity string
		matches := regexExtract(line, `position=<(.*)> velocity=<(.*)>`)
		position = matches[1]
		velocity = matches[2]
		point.X0, point.Y0 = parse(position)
		point.VX, point.VY = parse(velocity)
		point.X = point.X0
		point.Y = point.Y0
		rawPoints = append(rawPoints, &point)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	os.RemoveAll("output")
	os.Mkdir("output", 0755)

	pts := points(rawPoints).Copy()
	minDistance := math.MaxFloat64
	var index int
	var minSet points
	for x := 0; x < 11000; x++ {
		pts.Advance()
		distance := pts.AverageDistance()
		if distance < minDistance {
			minDistance = distance
			minSet = pts.Copy()
			index = x
		}
	}

	buffer := new(bytes.Buffer)
	minSet.Write(buffer, width, height)
	ioutil.WriteFile("output/image.png", buffer.Bytes(), 0644)
	println(index)
}

func regexExtract(corpus, expr string) []string {
	re := regexp.MustCompile(expr)
	allResults := re.FindAllStringSubmatch(corpus, -1)
	results := []string{}
	for _, resultSet := range allResults {
		for _, result := range resultSet {
			results = append(results, result)
		}
	}
	return results
}

func parse(str string) (int, int) {
	str = strings.Replace(str, " ", "", -1)
	pieces := strings.Split(str, ",")
	a, _ := strconv.Atoi(pieces[0])
	b, _ := strconv.Atoi(pieces[1])
	return a, b
}

type point struct {
	X0, Y0 int
	VX, VY int
	X, Y   int
}

type points []*point

func (pts points) AverageDistance() float64 {
	var total float64
	var count int
	for i0, p0 := range pts {
		for i1, p1 := range pts {
			if i0 != i1 {
				x := float64(p1.X) - float64(p0.X)
				x2 := x * x
				y := float64(p1.Y) - float64(p0.Y)
				y2 := y * y
				total += math.Sqrt(x2 + y2)
				count++
			}
		}
	}
	return total / float64(count)
}

func (pts points) Write(r io.Writer, width, height int) error {
	buffer := image.NewRGBA(image.Rect(0, 0, width, height))

	minx, miny, maxx, maxy := pts.Bounds()
	println(minx, miny, maxx, maxy)

	nmaxy := maxy - miny
	nmaxx := maxx - minx

	for _, p := range pts {
		pyn := float64(p.Y-miny) / float64(nmaxy+1)
		pxn := float64(p.X-minx) / float64(nmaxx+1)

		py := int(pyn * float64(height))
		px := int(pxn * float64(width))

		buffer.Set(px, py, ColorRed)
	}

	return png.Encode(r, buffer)
}

func (pts points) Advance() {
	for _, p := range pts {
		p.X = p.X + p.VX
		p.Y = p.Y + p.VY
	}
}

func (pts points) Bounds() (minx, miny, maxx, maxy int) {
	minx, miny = int(math.MaxInt64), int(math.MaxInt64)
	maxx, maxy = int(math.MinInt64), int(math.MinInt64)
	for _, p := range pts {
		if p.X < minx {
			minx = p.X
		}
		if maxx < p.X {
			maxx = p.X
		}
		if p.Y < miny {
			miny = p.Y
		}
		if maxy < p.Y {
			maxy = p.Y
		}
	}
	return
}

func (pts points) Copy() points {
	output := make([]*point, len(pts))
	for x := 0; x < len(pts); x++ {
		v := *pts[x]
		output[x] = &point{
			X:  v.X,
			Y:  v.Y,
			X0: v.X0,
			Y0: v.Y0,
			VX: v.VX,
			VY: v.VY,
		}
	}
	return output
}

var (
	// ColorRed is red.
	ColorRed = Color{R: 255, G: 0, B: 0, A: 255}
)

// Color is our internal color type because color.Color is bullshit.
type Color struct {
	R, G, B, A uint8
}

// RGBA returns the color as a pre-alpha mixed color set.
func (c Color) RGBA() (r, g, b, a uint32) {
	fa := float64(c.A) / 255.0
	r = uint32(float64(uint32(c.R)) * fa)
	r |= r << 8
	g = uint32(float64(uint32(c.G)) * fa)
	g |= g << 8
	b = uint32(float64(uint32(c.B)) * fa)
	b |= b << 8
	a = uint32(c.A)
	a |= a << 8
	return
}

func square(i *image.RGBA, x, y int, width int, c color.Color) {
	w2 := width >> 1
	for y0 := y - w2; y0 < y+w2; y0++ {
		for x0 := x - w2; x0 < x+w2; x0++ {
			i.Set(x0, y0, c)
		}
	}
}
