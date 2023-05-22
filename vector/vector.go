package vector

import "math"

type Vector2Int struct {
	X int
	Y int
}

type Vector2 struct {
	X float64
	Y float64
}

func (a Vector2Int) ToVector2() Vector2 {
	return Vector2{float64(a.X), float64(a.Y)}
}

func (a Vector2) ToVector2Int() Vector2Int {
	return Vector2Int{int(a.X), int(a.Y)}
}

func (a Vector2Int) magnitude() float64 {
	ax := float64(a.X)
	ay := float64(a.Y)
	return math.Sqrt(math.Pow(ax, 2) + math.Pow(ay, 2))
}

func (a Vector2) Magnitude() float64 {
	return math.Sqrt(math.Pow(a.X, 2) + math.Pow(a.Y, 2))
}

func (a Vector2) Normalize() Vector2 {
	mag := a.Magnitude()
	if mag == 0 {
		return a
	}

	return a.Divide(mag)
}

func (a Vector2Int) Normalize() Vector2Int {
	mag := a.magnitude()
	if mag == 0 {
		return a
	}

	return a.Divide(mag)
}

func (a Vector2) Sub(b Vector2) Vector2 {
	cx := a.X - b.X
	cy := a.Y - b.Y
	return Vector2{cx, cy}
}

func (a Vector2Int) Sub(b Vector2Int) Vector2Int {
	cx := a.X - b.X
	cy := a.Y - b.Y
	return Vector2Int{cx, cy}
}

func (a Vector2Int) Add(b Vector2Int) Vector2Int {
	cx := a.X + b.X
	cy := a.Y + b.Y
	return Vector2Int{cx, cy}
}

func (a Vector2) Add(b Vector2) Vector2 {
	cx := a.X + b.X
	cy := a.Y + b.Y
	return Vector2{cx, cy}
}

func (a Vector2Int) Divide(x float64) Vector2Int {
	ax := int(math.Round(float64(a.X) / x))
	ay := int(math.Round(float64(a.Y) / x))
	return Vector2Int{ax, ay}
}

func (a Vector2) Divide(x float64) Vector2 {
	ax := a.X / x
	ay := a.Y / x
	return Vector2{ax, ay}
}

func (a Vector2Int) Multiply(x float64) Vector2Int {
	ax := int(math.Round(float64(a.X) * x))
	ay := int(math.Round(float64(a.Y) * x))
	return Vector2Int{ax, ay}
}
func (a Vector2) Multiply(x float64) Vector2 {
	ax := a.X * x
	ay := a.Y * x
	return Vector2{ax, ay}
}

func (a Vector2) Distance(b Vector2) float64 {
	//d=√((x2 – x1)² + (y2 – y1)²) // formula for distance

	return math.Abs(math.Sqrt(math.Pow((b.X-a.X), 2) + math.Pow((b.Y-a.Y), 2)))
}

func (a Vector2Int) Distance(b Vector2Int) int {
	//d=√((x2 – x1)² + (y2 – y1)²) // formula for distance
	ax := float64(a.X)
	ay := float64(a.Y)
	bx := float64(b.X)
	by := float64(b.Y)

	return int(math.Round(math.Abs(math.Sqrt(math.Pow((bx-ax), 2) + math.Pow((by-ay), 2)))))
}

func (a Vector2) Dot(b Vector2) float64 {
	ax := a.X
	ay := a.Y
	bx := b.X
	by := b.Y

	dprod := (ax * bx) + (ay * by)

	if dprod > 1 {
		return 1
	} else if dprod < -1 {
		return -1
	}

	return dprod
}
