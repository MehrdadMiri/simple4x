package geometry

import "github.com/hajimehoshi/ebiten/v2"

type Line struct {
	start Point
	end   Point
}

func NewLine(start, end Point) *Line {
	return &Line{
		start: start,
		end:   end,
	}
}

func (l Line) Start() Point {
	return l.start
}

func (l Line) End() Point {
	return l.end
}

type Point struct {
	x, y float64
}

func (p Point) XY() (float64, float64) {
	return p.x, p.y
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

type Polygon interface {
	GetLines() []Line
	DrawLines(*ebiten.Image)
}
