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

type Vector2 struct {
	x float64
	y float64
}

type Cell struct {
	coords Vector2
	color  color.RGBA
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

func distance(a Vector2, b Vector2) float64 {
	//d=√((x2 – x1)² + (y2 – y1)²) // formula for distance
	return math.Sqrt(math.Pow((b.x-a.x), 2) + math.Pow((b.y-a.y), 2))
}

func magnitude(a Vector2) float64 {
	return math.Sqrt(math.Pow(a.x, 2) + math.Pow(a.y, 2))
}

func normalize(a Vector2) Vector2 {
	mag := magnitude(a)
	if mag == 0 {
		return a
	}

	return Vector2Div(a, mag)
}

func Vector2Div(a Vector2, x float64) Vector2 {
	return Vector2{a.x / x, a.y / x}
}

func Vector2Mult(a Vector2, x float64) Vector2 {
	return Vector2{a.x * x, a.y * x}
}

func Vector2Sub(a Vector2, b Vector2) Vector2 {
	return Vector2{a.x - b.x, a.y - b.y}
}

func Vector2Add(a Vector2, b Vector2) Vector2 {
	return Vector2{a.x + b.x, a.y + b.y}
}

func dotProduct(a Vector2, b Vector2) float64 {
	return a.x*b.x + a.y*a.y
}

func offsetCoords(cells []Cell, xOffset int, yOffset int) []Cell {
	crds := []Cell{}
	for i := 0; i < 64; i++ {
		x := int(cells[i].coords.x) + xOffset
		y := int(cells[i].coords.y) + yOffset
		c := cells[i].color
		crds = append(crds, Cell{Vector2{float64(x), float64(y)}, c})
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

	for i := 0; i < len(tcrds); i++ {
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
				Vector2{
					float64(i*unit + unit/2 + (randrange(-(unit/2 - margin), unit/2-5))),
					float64(j*unit + unit/2 + (randrange(-(unit/2 - margin), unit/2-5)))},
				randColor()})
			// i*unit + unit/2, // no offset
			// j*unit + unit/2,
			// shadeOfGrey(uint8((8*j + i) * 4))} // grey gradient for debugging
		}
	}

	tcrds := tileableCoords(crds, size)

	// closest := Vector2{tcrds[0].coords.x, tcrds[0].coords.y}
	// lastclosest := Vector2{closest.x, closest.y}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			current := Vector2{float64(x), float64(y)}
			min := distance(current, tcrds[0].coords)
			col := tcrds[0].color

			for j := 1; j < len(tcrds); j++ {
				dist := distance(current, tcrds[j].coords)
				if dist < min {
					min = dist

					current := tcrds[j]
					col = current.color

					// lastclosest = Vector2{closest.x, closest.y}
					// closest = Vector2{current.coords.x, current.coords.y}
				}
			}
			// toLastClosest := Vector2Sub(lastclosest, closest)
			// toClosest := Vector2Sub(closest, current)
			// difference := Vector2Sub(toClosest, toLastClosest)
			// center := Vector2Mult(Vector2Add(closest, lastclosest), 0.5)
			// pixToCenter := Vector2Sub(center, current)
			// dp := math.Abs(dotProduct(pixToCenter, difference))

			if min < float64(margin) {
				// || edgeDist < 0.5 { // draw a dot where the coord is
				col = color.RGBA{0, 0, 0, 0xff}
			}

			img.Set(x, y, col)
		}
	}

	for x := 0; x < size; x++ { // draw grid
		for y := 0; y < size; y++ {
			if x%unit == 0 || y%unit == 0 {
				img.Set(x, y, color.RGBA{127, 127, 127, 0xff})
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
