package main

import (
	"gorgon/vector"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

type Vector2Int = vector.Vector2Int
type Vector2 = vector.Vector2

type Cell struct {
	coords vector.Vector2Int
	color  color.RGBA
}

func main() {
	size := 1024
	unit := size / 8
	margin := 10
	drawCellCoord := false
	drawGrid := false
	lineWeight := 0.90 // smaller

	upLeft := image.Point{0, 0}
	lowRight := image.Point{size, size}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	var seed int64 = time.Now().UnixNano()
	rand.Seed(seed)

	cells := createCellGrid(8, unit, margin)
	tilableCells := tileableCoords(cells, 8, size)

	var closest Cell
	var secondClosest Cell

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			pixel := Vector2Int{X: x, Y: y}

			var min int = size

			for y1 := 0; y1 < len(tilableCells); y1++ {
				for x1 := 0; x1 < len(tilableCells[y1]); x1++ {
					cell := tilableCells[y1][x1]
					dist := pixel.Distance(cell.coords)

					if dist < min {
						min = dist
						closest = cell
					}
				}
			}

			img.Set(x, y, closest.color)
		}
	}

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			pixel := Vector2Int{X: x, Y: y}
			var min int = size

			for y1 := 0; y1 < len(tilableCells); y1++ {
				for x1 := 0; x1 < len(tilableCells[y1]); x1++ {
					cell := tilableCells[y1][x1]
					dist := pixel.Distance(cell.coords)

					if dist < min {
						min = dist
						secondClosest = closest
						closest = cell
					}
				}
			}

			if drawCellCoord && pixel.Distance(closest.coords) < margin { // draw cell coord
				img.Set(x, y, color.RGBA{0, 0, 0, 0xff})
			}

			center := closest.coords.Add(secondClosest.coords.Sub(closest.coords).Multiply(0.5))

			dirToSecondClosest := secondClosest.coords.Sub(closest.coords).ToVector2().Normalize()
			pixelToCenter := center.Sub(pixel).ToVector2()
			dprod := pixelToCenter.Dot(dirToSecondClosest.Multiply(1 - lineWeight))

			if dprod < 0.95 {
				img.Set(x, y, color.RGBA{0, 0, 0, 0xff})
			}
		}
	}

	if drawGrid {
		for x := 0; x < size; x++ { // draw grid
			for y := 0; y < size; y++ {
				if x%unit == 0 || y%unit == 0 || x == size-1 || y == size-1 {
					img.Set(x, y, color.RGBA{255, 255, 255, 0xff})
				}
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
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

func (cell Cell) OffsetCoords(offset Vector2Int) Cell {
	return Cell{cell.coords.Add(offset), cell.color}
}

func tileableCoords(cells [][]Cell, length int, size int) [][]Cell {
	len := length + 2
	tcells := [][]Cell{}
	offsetAmount := (size / length)
	var val Cell

	for y := 0; y < len; y++ {
		tcells = append(tcells, []Cell{})
		for x := 0; x < len; x++ {
			if x == 0 && y == 0 { // top left corner
				val = cells[length-1][length-1].OffsetCoords(Vector2Int{X: offsetAmount, Y: offsetAmount})
			} else if x == 0 && y == len-1 { // top right corner
				val = cells[0][length-1].OffsetCoords(Vector2Int{X: -offsetAmount * length, Y: offsetAmount * length})
			} else if x == len-1 && y == 0 { // bottom left corner
				val = cells[length-1][0].OffsetCoords(Vector2Int{X: offsetAmount * length, Y: -offsetAmount * length})
			} else if x == len-1 && y == len-1 { // bottom right corner
				val = cells[0][0].OffsetCoords(Vector2Int{X: -offsetAmount, Y: -offsetAmount})
			} else if x == 0 { // right edge
				val = cells[y-1][length-1].OffsetCoords(Vector2Int{X: -offsetAmount * length, Y: 0})
			} else if y == 0 { // bottom edge
				val = cells[length-1][x-1].OffsetCoords(Vector2Int{X: 0, Y: -offsetAmount * length})
			} else if x == len-1 { // left edge
				val = cells[y-1][0].OffsetCoords(Vector2Int{X: offsetAmount * length, Y: 0})
			} else if y == len-1 { // top edge
				val = cells[0][x-1].OffsetCoords(Vector2Int{X: 0, Y: offsetAmount * length})
			} else { // middle
				val = cells[y-1][x-1]
			}

			tcells[y] = append(tcells[y], val)
		}
	}

	return tcells
}

func randomOffset(unit int, margin int) int {
	return unit/2 + randrange(-(unit/2-margin), unit/2-margin)
}

func randomizedCoord(coord Vector2Int, unit int, margin int) Vector2Int {
	x := coord.X*unit + randomOffset(unit, margin)
	y := coord.Y*unit + randomOffset(unit, margin)
	return Vector2Int{X: x, Y: y}
}

func createCellGrid(cellSize int, unit int, margin int) [][]Cell {
	cells := [][]Cell{}
	for y := 0; y < cellSize; y++ {
		cells = append(cells, []Cell{})
		for x := 0; x < cellSize; x++ {
			cell := Cell{randomizedCoord(Vector2Int{X: x, Y: y}, unit, margin), randColor()}
			cells[y] = append(cells[y], cell)
		}
	}
	return cells
}
