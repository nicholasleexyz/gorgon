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

	column := 128
	row := 128

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}
	grey := color.RGBA{127, 127, 127, 0xff}
	rand.Seed(time.Now().UnixNano())

	numcrds := 8
	crds := []Cell{}

	for i := 0; i < numcrds; i++ {
		col := cyan
		if i%2 == 0 {
			col = color.RGBA{255, 255, 255, 0xff}
		}

		crd := Cell{randrange(0, column-1), randrange(0, row-1), col}
		crds = append(crds, crd)
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, grey)
		}
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			/*
				pick n amount of random coordinates
				for each pixel find the nearest of the random coords
				profit?! :P
				random number between a range
				calculate the distance between two 2-dimensional coordinates
			*/

			// normalized x and y
			nx := (x * column) / width
			ny := (y * row) / height

			for i := 0; i < numcrds; i++ {
				crd := crds[i]
				if nx == crd.x && ny == crd.y {
					img.Set(x, y, crd.color)
				}
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
