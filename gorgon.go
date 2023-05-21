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

type Vector2Int struct {
	x int
	y int
}

type Vector2 struct {
	x float64
	y float64
}

type Cell struct {
	coords Vector2Int
	color  color.RGBA
}

func main() {
	size := 512
	unit := size / 8
	margin := 5

	upLeft := image.Point{0, 0}
	lowRight := image.Point{size, size}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	rand.Seed(time.Now().UnixNano())

	cells := createCellGrid(8, unit, margin)
	tilableCells := tileableCoords(cells, 8, size)

	var closest Cell
	var secondClosest Cell

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			pixel := Vector2Int{x, y}

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
			pixel := Vector2Int{x, y}
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

			// if pixel.Distance(closest.coords) < margin { // draw cell coord
			// 	img.Set(x, y, color.RGBA{0, 0, 0, 0xff})
			// }

			center := closest.coords.Add(secondClosest.coords.Sub(closest.coords).Multiply(0.5))
			// if pixel.Distance(center) < margin {
			// 	img.Set(x, y, color.RGBA{127, 127, 127, 0xff})
			// }

			dirToSecondClosest := secondClosest.coords.Sub(closest.coords).toVector2().normalize()
			pixelToCenter := center.Sub(pixel).toVector2()
			dprod := pixelToCenter.Dot(dirToSecondClosest.Multiply(0.15))

			if dprod < 0.95 {
				img.Set(x, y, color.RGBA{0, 0, 0, 0xff})
			}
		}
	}

	// for x := 0; x < size; x++ { // draw grid
	// 	for y := 0; y < size; y++ {
	// 		if x%unit == 0 || y%unit == 0 || x == size-1 || y == size-1 {
	// 			img.Set(x, y, color.RGBA{255, 255, 255, 0xff})
	// 		}
	// 	}
	// }

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

func invertColor(col color.RGBA) color.RGBA {
	r := (col.R + 127) % 255
	g := (col.G + 127) % 255
	b := (col.B + 127) % 255
	return color.RGBA{r, g, b, 0xff}
}

func (a Vector2Int) toVector2() Vector2 {
	return Vector2{float64(a.x), float64(a.y)}
}

func (a Vector2) toVector2Int() Vector2Int {
	return Vector2Int{int(a.x), int(a.y)}
}

func (a Vector2Int) magnitude() float64 {
	ax := float64(a.x)
	ay := float64(a.y)
	return math.Sqrt(math.Pow(ax, 2) + math.Pow(ay, 2))
}

func (a Vector2) magnitude() float64 {
	return math.Sqrt(math.Pow(a.x, 2) + math.Pow(a.y, 2))
}

func (a Vector2) normalize() Vector2 {
	mag := a.magnitude()
	if mag == 0 {
		return a
	}

	return a.Divide(mag)
}

func (a Vector2Int) normalize() Vector2Int {
	mag := a.magnitude()
	if mag == 0 {
		return a
	}

	return a.Divide(mag)
}

func (a Vector2) Sub(b Vector2) Vector2 {
	cx := a.x - b.x
	cy := a.y - b.y
	return Vector2{cx, cy}
}

func (a Vector2Int) Sub(b Vector2Int) Vector2Int {
	cx := a.x - b.x
	cy := a.y - b.y
	return Vector2Int{cx, cy}
}

func (a Vector2Int) Add(b Vector2Int) Vector2Int {
	cx := a.x + b.x
	cy := a.y + b.y
	return Vector2Int{cx, cy}
}

func (a Vector2) Add(b Vector2) Vector2 {
	cx := a.x + b.x
	cy := a.y + b.y
	return Vector2{cx, cy}
}

func (a Vector2Int) Divide(x float64) Vector2Int {
	ax := int(math.Round(float64(a.x) / x))
	ay := int(math.Round(float64(a.y) / x))
	return Vector2Int{ax, ay}
}

func (a Vector2) Divide(x float64) Vector2 {
	ax := a.x / x
	ay := a.y / x
	return Vector2{ax, ay}
}

func (a Vector2Int) Multiply(x float64) Vector2Int {
	ax := int(math.Round(float64(a.x) * x))
	ay := int(math.Round(float64(a.y) * x))
	return Vector2Int{ax, ay}
}
func (a Vector2) Multiply(x float64) Vector2 {
	ax := a.x * x
	ay := a.y * x
	return Vector2{ax, ay}
}

func (a Vector2) Distance(b Vector2) float64 {
	//d=√((x2 – x1)² + (y2 – y1)²) // formula for distance

	return math.Abs(math.Sqrt(math.Pow((b.x-a.x), 2) + math.Pow((b.y-a.y), 2)))
}

func (a Vector2Int) Distance(b Vector2Int) int {
	//d=√((x2 – x1)² + (y2 – y1)²) // formula for distance
	ax := float64(a.x)
	ay := float64(a.y)
	bx := float64(b.x)
	by := float64(b.y)

	return int(math.Round(math.Abs(math.Sqrt(math.Pow((bx-ax), 2) + math.Pow((by-ay), 2)))))
}

func (a Vector2) Dot(b Vector2) float64 {
	ax := a.x
	ay := a.y
	bx := b.x
	by := b.y

	dprod := (ax * bx) + (ay * by)

	if dprod > 1 {
		return 1
	} else if dprod < -1 {
		return -1
	}

	return dprod
}

func (cell Cell) OffsetCoords(offset Vector2Int) Cell {
	return Cell{cell.coords.Add(offset), cell.color}
}

func tileableCoords(cells [][]Cell, length int, size int) [][]Cell {
	len := length + 2
	tcells := [][]Cell{}
	offsetAmount := (size / length)

	for y := 0; y < len; y++ {
		tcells = append(tcells, []Cell{})
		for x := 0; x < len; x++ {
			if x == 0 && y == 0 { // top left corner
				val := cells[length-1][length-1].OffsetCoords(Vector2Int{offsetAmount, offsetAmount})
				tcells[y] = append(tcells[y], val)
			} else if x == 0 && y == len-1 { // top right corner
				val := cells[0][length-1].OffsetCoords(Vector2Int{-offsetAmount * length, offsetAmount * length})
				tcells[y] = append(tcells[y], val)
			} else if x == len-1 && y == 0 { // bottom left corner
				val := cells[length-1][0].OffsetCoords(Vector2Int{offsetAmount * length, -offsetAmount * length})
				tcells[y] = append(tcells[y], val)
			} else if x == len-1 && y == len-1 { // bottom right corner
				val := cells[0][0].OffsetCoords(Vector2Int{-offsetAmount, -offsetAmount})
				tcells[y] = append(tcells[y], val)
			} else if x == 0 { // right edge
				val := cells[y-1][length-1].OffsetCoords(Vector2Int{-offsetAmount * length, 0})
				tcells[y] = append(tcells[y], val)
			} else if y == 0 { // bottom edge
				val := cells[length-1][x-1].OffsetCoords(Vector2Int{0, -offsetAmount * length})
				tcells[y] = append(tcells[y], val)
			} else if x == len-1 { // left edge
				val := cells[y-1][0].OffsetCoords(Vector2Int{offsetAmount * length, 0})
				tcells[y] = append(tcells[y], val)
			} else if y == len-1 { // top edge
				val := cells[0][x-1].OffsetCoords(Vector2Int{0, offsetAmount * length})
				tcells[y] = append(tcells[y], val)
			} else { // middle
				val := cells[y-1][x-1]
				tcells[y] = append(tcells[y], val)
			}
		}
	}

	return tcells
}

func randomOffset(unit int, margin int) int {
	return unit/2 + randrange(-(unit/2-margin), unit/2-margin)
}

func randomizedCoord(coord Vector2Int, unit int, margin int) Vector2Int {
	x := coord.x*unit + randomOffset(unit, margin)
	y := coord.y*unit + randomOffset(unit, margin)
	return Vector2Int{x, y}
}

func createCellGrid(cellSize int, unit int, margin int) [][]Cell {
	cells := [][]Cell{}
	for y := 0; y < cellSize; y++ {
		cells = append(cells, []Cell{})
		for x := 0; x < cellSize; x++ {
			cell := Cell{randomizedCoord(Vector2Int{x, y}, unit, margin), randColor()}
			// cell := Cell{Vector2{x*unit + unit/2, y*unit + unit/2}, randColor()}
			// cell := Cell{Vector2{(x+1)*unit + unit/2, (y+1)*unit + unit/2}, randColor()}
			cells[y] = append(cells[y], cell)
		}
	}
	return cells
}

func cellNeighbors(x int, y int, cellSize int, cells [][]Cell) []Cell {
	nbrs := []Cell{
		cells[(y+1)%cellSize][(x-1+cellSize)%cellSize],
		cells[(y+1)%cellSize][x%cellSize],
		cells[(y+1)%cellSize][(x+1)%cellSize],
		cells[y%cellSize][(x-1+cellSize)%cellSize],
		cells[y%cellSize][(x+1)%cellSize],
		cells[(y-1+cellSize)%cellSize][(x-1+cellSize)%cellSize],
		cells[(y-1+cellSize)%cellSize][x%cellSize],
		cells[(y-1+cellSize)%cellSize][(x+1)%cellSize],
	}

	return nbrs
}

func (cell Cell) SetZero() Cell { // usful for debugging
	return Cell{Vector2Int{0, 0}, cell.color}
}

// func (a Vector2) Invert() Vector2 {
// 	return a.Multiply(-1)
// }

func shadeOfGrey(i uint8) color.RGBA {
	c := color.RGBA{i, i, i, 0xff}
	return c
}

// func (a Vector2) Abs() Vector2 {
// 	return Vector2{int(math.Abs(float64(a.x))), int(math.Abs(float64(a.y)))}
// }
