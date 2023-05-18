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
	size := 512
	unit := size / 8

	upLeft := image.Point{0, 0}
	lowRight := image.Point{size, size}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	rand.Seed(time.Now().UnixNano())

	crds := [64]Cell{}

	for j := 0; j < 8; j++ {
		for i := 0; i < 8; i++ {
			crds[8*j+i] = Cell{i*unit + unit/2 + (randrange(-(unit / 2), unit/2)), j*unit + unit/2 + (randrange(-(unit / 2), unit/2)), randColor()}
		}
	}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			min := distancefromcell(x, y, crds[0])
			col := crds[0].color

			for j := 1; j < 64; j++ {
				dist := distancefromcell(x, y, crds[j])
				if dist < min {
					min = dist
					col = crds[j].color
				}
			}

			if min < 5 {
				col = color.RGBA{0, 0, 0, 0xff}
			}

			img.Set(x, y, col)
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
