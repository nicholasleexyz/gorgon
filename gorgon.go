package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"
)

type Cell struct {
	x     int
	y     int
	color color.RGBA
}

func randrange(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func distancefromcell(x int, y int, cell Cell) float64 {
	//d=√((x2 – x1)² + (y2 – y1)²) // formula for distance
	x1 := float64(x)
	y1 := float64(y)
	x2 := float64(cell.x)
	y2 := float64(cell.y)

	return math.Sqrt(math.Pow((x2-x1), 2) + math.Pow((y2-y1), 2))
}

func main() {
	width := 512
	height := 512

	column := 512
	row := 512

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	rand.Seed(time.Now().UnixNano())

	numcrds := 8
	crds := [8]Cell{ // TODO: random color function
		{randrange(0, column-1), randrange(0, row-1), color.RGBA{255, 0, 0, 0xff}},
		{randrange(0, column-1), randrange(0, row-1), color.RGBA{0, 255, 0, 0xff}},
		{randrange(0, column-1), randrange(0, row-1), color.RGBA{0, 0, 255, 0xff}},
		{randrange(0, column-1), randrange(0, row-1), color.RGBA{255, 255, 0, 0xff}},
		{randrange(0, column-1), randrange(0, row-1), color.RGBA{0, 255, 255, 0xff}},
		{randrange(0, column-1), randrange(0, row-1), color.RGBA{255, 0, 255, 0xff}},
		{randrange(0, column-1), randrange(0, row-1), color.RGBA{127, 255, 127, 0xff}},
		{randrange(0, column-1), randrange(0, row-1), color.RGBA{255, 127, 127, 0xff}},
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			// normalized x and y
			nx := (x * column) / width
			ny := (y * row) / height

			min := distancefromcell(nx, ny, crds[0])
			col := crds[0].color

			for j := 1; j < numcrds; j++ {
				dist := distancefromcell(nx, ny, crds[j])
				if dist < min {
					min = dist
					col = crds[j].color
				}
			}

			img.Set(x, y, col)
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
