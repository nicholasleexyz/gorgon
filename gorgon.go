package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

func randRange(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

type Coord struct {
	x int
	y int
}

func returnCoord(x int, y int) Coord {
	return Coord{x, y}
}

func main() {
	width := 512
	height := 512

	column := 16
	row := 16

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}

	rand.Seed(time.Now().UnixNano())

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			/*

				pick n amount of random coordinates
				for each pixel find the nearest of the random coords

				profit?! :P

				random number between a range
				calculate the distance between two 2-dimensional coordinates

			*/

			//normalized x and y could also be used with this

			nx := (x * column) / width
			ny := (y * row) / height

			if nx%2 == ny%2 {
				img.Set(x, y, cyan)
			} else {
				img.Set(x, y, color.White)
			}

		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
