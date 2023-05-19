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

func randColor() color.RGBA {
	r := uint8(randrange(0, 256))
	g := uint8(randrange(0, 256))
	b := uint8(randrange(0, 256))
	c := color.RGBA{r, g, b, 0xff}
	return c
}

func shadeOfGrey(i uint8) color.RGBA {
	c := color.RGBA{i, i, i, 0xff}
	return c
}

func randrange(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

/*
	detect edges
	normalize
	dot product
	abs
*/

func distancefromcell(x int, y int, cell Cell) float64 {
	//d=√((x2 – x1)² + (y2 – y1)²) // formula for distance
	x1 := float64(x)
	y1 := float64(y)
	x2 := float64(cell.x)
	y2 := float64(cell.y)

	return math.Sqrt(math.Pow((x2-x1), 2) + math.Pow((y2-y1), 2))
}

func offsetCoords(cells []Cell, xOffset int, yOffset int) []Cell {
	crds := []Cell{}
	for i := 0; i < 64; i++ {
		x := cells[i].x + xOffset
		y := cells[i].y + yOffset
		c := cells[i].color
		crds = append(crds, Cell{x, y, c})
	}

	return crds
}

func tileableCoords(cells []Cell, size int) []Cell {
	tcrds := [][]Cell{
		offsetCoords(cells, -size, size),
		offsetCoords(cells, 0, size),
		offsetCoords(cells, size, size),
		offsetCoords(cells, -size, 0),
		offsetCoords(cells, 0, 0),
		offsetCoords(cells, size, 0),
		offsetCoords(cells, -size, -size),
		offsetCoords(cells, 0, -size),
		offsetCoords(cells, size, -size),
	}
	arr := []Cell{}

	for i := 0; i < 9; i++ {
		arr = append(arr, tcrds[i]...)
	}

	return arr
}

func main() {
	size := 512
	unit := size / 8
	margin := 5

	upLeft := image.Point{0, 0}
	lowRight := image.Point{size, size}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	rand.Seed(time.Now().UnixNano())

	crds := []Cell{}

	for j := 0; j < 8; j++ {
		for i := 0; i < 8; i++ {
			crds = append(crds, Cell{
				i*unit + unit/2 + (randrange(-(unit/2 - margin), unit/2-5)),
				j*unit + unit/2 + (randrange(-(unit/2 - margin), unit/2-5)),
				randColor()})
			// i*unit + unit/2, // no offset
			// j*unit + unit/2,
			// shadeOfGrey(uint8((8*j + i) * 4))} // grey gradient for debugging
		}
	}

	tcrds := tileableCoords(crds, size)
	// fmt.Println(tcrds)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			min := distancefromcell(x, y, tcrds[0])
			col := tcrds[0].color

			for j := 1; j < len(tcrds); j++ {
				dist := distancefromcell(x, y, tcrds[j])
				if dist < min {
					min = dist
					col = tcrds[j].color
				}
			}

			// if min < float64(margin) { // draw a dot where the coord is
			// 	col = color.RGBA{0, 0, 0, 0xff}
			// }

			img.Set(x, y, col)
		}
	}

	// for x := 0; x < size; x++ { // draw grid
	// 	for y := 0; y < size; y++ {
	// 		if x%unit == 0 || y%unit == 0 {
	// 			img.Set(x, y, color.RGBA{127, 127, 127, 0xff})
	// 		}
	// 	}
	// }

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
