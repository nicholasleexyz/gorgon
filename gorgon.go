package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

type Coord struct {
	x int
	y int
}

func randrange(min int, max int) int {
	return rand.Intn(max-min+1) + min
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

	rand.Seed(time.Now().UnixNano())

	numcrds := 8
	crds := []Coord{}

	for i := 0; i < numcrds; i++ {
		crd := Coord{randrange(0, column-1), randrange(0, row-1)}
		crds = append(crds, crd)
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, cyan)
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
					img.Set(x, y, color.White)
				}
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
