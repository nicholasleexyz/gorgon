package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	width := 512
	height := 512

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			nx := (x * 4) / width
			ny := (y * 4) / height

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
