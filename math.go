package dframe

type Point struct { // A basic float64 point struct.
	X float64
	Y float64
}

func (p1 Point) Div(p2 Point) Point {
	p1.X /= p2.X
	p1.Y /= p2.Y
	return p1
}

func (p1 Point) Mul(p2 Point) Point {
	p1.X *= p2.X
	p1.Y *= p2.Y
	return p1
}

func (p1 Point) Add(p2 Point) Point {
	p1.X += p2.X
	p1.Y += p2.Y
	return p1
}

func (p1 Point) Sub(p2 Point) Point {
	p1.X -= p2.X
	p1.Y -= p2.Y
	return p1
}

type IntPoint struct { // A basic int point struct.
	X int
	Y int
}

func (p1 IntPoint) Div(p2 IntPoint) IntPoint {
	p1.X /= p2.X
	p1.Y /= p2.Y
	return p1
}

func (p1 IntPoint) Mul(p2 IntPoint) IntPoint {
	p1.X *= p2.X
	p1.Y *= p2.Y
	return p1
}

func (p1 IntPoint) Add(p2 IntPoint) IntPoint {
	p1.X += p2.X
	p1.Y += p2.Y
	return p1
}

func (p1 IntPoint) Sub(p2 IntPoint) IntPoint {
	p1.X -= p2.X
	p1.Y -= p2.Y
	return p1
}
