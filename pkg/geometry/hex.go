package geometry

import "math"

type Hex struct {
	x           float64
	y           float64
	r           float64
	row, column int
	lines       []Line
}

func NewHex(x, y, r float64, row, column int) *Hex {
	lines := make([]Line, 6)
	points := make([]Point, 6)
	rotate30 := complex(math.Cos(math.Pi/6), math.Sin(math.Pi/6))
	rotate60 := complex(math.Cos(math.Pi/3), math.Sin(math.Pi/3))
	p := complex(r, 0)
	p = p * rotate30
	for i := 0; i < 6; i++ {
		points[i] = *NewPoint(real(p)+x, imag(p)+y)
		p = p * rotate60
	}
	for i := 0; i < 6; i++ {
		lines[i] = *NewLine(points[i], points[(i+1)%6])
	}
	return &Hex{
		x:      x,
		y:      y,
		r:      r,
		row:    row,
		column: column,
		lines:  lines,
	}
}

func (h Hex) GetLines() []Line {
	return h.lines
}
